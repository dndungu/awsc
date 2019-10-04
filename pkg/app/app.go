package app

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"sire.run/awsc/pkg/network/vpc"
	"sire.run/awsc/pkg/op"
	"sire.run/awsc/pkg/parser"
)

type App struct {
	Args          []string
	Input, Output string
}

func (app App) Do(ctx context.Context) error {
	var config, err = external.LoadDefaultAWSConfig()

	if err != nil {
		return err
	}
	var resource, verb string
	if len(app.Args) != 3 {
		return fmt.Errorf("expected 3 arguments, got %d (%v)", len(app.Args), app.Args)
	}
	verb = app.Args[1]
	resource = app.Args[2]
	var name = verb + "." + resource
	var p = parser.New(parser.WithInput(app.Input), parser.WithOutput(app.Output))
	return op.New(
		op.WithConfig(config),
		op.WithDecorator(decorator(name)),
		op.WithParser(p),
	).Do(ctx)
}

func decorator(name string) func(aws.Config) func(*op.Op) *op.Op {
	switch name {
	case "create.vpc":
		return vpc.Create
	default:
		return noop
	}
}
