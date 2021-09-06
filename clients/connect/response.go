package connect

import "net/http"

type ApiResponse struct {
	HttpResponse *http.Response
	WasRequested bool
}
