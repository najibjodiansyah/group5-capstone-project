package util

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func GetAWSSession() (*session.Session, error) {
	if sess == nil {
		newSession, err := session.NewSession(
			&aws.Config{
				Region: aws.String(os.Getenv("AWS_REGION")),
				Credentials: credentials.NewStaticCredentials(
					os.Getenv("AWS_ACCESSKEYID"),
					os.Getenv("AWS_SECRETKEY"),
					"",
				),
			},
		)

		if err != nil {
			return nil, err
		}

		sess = newSession
	}

	return sess, nil
}
