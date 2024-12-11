package aws

import "github.com/aws/aws-lambda-go/events"

type AWSPolicy interface {
    GeneratePolicy(principalID string, effect string, resource string, claims map[string]interface{}) events.APIGatewayCustomAuthorizerResponse
}
