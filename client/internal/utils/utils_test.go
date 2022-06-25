package internal

import (
	"bytes"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				data: "Urok twój, woń, wpływ, czar, aury żar",
			},
			want:    bytes.NewBufferString(`"Urok twój, woń, wpływ, czar, aury żar"`),
			wantErr: false,
		},
		{
			name: "int",
			args: args{
				data: int(2),
			},
			want:    bytes.NewBufferString("2"),
			wantErr: false,
		},
		{
			name: "struct",
			args: args{
				data: struct {
					Field int `json:"field"`
				}{Field: 42},
			},
			want:    bytes.NewBufferString(`{"field":42}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetParam(t *testing.T) {
	type args struct {
		url   *url.URL
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		{
			name: "basic",
			args: args{
				url: func() *url.URL {
					url, _ := url.Parse("http://foo")
					return url
				}(),
				key:   "Key",
				value: "Val",
			},
			want: func() *url.URL {
				url, _ := url.Parse("http://foo?Key=Val")
				return url
			}(),
		},
		{
			name: "add to already existing params",
			args: args{
				url: func() *url.URL {
					url, _ := url.Parse("http://foo?param=asd")
					return url
				}(),
				key:   "Key",
				value: "Val",
			},
			want: func() *url.URL {
				url, _ := url.Parse("http://foo?Key=Val&param=asd")
				return url
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetParam(tt.args.url, tt.args.key, tt.args.value)
			assert.Equal(t, tt.want, tt.args.url)
		})
	}
}
