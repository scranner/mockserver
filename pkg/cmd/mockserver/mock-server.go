package mockserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scranner/mockserver/internal/config"
	"github.com/scranner/mockserver/internal/route"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Server Object
type Server struct {
	Router *mux.Router
	Config *config.Config
}

// NewMockServer Creates new mock http server based on supplied
// JSON config
func NewMockServer() Server {

	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()

	serverConfig, routes := config.LoadConfigFromEnv()

	logger := logrus.New()
	logger.SetLevel(serverConfig.LogLevel)

	addProbeHandler("/live", router, logger)
	addProbeHandler("/ready", router, logger)

	for _, suppliedRoute := range routes {
		suppliedRoute.CreateHandler(router, logger)
	}

	return Server{
		router,
		&serverConfig,
	}
}

// StartServer will start the http server
func (s *Server) StartServer() {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("0.0.0.0:%s", s.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func addProbeHandler(path string, r *mux.Router, logger *logrus.Logger) {
	route.ProcessedRoute{
		Path: path,
		Description: route.ProcessedDescription{
			Response:     route.ProcessedResponse{
				Key:               "root",
				Untemplated:       map[string]interface{}{ "Status": "OK" },
				Templated:         nil,
				UntemplatedArries: nil,
				Order:             nil,
				RepeatConfig:      route.RepeatConfig{
					Repeat: false,
					Min:    1,
					Max:    1,
				},
			},
			ResponseCode: 200,
			Method:       "GET",
		},
	}.CreateHandler(r, logger)
}