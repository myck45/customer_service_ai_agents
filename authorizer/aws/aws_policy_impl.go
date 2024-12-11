package aws

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

type AWSPolicyImpl struct{}

// GeneratePolicy implements AWSPolicy.
func (a *AWSPolicyImpl) GeneratePolicy(principalID string, effect string, resource string, claims map[string]interface{}) events.APIGatewayCustomAuthorizerResponse {

	if principalID == "" {
		logrus.Error("[GeneratePolicy] principalID is empty")
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: "unathorized",
		}
	}

	if effect == "" {
		effect = "Deny"
	}

	authResponse := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: principalID,
	}

	if resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	authResponse.Context = map[string]interface{}{
		"id":        principalID,
		"auth_time": time.Now().Unix(),
		"email":     claims["email"],
		"role":      claims["role"],
	}

	return authResponse
}

func NewAWSPolicy() AWSPolicy {
	return &AWSPolicyImpl{}
}
