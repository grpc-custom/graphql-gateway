package codec

import (
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	Accept  = "Accept"
	MsgPack = "application/x-msgpack"
)

func NewEncoder(w http.ResponseWriter) Encoder {
	accept := w.Header().Get(Accept)
	switch strings.ToLower(accept) {
	case MsgPack:
		return NewMsgPackEncoder(w)
	default:
		return NewJSONEncoder(w)
	}
}

func NewDecoder(r *http.Request) Decoder {
	accept := r.Header.Get(Accept)
	switch strings.ToLower(accept) {
	case MsgPack:
		return NewMsgPackDecoder(r.Body)
	default:
		return NewJSONDecoder(r.Body)
	}
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

func NewMsgPackEncoder(w io.Writer) Encoder {
	return msgpack.NewEncoder(w)
}

func NewMsgPackDecoder(r io.Reader) Decoder {
	return msgpack.NewDecoder(r)
}
