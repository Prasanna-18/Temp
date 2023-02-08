package github

type GithubErrorResponse struct {
	StatusCode int `json:"statuscode"`
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentationUrl"`
	Errors           []GithubError `json:"errors"`
}

type GithubError struct {
	Resorce string `json:"resorce"`
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
