package parser

type Encoder interface {
	Encode(interface{}) error
}
