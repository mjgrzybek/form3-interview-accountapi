package internal

import (
	"bytes"
	"encoding/json"
	"net/url"
	"path"
)

func Encode(data any) (*bytes.Buffer, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}

func SetParam(url *url.URL, key, value string) {
	values := url.Query()
	values.Set(key, value)
	url.RawQuery = values.Encode()
}

// Once Go1.19 is released, replace it with https://pkg.go.dev/net/url@master#JoinPath
// THIS IMPLEMENTATION COVERS ONLY CASES NEEDED FOR CURRENT NEEDS, SEE UNIT TESTS
func JoinPathUrl(url url.URL, elems ...string) *url.URL {
	if len(elems) > 0 {
		url.Path = path.Join(append([]string{url.Path}, append([]string{"/"}, elems...)...)...)
	}
	return &url
}
