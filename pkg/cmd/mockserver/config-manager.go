package mockserver

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	"strings"
// )

// func loadConfigFromEnv(unparsedRoutes string) (Config, []Route) {

// 	var routes map[string]RouteDescription
// 	err := json.Unmarshal([]byte(unparsedRoutes), &routes)

// 	if err != nil {
// 		panic(fmt.Sprintf("Invalid JSON config provided. [%s]", err))
// 	}

// 	parsedRoutes := make([]Route, 0)

// 	for path, routeDesc := range routes {
// 		for key, data := range routeDesc.Response {
// 			if strings.HasPrefix(string(data), "[") && strings.HasSuffix(data, "]") {
// 				append(routeDesc.ProcessedResponse.Templated, data)
// 			}
// 		}
// 		parsedRoutes = append(parsedRoutes, Route{
// 			path,
// 			routeDesc,
// 		})
// 	}

// 	return Config{
// 			"8080",
// 		},
// 		parsedRoutes
// }

// // CreateHandler handle it
// func (r Route) CreateHandler(router *mux.Router) {
// 	handleFunc := func(w http.ResponseWriter, r *http.Request) {
// 		responseToSend := make([]map[string]interface{}, 0)

// 		if route.Description.RepeatResponse {
// 			for i := 0; i < rand.Intn(route.Description.MaxRepeats-route.Description.MinRepeats+1)+route.Description.MinRepeats; i++ {
// 				responseToSend = append(responseToSend, route.Description.Response)
// 			}
// 		} else {
// 			responseToSend = append(responseToSend, route.Description.Response)
// 		}

// 		for i := 0; i < len(responseToSend); i++ {
// 			for j := 0; j < len(responseToSend[i]); i++ {
// 				repo := responseToSend[i][j].(func() string)()
// 				fmt.Println(repo)
// 			}
// 		}

// 		// for i, response := range responseToSend {
// 		// 	for j, k := range response {
// 		// 		// responseToSend[i] = string(test1.(func() string)())
// 		// 		repo := responseToSend[i][j].(func() string)()
// 		// 		fmt.Println(k)
// 		// 		responseToSend[i][j] = repo
// 		// 	}
// 		// }

// 		b, _ := json.Marshal(responseToSend)
// 		w.WriteHeader(route.Description.ResponseCode)
// 		fmt.Fprintf(w, string(b))
// 	}
// 	router.HandleFunc(route.Path, handleFunc)
// }
