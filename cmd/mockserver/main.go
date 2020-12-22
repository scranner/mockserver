package main

import (
	"github.com/scranner/mockserver/pkg/cmd/mockserver"
	"os"
)

func main() {
	MockServer := mockserver.NewMockServer(os.Getenv("CONFIG"))
	MockServer.StartServer()
}
