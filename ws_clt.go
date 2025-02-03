package gosf

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"

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

type SimpleWsMsgHandler struct {
	mux    sync.Mutex
	sigs   map[string]Tx
	logger *slog.Logger
}

func NewSimpleWsMsgHandler(logger *slog.Logger) *SimpleWsMsgHandler {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &SimpleWsMsgHandler{
		sigs:   make(map[string]Tx),
		logger: logger,
	}
}

func (h *SimpleWsMsgHandler) Handle(msg []byte) {
	var tx Tx
	err := json.Unmarshal(msg, &tx)
	if err != nil {
		h.logger.Error("ws message", "error", "Failed to unmarshal message", "error", err)
		return
	}
	h.mux.Lock()
	defer h.mux.Unlock()
	for _, sig := range tx.Signatures {
		if t, ok := h.sigs[sig]; ok {
			fmt.Printf("%+v\n", tx)
			fmt.Printf("%+v\n", t)
			panic(1)
		}
		h.sigs[sig] = tx
	}
}
