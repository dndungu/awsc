package parser

import (
	"errors"
	"io"
	"os"
)

type Parser struct {
	Input, Output string
	NewDecoder    func(io.Reader) Decoder
	NewEncoder    func(io.Writer) Encoder
	Reader        io.Reader
	Writer        io.Writer
}

func (p *Parser) Decode(v interface{}) error {
	if len(p.Input) > 0 {
		var file, err = os.Open(p.Input)
		if err != nil {
			return err
		}

		defer file.Close()
		p.Reader = file
	}

	if p.Reader == nil {
		return errors.New("Parser.Reader can not be nil")
	}

	if p.NewDecoder == nil {
		return errors.New("Parser.NewDecoder can not be nil")
	}

	return p.NewDecoder(p.Reader).Decode(v)
}

func (p *Parser) Encode(v interface{}) error {
	if len(p.Output) > 0 {
		var file, err = os.Create(p.Output)
		if err != nil {
			return err
		}

		defer file.Close()
		p.Writer = file
	}

	if p.Writer == nil {
		return errors.New("Parser.Writer can not be nil")
	}

	if p.NewEncoder == nil {
		return errors.New("Parser.NewEncoder can not be nil")
	}

	return p.NewEncoder(p.Writer).Encode(v)
}

func (p *Parser) Parse() error {
	return nil
}

func New(options ...Option) *Parser {
	var p Parser
	for _, fn := range options {
		fn(&p)
	}
	return &p
}
