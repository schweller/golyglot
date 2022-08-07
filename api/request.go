package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method string
	URL    *url.URL
	Path   string
	Host   string
	Params url.Values
	Token  string
	Body   io.Reader
}

func (req *Request) NewRequestHTTP() (*http.Request, error) {
	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("DeepL-Auth-Key %s", req.Token)},
	}

	rawQuery := ""

	if req.Params != nil {
		rawQuery = req.Params.Encode()
	}

	foo := &http.Request{
		Method: req.Method,
		URL: &url.URL{
			Host:     req.Host,
			Scheme:   "https",
			Path:     req.Path,
			RawQuery: rawQuery,
		},
		Header: headers,
	}

	return foo, nil
}
