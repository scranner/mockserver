package config

import (
	"encoding/json"
	"fmt"
	"github.com/scranner/mockserver/internal/route"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Config struct {
	Port string
	LogLevel logrus.Level
}

func LoadConfigFromEnv() (Config, []route.ProcessedRoute) {

	jsonFile, fileErr := os.Open("/opt/routes.json"); if fileErr != nil {
		panic(fileErr)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var routes map[string]route.Description
	err := json.Unmarshal(byteValue, &routes)

	if err != nil {
		panic(fmt.Sprintf("Invalid JSON config provided. [%s]", err))
	}

	parsedRoutes := make([]route.ProcessedRoute, 0)

	for path, routeDesc := range routes {

		processedRoute := routeDesc.ProcessRouteDescription()

		parsedRoutes = append(parsedRoutes, route.ProcessedRoute{
			Path:        path,
			Description: processedRoute,
		})
	}

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))

	if err != nil {
		panic(err)
	}

	return Config{
			os.Getenv("PORT"),
		logLevel,
		},
		parsedRoutes
}