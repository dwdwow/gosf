package gosf

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

// SimpleWSServer represents a WebSocket server
type SimpleWSServer struct {
	// clients holds all connected clients
	clients map[*websocket.Conn]bool
	// broadcast is a channel for sending messages to all clients
	broadcast chan []byte
	// mutex for protecting clients map
	mu sync.RWMutex
	// upgrader for upgrading HTTP connections to WebSocket connections
	upgrader websocket.Upgrader
	// port for the server
	port string

	// ctx and cancel for the server
	ctx    context.Context
	cancel context.CancelFunc

	// logger for logging
	logger *slog.Logger
}

// NewSimpleWSServer creates a new WebSocket server
func NewSimpleWSServer(port string, logger *slog.Logger) *SimpleWSServer {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SimpleWSServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte, 10000),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins for now
			},
		},
		port:   port,
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

// Serve starts the WebSocket server on the specified port
func (s *SimpleWSServer) Serve(path string) {
	// Start the broadcast handler
	go s.handleBroadcast()

	// Set up the HTTP handler for WebSocket connections
	http.HandleFunc(path, s.handleConnections)

	// Start the server
	s.logger.Info("WebSocket server starting on port", "port", s.port)

	go func() {
		err := http.ListenAndServe(":"+s.port, nil)
		if err != nil {
			s.logger.Error("ws server", "error", "Failed to listen and serve", "error", err)
		}
		s.cancel()
		panic(1)
	}()
}

// handleConnections handles new WebSocket connections
func (s *SimpleWSServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("ws connection request", "method", r.Method, "ip", r.RemoteAddr)
	// Upgrade the HTTP connection to a WebSocket connection
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register new client
	s.mu.Lock()
	s.clients[ws] = true
	s.mu.Unlock()

	// Remove client when function returns
	defer func() {
		s.mu.Lock()
		delete(s.clients, ws)
		s.mu.Unlock()
	}()

	// Keep the connection alive
	for {
		// Read any incoming messages (not used in this simple implementation)
		_, b, err := ws.ReadMessage()
		if err != nil {
			s.logger.Error("ws connection request", "method", r.Method, "ip", r.RemoteAddr, "error", "Failed to read message", "error", err)
			break
		}
		s.logger.Info("ws message", "message", string(b))
	}
}

// handleBroadcast handles broadcasting messages to all connected clients
func (s *SimpleWSServer) handleBroadcast() {
	for {
		select {
		case <-s.ctx.Done():
			s.logger.Info("ws broadcast", "error", "Context done")
			return
		case msg := <-s.broadcast:
			// Send it to every client
			s.mu.RLock()
			for client := range s.clients {
				client := client
				go func() {
					err := client.WriteMessage(websocket.TextMessage, msg)
					if err != nil {
						s.logger.Error("ws broadcast", "error", "Failed to write message", "error", err)
						client.Close()
						delete(s.clients, client)
					}
				}()
			}
			s.mu.RUnlock()
		}
	}
}

// Broadcast sends a message to all connected clients
func (s *SimpleWSServer) Broadcast(msg []byte) {
	select {
	case s.broadcast <- msg:
	default:
		s.logger.Error("ws broadcast", "error", "Channel full")
	}
}
