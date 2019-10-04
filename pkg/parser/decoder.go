package parser

type Decoder interface {
	Decode(interface{}) error
}
