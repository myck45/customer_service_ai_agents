package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/proyectos01-a/authorizer/aws"
	"github.com/proyectos01-a/authorizer/handler"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting authorizer...")
	jwtValidator := utils.NewAuthUtilsImpl()
	policy := aws.NewAWSPolicy()
	authHandler := handler.NewAuthHandler(policy, jwtValidator)

	lambda.Start(authHandler.HandleAuth)

	logrus.Info("Authorizer started")
}
