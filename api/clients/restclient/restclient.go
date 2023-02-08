package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// func Somethng(){
// 	Post("https://api.github.com/user/repos",`{"name":"Github API"}`)
// }

var (
	enablemocks = false
	mocks       = make(map[string]*Mock)
)

func StartMockups() {
	enablemocks = true
}

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func StopMockups() {
	enablemocks = false
}

func GetId(httpmethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpmethod, url)
}

func AddMockup(mock Mock) {
	mocks[GetId(mock.HttpMethod, mock.Url)] = &mock
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enablemocks {
		mock := mocks[GetId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found")
		}
		return mock.Response, mock.Err
	}
	
	jsonbytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonbytes))
	request.Header = headers
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	return client.Do(request)
}
