package gosf

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

const (
	CB_SERVER_TX_CALLBACK_PATH   = "/txcallback-hlsfdgj654389127dsfgbhjague"
	CB_SERVER_ACCT_CALLBACK_PATH = "/acctcallback-hlsfdgj654389127dsfgbhjague"
)

type CallbackServer struct {
	port     string
	certFile string
	keyFile  string
	txCh     chan CBTx
	acctCh   chan CBAcct
	logger   *slog.Logger
}

func NewCallbackServer(port, certFile, keyFile string, logger *slog.Logger) *CallbackServer {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &CallbackServer{
		port:     port,
		certFile: certFile,
		keyFile:  keyFile,
		txCh:     make(chan CBTx, 100),
		acctCh:   make(chan CBAcct, 100),
		logger:   logger,
	}
}

func (s *CallbackServer) GetTxChannel() <-chan CBTx {
	return s.txCh
}

func (s *CallbackServer) GetAcctChannel() <-chan CBAcct {
	return s.acctCh
}

func (s *CallbackServer) handleTest(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("test request", "method", r.Method, "ip", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func (s *CallbackServer) handleTxCallback(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("tx callback request", "method", r.Method, "ip", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("\n"))
	if err != nil {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to write response", "error", err)
	}

	if r.Method != http.MethodPost {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Method not allowed")
		return
	}

	var tx CBTx
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Invalid request body", "error", err)
		return
	}

	select {
	case s.txCh <- tx:
		s.logger.Info("send tx to channel")
	default:
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Channel full")
	}
}

func (s *CallbackServer) handleAcctCallback(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("acct callback request", "method", r.Method, "ip", r.RemoteAddr)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var acct CBAcct
	if err := json.NewDecoder(r.Body).Decode(&acct); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	select {
	case s.acctCh <- acct:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Channel full", http.StatusServiceUnavailable)
	}
}

func (s *CallbackServer) HTTPServe() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hlaskjdfyreiqwt658-91740agjhs-ldvhjk19asdkvy5", s.handleTest)
	mux.HandleFunc(CB_SERVER_TX_CALLBACK_PATH, s.handleTxCallback)
	mux.HandleFunc(CB_SERVER_ACCT_CALLBACK_PATH, s.handleAcctCallback)
	s.logger.Info("Starting HTTPS callback server", "port", s.port, "certFile", s.certFile, "keyFile", s.keyFile)
	return http.ListenAndServe(":"+s.port, mux)
}
