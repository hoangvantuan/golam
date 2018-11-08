package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/urfave/cli"
)

// Update is update lambda function code
func Update(ctx *cli.Context) error {
	flags := (&Flags{}).Scan(ctx)

	loadAWSConfig(flags.Region)

	if flags.IsUpdateFunctionCode() {
		_, err := updateFunctionCode(flags.Name, flags.Path, flags.IsNewVersion)

		if err != nil {
			log.Printf("deploy function failure %s", err)
		}

		return err
	}

	return fmt.Errorf("command not found")
}

func updateFunctionCode(lambdaFnName, path string, isPublish bool) (*lambda.FunctionConfiguration, error) {
	svc := lambda.New(sess)

	log.Println("[info] starting update and publish function")
	log.Printf("[info] source code %s", path)
	log.Printf("[info] lambda function name: %s", lambdaFnName)

	zipFilePath := filepath.Join(path, "index.zip")

	err := zip(path, zipFilePath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// remove all file after deploy
	defer os.RemoveAll(zipFilePath)

	zipFile, err := ioutil.ReadFile(zipFilePath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(lambdaFnName),
		ZipFile:      zipFile,
		Publish:      aws.Bool(isPublish),
	}

	conf, err := svc.UpdateFunctionCode(input)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	log.Printf("[info] new function was be publish to version %s", *conf.Version)

	return conf, nil
}
