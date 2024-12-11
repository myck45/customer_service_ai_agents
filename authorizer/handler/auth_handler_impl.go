package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/proyectos01-a/authorizer/aws"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

type AuthHandlerImpl struct {
	policy       aws.AWSPolicy
	jwtValidator utils.AuthUtils
}

// HandleAuth implements AuthHandler.
func (a *AuthHandlerImpl) HandleAuth(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := strings.TrimPrefix(event.AuthorizationToken, "Bearer ")

	claims, err := a.jwtValidator.ValidateToken(token)
	if err != nil {
		logrus.Error("[HandleAuth] Error validating token: ", err)
		return a.policy.GeneratePolicy("", "Deny", event.MethodArn, nil), fmt.Errorf("token validation failed: %v", err)
	}

	principalID, ok := claims["id"].(float64)
	if !ok {
		logrus.Error("[HandleAuth] Invalid principal ID")
		return a.policy.GeneratePolicy("", "Deny", event.MethodArn, nil), fmt.Errorf("invalid principal ID")
	}

	uintPrincipalID := uint(principalID)

	email, ok := claims["email"].(string)
	if !ok {
		logrus.Error("[HandleAuth] Invalid email")
		return a.policy.GeneratePolicy("", "Deny", event.MethodArn, nil), fmt.Errorf("invalid email")
	}

	role, ok := claims["role"].(string)
	if !ok {
		logrus.Error("[HandleAuth] Invalid role")
		return a.policy.GeneratePolicy("", "Deny", event.MethodArn, nil), fmt.Errorf("invalid role")
	}

	tokenClaims := map[string]any{
		"email": email,
		"role":  role,
	}

	logrus.Infof("User authenticated: ID=%v, Email=%s, Role=%s", uintPrincipalID, email, role)

	// Allow access and pass the claims to the backend
	return a.policy.GeneratePolicy(fmt.Sprintf("%d", uintPrincipalID), "Allow", event.MethodArn, tokenClaims), nil
}

func NewAuthHandler(policy aws.AWSPolicy, jwtValidator utils.AuthUtils) AuthHandler {
	return &AuthHandlerImpl{
		policy:       policy,
		jwtValidator: jwtValidator,
	}
}
