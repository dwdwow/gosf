package gosf

import (
	"io"
	"log/slog"
	"net/http"
	"os"
)

const (
	CB_SERVER_TX_CALLBACK_PATH   = "/txcallback-hlsfdgj654389127dsfgbhjague"
	CB_SERVER_ACCT_CALLBACK_PATH = "/acctcallback-hlsfdgj654389127dsfgbhjague"
	CB_SERVER_TEST_PATH          = "/test-hlsfdgj654389127dsfgbhjague"
)

type CallbackServer struct {
	port   string
	txWs   *SimpleWSServer
	acctWs *SimpleWSServer
	logger *slog.Logger
}

func NewCallbackServer(httpPort, txWsPort, acctWsPort string, logger *slog.Logger) *CallbackServer {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &CallbackServer{
		port:   httpPort,
		txWs:   NewSimpleWSServer(txWsPort, nil),
		acctWs: NewSimpleWSServer(acctWsPort, nil),
		logger: logger,
	}
}

func (s *CallbackServer) handleTest(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("test request", "method", r.Method, "ip", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		s.logger.Error("test request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to write response", "error", err)
	}
}

func (s *CallbackServer) handleTxCallback(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("tx callback request", "method", r.Method, "ip", r.RemoteAddr)

	if r.Method != http.MethodPost {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Method not allowed")
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to read body", "error", err)
		return
	}
	s.txWs.Broadcast(bytes)
}

func (s *CallbackServer) handleAcctCallback(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("acct callback request", "method", r.Method, "ip", r.RemoteAddr)
	if r.Method != http.MethodPost {
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Method not allowed")
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to read body", "error", err)
		return
	}
	s.acctWs.Broadcast(bytes)
}

func (s *CallbackServer) Serve() {
	s.txWs.Serve("/")
	s.acctWs.Serve("/")
	mux := http.NewServeMux()
	mux.HandleFunc(CB_SERVER_TEST_PATH, s.handleTest)
	mux.HandleFunc(CB_SERVER_TX_CALLBACK_PATH, s.handleTxCallback)
	mux.HandleFunc(CB_SERVER_ACCT_CALLBACK_PATH, s.handleAcctCallback)
	s.logger.Info("Starting HTTP callback server", "port", s.port)
	err := http.ListenAndServe(":"+s.port, mux)
	if err != nil {
		s.logger.Error("Failed to start HTTP callback server", "error", err)
	}
}
