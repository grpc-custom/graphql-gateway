package encoder

type Accept int

const (
	AcceptNone Accept = iota
	AcceptJSON
	AcceptMsgPack
)

type Encoder interface {
	Encode()
}

var (
	_ = (*JSONEncoder)(nil)
)

type JSONEncoder struct {
}

func (e *JSONEncoder) Encode() {

}
