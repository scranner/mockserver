package mockserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

// Server Object
type Server struct {
	Router *mux.Router
	Config *Config
}

// Config Server Config
type Config struct {
	Port string
}

// NewMockServer Creates new mock http server based on supplied
// JSON config
func NewMockServer(suppliedConfig string) Server {
	router := mux.NewRouter()
	router.HandleFunc("/live", probeHandler)
	router.HandleFunc("/ready", probeHandler)

	return Server{
		router,
		parseConfig(suppliedConfig),
	}
}

// StartServer will start the http server
func (s *Server) StartServer() {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", s.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func parseConfig(unparsedConfig string) *Config {
	return &Config{
		"8080",
	}
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
