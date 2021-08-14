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

type responseWriter struct {
	http.ResponseWriter
	// size   int
	// status int
}

var _ ResponseWriter = &responseWriter{}

func (w *responseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	// w.size = noWritten
	// w.status = defaultStatus
}

// func (w *responseWriter) Status() int {
// 	return w.status
// }

// func (w *responseWriter) Size() int {
// 	return w.size
// }
