package route

// Route bla
type Route struct {
	Path        string
	Description Description
}

// Description thing
type Description struct {
	Response       Response
	ResponseCode   int
	RepeatResponse bool
	MinRepeats     int
	MaxRepeats     int
}

// Response obj
type Response struct {
	Untemplated map[string]string
	Templated   map[string]string
	Order       []string
}
