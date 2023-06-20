package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Cognito() *cognito.CognitoIdentityProvider {
	region := os.Getenv("AWS_REGION")
	mySession := session.Must(session.NewSession())
	return cognito.New(mySession, aws.NewConfig().WithRegion(region))
}
