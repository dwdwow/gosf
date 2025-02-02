package gosf

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tx CBTx
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	select {
	case s.txCh <- tx:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Channel full", http.StatusServiceUnavailable)
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

func (s *CallbackServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hlaskjdfyreiqwt658-91740agjhs-ldvhjk19asdkvy5", s.handleTest)
	mux.HandleFunc("/txcallback-hlsfdgj654389127dsfgbhjague", s.handleTxCallback)
	mux.HandleFunc("/acctcallback-hlsfdgj654389127dsfgbhjague", s.handleAcctCallback)
	log.Printf("Starting HTTPS callback server on port %s", s.port)
	return http.ListenAndServeTLS(":"+s.port, s.certFile, s.keyFile, mux)
}
