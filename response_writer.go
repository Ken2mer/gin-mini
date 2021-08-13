package gin

import "net/http"

// ResponseWriter ...
type ResponseWriter interface {
	http.ResponseWriter

	// Returns the HTTP response status code of the current request.
	// Status() int

	// Returns the number of bytes already written into the response http body.
	// Size() int
}
