package codec

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const (
	Accept = "Accept"
)

func NewEncoder(w http.ResponseWriter) Encoder {
	return NewJSONEncoder(w)
}

func NewDecoder(r *http.Request) Decoder {
	return NewJSONDecoder(r.Body)
}

type Encoder interface {
	Encode(v interface{}) error
}

type Decoder interface {
	Decode(v interface{}) error
}

func NewJSONEncoder(w io.Writer) Encoder {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.NewEncoder(w)
}

func NewJSONDecoder(r io.Reader) Decoder {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.NewDecoder(r)
}
