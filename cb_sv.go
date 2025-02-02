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
	txCh     chan Tx
	acctCh   chan Acct
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
		txCh:     make(chan Tx, 100),
		acctCh:   make(chan Acct, 100),
		logger:   logger,
	}
}

func (s *CallbackServer) GetTxChannel() <-chan Tx {
	return s.txCh
}

func (s *CallbackServer) GetAcctChannel() <-chan Acct {
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

	var tx Tx
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Invalid request body", "error", err)
		return
	}

	if tx.Timestamp == "" {
		return
	}

	if err := ParseTxActions(&tx); err != nil {
		s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to parse tx actions", "error", err)
		return
	}

	// select {
	// case s.txCh <- tx:
	// default:
	// 	s.logger.Error("tx callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Channel full")
	// }
}

func (s *CallbackServer) handleAcctCallback(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("acct callback request", "method", r.Method, "ip", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("\n"))
	if err != nil {
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to write response", "error", err)
	}

	if r.Method != http.MethodPost {
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Method not allowed")
		return
	}

	var acct Acct
	if err := json.NewDecoder(r.Body).Decode(&acct); err != nil {
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Invalid request body", "error", err)
		return
	}

	if acct.Account == "" {
		return
	}

	select {
	case s.acctCh <- acct:
	default:
		s.logger.Error("acct callback request", "method", r.Method, "ip", r.RemoteAddr, "error", "Channel full")
	}
}

func (s *CallbackServer) StartHTTPServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hlaskjdfyreiqwt658-91740agjhs-ldvhjk19asdkvy5", s.handleTest)
	mux.HandleFunc(CB_SERVER_TX_CALLBACK_PATH, s.handleTxCallback)
	mux.HandleFunc(CB_SERVER_ACCT_CALLBACK_PATH, s.handleAcctCallback)
	s.logger.Info("Starting HTTP callback server", "port", s.port)
	go func() {
		err := http.ListenAndServe(":"+s.port, mux)
		if err != nil {
			s.logger.Error("Failed to start HTTP callback server", "error", err)
		}
	}()
	return nil
}
