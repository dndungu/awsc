package app

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"sire.run/awsc/pkg/op"
)

type noopInput struct{}

func (noopInput) Validate() error { return nil }

func noop(aws.Config) func(*op.Op) *op.Op {
	return func(o *op.Op) *op.Op {
		o.Input = noopInput{}
		var sender op.Sender = func(context.Context) (interface{}, error) {
			type response struct {
				Foo, Bar string
			}
			var v response
			return &v, nil
		}
		o.Sender = sender
		return o
	}
}
