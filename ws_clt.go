package gosf

import (
	"log/slog"
	"os"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	ul     string
	ws     *websocket.Conn
	logger *slog.Logger
}

func NewWsClient(url string, logger *slog.Logger) *WsClient {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	c := &WsClient{
		ul:     url,
		logger: logger,
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
			c.logger.Info("ws message", "message", string(b))
		case websocket.BinaryMessage:
			c.logger.Info("ws binary message", "message", string(b))
		case websocket.CloseMessage:
			c.logger.Info("ws close message")
		}
	}
}
