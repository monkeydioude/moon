package moon

import (
	"log"
	"net/http"
	"time"
)

// ServerRunWithConfig starts http server using http.Server configuration
func ServerRunWithConfig(server *http.Server) error {
	log.Printf("[INFO] Server running: %+v", server)
	return server.ListenAndServe()
}

// ServerRunWithParams wraps ServerRunWithConfig() using configuration built over given parameters
func ServerRunWithParams(addr string, handler *Handler, readTimeout, writeTimeout time.Duration, maxHeaderBytes int) error {
	return ServerRunWithConfig(&http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	})
}

// ServerRun wraps ServerRunWithConfig() using defined default configuration
func ServerRun(addr string, handler *Handler) error {
	return ServerRunWithConfig(&http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})
}
