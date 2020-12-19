package mockserver

type mockServer struct {

}

type mockServerConfig struct {

}

// NewMockServer Creates new mock http server based on supplied
// JSON config
func NewMockServer(config string) *mockServer {
	return &mockServer{}
}

func parseConfig(config string) *mockServerConfig {
	return &mockServerConfig{}
}