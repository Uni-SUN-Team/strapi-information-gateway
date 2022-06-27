package ports

import "net/http"

type HTTPService interface {
	HTTPRequest(url string, method string, payload []byte) (*http.Response, error)
}
