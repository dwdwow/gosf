package gosf

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WsMsgHandler interface {
	Handle(msg []byte)
}

type WsClient struct {
	ul      string
	ws      *websocket.Conn
	handler WsMsgHandler
	logger  *slog.Logger
}

func NewWsClient(url string, handler WsMsgHandler, logger *slog.Logger) *WsClient {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	c := &WsClient{
		ul:      url,
		handler: handler,
		logger:  logger,
	}
	return c
}

func (c *WsClient) Start() error {
	ws, _, err := websocket.DefaultDialer.Dial(c.ul, nil)
	if err != nil {
		return err
	}
	c.ws = ws
	go c.listener()
	return nil
}

func (c *WsClient) listener() {
	for {
		t, b, err := c.ws.ReadMessage()
		if err != nil {
			c.logger.Error("ws listener", "error", "Failed to read message", "error", err)
			break
		}
		switch t {
		case websocket.PingMessage:
			c.logger.Info("ws ping message")
		case websocket.PongMessage:
			c.logger.Info("ws pong message")
		case websocket.TextMessage:
			c.handler.Handle(b)
		case websocket.BinaryMessage:
			c.logger.Info("ws binary message", "message", string(b))
		case websocket.CloseMessage:
			c.logger.Info("ws close message")
		}
	}
}

type SigChecker struct {
	mux  sync.Mutex
	sigs map[string]bool
}

func NewSigChecker() *SigChecker {
	return &SigChecker{
		sigs: make(map[string]bool),
	}
}

func (c *SigChecker) IsExist(sig string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.sigs[sig] {
		return true
	}
	c.sigs[sig] = true
	go func() {
		time.Sleep(10 * time.Second)
		c.mux.Lock()
		delete(c.sigs, sig)
		c.mux.Unlock()
	}()
	return false
}

type SimpleWsMsgHandler struct {
	sigChk *SigChecker
	logger *slog.Logger
}

func NewSimpleWsMsgHandler(logger *slog.Logger) *SimpleWsMsgHandler {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &SimpleWsMsgHandler{
		sigChk: NewSigChecker(),
		logger: logger,
	}
}

func (h *SimpleWsMsgHandler) Handle(msg []byte) {
	var tx Tx
	err := json.Unmarshal(msg, &tx)
	if err != nil {
		h.logger.Error("ws message", "raw", string(msg), "error", err)
		return
	}
	if tx.Timestamp == "" {
		h.logger.Error("ws message", "raw", string(msg), "error", "timestamp is empty")
		return
	}
	for _, sig := range tx.Signatures {
		if h.sigChk.IsExist(sig) {
			return
		}
	}
	// n := time.Now()
	// t, err := time.Parse(time.RFC3339Nano, tx.Timestamp)
	// if err != nil {
	// 	h.logger.Error("ws message", "raw", tx.Timestamp, "error", err)
	// 	panic(1)
	// }
	// du := n.Sub(t)
	fmt.Println(tx.Protocol.Address, tx.Protocol.Name)
}
