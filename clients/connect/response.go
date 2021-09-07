package connect

type ApiResponse struct {
	HttpResponse *HttpResponse `json:"http_response"`
	WasRequested bool          `json:"was_requested"`
	WasCached    bool          `json:"was_cached"`
}

type HttpResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}
