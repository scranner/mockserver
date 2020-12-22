package mockserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/scranner/mock-server/internal/route"
)

// Server Object
type Server struct {
	Router *mux.Router
	Config *Config
}

// ProcessedResponse we
type ProcessedResponse struct {
	Untemplated map[string]string
	Templated   map[string]func() string
	Order       []string
}

// Config Server Config
type Config struct {
	Port string
}


// NewMockServer Creates new mock http server based on supplied
// JSON config
func NewMockServer(suppliedConfig string) Server {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()

	Route{
		"/live",
		RouteDescription{
			Response{
				map[string]string{"Status": "OK"},
				200,
				false,
				0,
				0,
			},
		},
	}.CreateHandler(router)

	// &Route{
	// 	"/ready",
	// 	RouteDescription{
	// 		Response{map[string]string{
	// 			"Status", "OK",
	// 		},
	// 			200,
	// 			false,
	// 			0,
	// 			0,
	// 		},
	// 	},
	// }.CreateHandler(router)

	// config, routes := loadConfigFromEnv(suppliedConfig)

	// for _, route := range routes {
	// 	createHandler(router, route)
	// }

	return Server{
		router,
		&Config{
			"8080",
		},
	}
}

// BuildResponse b
type (r Response) BuildResponse() {

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

func probeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

// CreateHandler handle it
func (r Route) CreateHandler(router *mux.Router) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {

		// Create final response object
		// generate tmeplated valu

		responseToSend := r.Response

		// if route.Description.RepeatResponse {
		// 	for i := 0; i < rand.Intn(route.Description.MaxRepeats-route.Description.MinRepeats+1)+route.Description.MinRepeats; i++ {
		// 		responseToSend = append(responseToSend, route.Description.Response)
		// 	}
		// } else {
		// 	responseToSend = append(responseToSend, route.Description.Response)
		// }

		// for i := 0; i < len(responseToSend); i++ {
		// 	for j := 0; j < len(responseToSend[i]); i++ {
		// 		repo := responseToSend[i][j].(func() string)()
		// 		fmt.Println(repo)
		// 	}
		// }

		// for i, response := range responseToSend {
		// 	for j, k := range response {
		// 		// responseToSend[i] = string(test1.(func() string)())
		// 		repo := responseToSend[i][j].(func() string)()
		// 		fmt.Println(k)
		// 		responseToSend[i][j] = repo
		// 	}
		// }

		b, _ := json.Marshal(responseToSend)
		w.WriteHeader(route.Description.ResponseCode)
		fmt.Fprintf(w, string(b))
	}
	router.HandleFunc(route.Path, handleFunc)
}

