package viewmodels

type ErrorMessage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
