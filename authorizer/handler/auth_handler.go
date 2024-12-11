package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type AuthHandler interface {
	HandleAuth(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error)
}
