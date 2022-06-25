package internal

import (
	"bytes"
	"encoding/json"
	"net/url"
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
