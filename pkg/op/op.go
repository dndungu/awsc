package op

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"sire.run/awsc/pkg/parser"
)

type Input interface {
	Validate() error
}

type Sender func(context.Context) (response interface{}, err error)

func (s Sender) Send(ctx context.Context) (response interface{}, err error) {
	return s(ctx)
}

type Op struct {
	Apply  func(*Op) *Op
	Config aws.Config
	Input  Input
	Parser *parser.Parser
	Sender Sender
}

func (op *Op) Do(ctx context.Context) error {
	var err = op.Parser.Parse()
	if err != nil {
		return err
	}

	op.Apply(op)

	err = op.Parser.Decode(op.Input)
	if err != nil {
		return err
	}

	err = op.Input.Validate()
	if err != nil {
		return err
	}

	if op.Sender == nil {
		return errors.New("Action.Send can not be nil")
	}

	var response interface{}
	response, err = op.Sender.Send(ctx)
	if err != nil {
		return err
	}

	return op.Parser.Encode(response)
}

type Option func(*Op)

func WithConfig(config aws.Config) Option {
	return func(op *Op) {
		op.Config = config
	}
}

func WithDecorator(fn func(aws.Config) func(*Op) *Op) Option {
	return func(op *Op) {
		op.Apply = fn(op.Config)
	}
}

func WithParser(p *parser.Parser) Option {
	return func(op *Op) {
		op.Parser = p
	}
}

func New(options ...Option) *Op {
	var op Op
	for _, fn := range options {
		fn(&op)
	}
	return &op
}
