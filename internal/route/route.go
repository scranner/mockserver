package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scranner/mockserver/internal/templating"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

type RepeatConfig struct {
	Repeat bool
	Min    int
	Max    int
}
// Route bla
type Route struct {
	Path        string
	Description Description
}

type ProcessedRoute struct {
	Path        string
	Description ProcessedDescription
}

// Description thing
type Description struct {
	Response       map[string]interface{}
	ResponseCode   int
	Method string
	RepeatConfig RepeatConfig
}

type ProcessedDescription struct {
	Response       ProcessedResponse
	ResponseCode   int
	Method         string
}

// Repeat obj
type Response struct {
	Untemplated map[string]string
	Templated   map[string]string
	Order       []string
}

// ProcessedResponse we
type ProcessedResponse struct {
	Key			string
	nodes *[]ProcessedResponse
	Untemplated map[string]interface{}
	Templated   map[string]func() string
	UntemplatedArries map[string][]interface{}
	Order       []string
	RepeatConfig RepeatConfig
}

// CreateHandler handle it
func (r ProcessedRoute) CreateHandler(router *mux.Router, logger *logrus.Logger) {
	handleFunc := func(writer http.ResponseWriter, req *http.Request) {
		responseToSend, _ := r.Description.Response.BuildResponse()
		JSONResponse, _ := json.Marshal(responseToSend)
		writer.WriteHeader(r.Description.ResponseCode)
		fmt.Fprintf(writer, string(JSONResponse))
		logger.WithFields(logrus.Fields{
			"Path": r.Path,
			"Method": r.Description.Method,
		}).Info("Request Successful")
	}
	router.HandleFunc(r.Path, handleFunc)
}

func (pr ProcessedResponse) BuildResponse() (map[string]interface{}, RepeatConfig) {
	response := make(map[string]interface{}, 0)

	if pr.Untemplated != nil {
		for key, value := range pr.Untemplated {
			response[key] = value
		}
	}
	if pr.Templated != nil {
		for key, value := range pr.Templated {
			response[key] = value()
		}
	}

	if pr.UntemplatedArries != nil {
		for key, value := range pr.UntemplatedArries {
			response[key] = value
		}
	}

	if pr.nodes != nil {
		if len(*pr.nodes) > 0 {
			for _, node := range *pr.nodes {
				resp, repeat := node.BuildResponse()

				if repeat.Repeat {
					repeatedResp := make([]interface{}, 0)
					rand.Seed(time.Now().UnixNano())
					for i := 0; i < rand.Intn(repeat.Max - repeat.Min + 1) + repeat.Min; i++ {
						repeatedResp = append(repeatedResp, resp)
					}
					response[node.Key] = repeatedResp
				}
			}
		}
	}

	return response, pr.RepeatConfig

}

func (routeDescription Description) ProcessRouteDescription() ProcessedDescription {

	procResp := processResponse(routeDescription.Response, "root")

	return ProcessedDescription{
		Response:       procResp,
		ResponseCode:   routeDescription.ResponseCode,
		Method:         routeDescription.Method,
	}

}

func processResponse(response map[string]interface{}, overkey string) ProcessedResponse {

	nodes := make([]ProcessedResponse, 0)
	untemplatedVariables := make(map[string]interface{}, 0)
	templatedVariables := make(map[string]func() string , 0)
	untemplatedArries := make(map[string][]interface{}, 0)
	order := make([]string, 0)

	var repeat bool
	var minRepeats, maxRepeats int

	for key, value := range response {

		if key == "[Repeat]" {
			repeat = value.(bool)
			continue
		}

		if key == "[MinRepeats]" {
			minRepeats = int(value.(float64))
			continue
		}

		if key == "[MaxRepeats]" {
			maxRepeats = int(value.(float64))
			continue
		}

		order = append(order, key)
		switch value.(type) {
		case map[string]interface{}:
			{
				nodes = append(nodes, processResponse(value.(map[string]interface{}), key))
			}
		case interface{}:
			{
				strValue, isStr := value.(string); if isStr {
				if templating.IsTemplateValue(strValue) {
					templatedVariables[key] = templating.ProcessTemplateVariable(strValue)
					continue
				}
			}
				untemplatedVariables[key] = value.(interface{})
				continue

			}
		case []interface{}:
			{
				untemplatedArries[key] = value.([]interface{})
				continue
			}
		}
	}

	return ProcessedResponse{
		Key: 			overkey,
		nodes: 			&nodes,
		Untemplated: 	untemplatedVariables,
		Templated: 		templatedVariables,
		UntemplatedArries:untemplatedArries,
		Order: 		 	order,
		RepeatConfig: RepeatConfig{
			Repeat: repeat,
			Min:    minRepeats,
			Max:    maxRepeats,
		},
	}
}
