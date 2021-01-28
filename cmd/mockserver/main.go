package main

import (
	"github.com/scranner/mockserver/pkg/cmd/mockserver"
)

func main() {
	MockServer := mockserver.NewMockServer()
	MockServer.StartServer()
}
