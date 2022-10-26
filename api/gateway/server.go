package gateway

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server represents an RPC server on the Node.
// TODO @renaynay: eventually, gateway server should be able to be toggled on and off.
type Server struct {
	srv      *http.Server
	srvMux   *mux.Router // http request multiplexer
	listener net.Listener
}

// NewServer returns a new RPC Server.
func NewServer(address string, port string) *Server {
	srvMux := mux.NewRouter()
	srvMux.Use(setContentType)

	server := &Server{
		srvMux: srvMux,
	}
	server.srv = &http.Server{
		Addr:    address + ":" + port,
		Handler: server,
		// the amount of time allowed to read request headers. set to the default 2 seconds
		ReadHeaderTimeout: 2 * time.Second,
	}
	return server
}

// Start starts the gateway Server, listening on the given address.
func (s *Server) Start(context.Context) error {
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.listener = listener
	log.Infow("Gateway server started", "listening on", listener.Addr().String())
	//nolint:errcheck
	go s.srv.Serve(listener)
	return nil
}

// Stop stops the gateway Server.
func (s *Server) Stop(context.Context) error {
	// if server already stopped, return
	if s.listener == nil {
		return nil
	}
	if err := s.listener.Close(); err != nil {
		return err
	}
	s.listener = nil
	log.Info("Gateway server stopped")
	return nil
}

// RegisterMiddleware allows to register a custom middleware that will be called before http.Request will reach handler.
func (s *Server) RegisterMiddleware(m mux.MiddlewareFunc) {
	// `router.Use` appends new middleware to existing
	s.srvMux.Use(m)
}

// RegisterHandlerFunc registers the given http.HandlerFunc on the Server's multiplexer
// on the given pattern.
func (s *Server) RegisterHandlerFunc(pattern string, handlerFunc http.HandlerFunc, method string) {
	s.srvMux.HandleFunc(pattern, handlerFunc).Methods(method)
}

// ServeHTTP serves inbound requests on the Server.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.srvMux.ServeHTTP(w, r)
}

// ListenAddr returns the listen address of the server.
func (s *Server) ListenAddr() string {
	return s.listener.Addr().String()
}
