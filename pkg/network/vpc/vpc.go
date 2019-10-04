package vpc

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"sire.run/awsc/pkg/op"
)

type service struct {
	ec2Client *ec2.Client
}

func (srv *service) create(o *op.Op) *op.Op {
	var input = &ec2.CreateVpcInput{}
	o.Input = input

	var sender op.Sender = func(ctx context.Context) (response interface{}, err error) {
		return srv.ec2Client.CreateVpcRequest(input).Send(ctx)
	}
	o.Sender = sender
	return o
}

func Create(config aws.Config) func(*op.Op) *op.Op {
	return (&service{ec2.New(config)}).create
}
