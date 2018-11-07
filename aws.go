package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func loadAWSConfig(region string) {
	if region != "" {
		sess = session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	} else {
		sess = session.Must(session.NewSession())
	}
}
