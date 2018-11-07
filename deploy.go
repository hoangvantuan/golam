package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/urfave/cli"
)

// Deploy is deploy lambda@edge to cloudfront
func Deploy(ctx *cli.Context) error {
	flags := (&Flags{}).Scan(ctx)

	loadAWSConfig(flags.Region)

	if flags.IsSetupFromVersion() {
		err := updateCloudFront(flags.Distribution, flags.PathPattern, flags.EventType, flags.Name, flags.Version)

		if err != nil {
			log.Fatalf("deploy function failure %s", err)
		}

		return err
	}

	if flags.IsSetupFromSourceCode() {
		cfg, err := updateFunctionCode(flags.Name, flags.Path, true)

		if err != nil {
			log.Fatalf("update function code failure %s", err)
			return err
		}

		err = updateCloudFront(flags.Distribution, flags.PathPattern, flags.EventType, flags.Name, *cfg.Version)

		if err != nil {
			log.Fatalf("deploy function failure %s", err)
		}

		return err
	}

	log.Fatalln("command not found")

	return fmt.Errorf("command not found")

}

func updateCloudFront(distributionID, pathPattern, eventType, lambdaFnName, version string) error {
	cloudFrontSvc := cloudfront.New(sess)
	lambdaSvc := lambda.New(sess)

	lambdaFn, err := lambdaSvc.GetFunction(&lambda.GetFunctionInput{FunctionName: aws.String(lambdaFnName)})

	if err != nil {
		return err
	}

	log.Printf("[info] starting setting  lambda %s with version %s", lambdaFnName, version)

	distributionCfg, err := cloudFrontSvc.GetDistributionConfig(&cloudfront.GetDistributionConfigInput{Id: aws.String(distributionID)})

	for _, i := range distributionCfg.DistributionConfig.CacheBehaviors.Items {
		if *i.PathPattern == pathPattern {
			for _, fn := range i.LambdaFunctionAssociations.Items {

				ldarn := strings.Split(*fn.LambdaFunctionARN, ":")

				currentVer := ldarn[len(ldarn)-1]

				curLbfn := strings.Join(ldarn[:len(ldarn)-1], ":")

				if curLbfn != *lambdaFn.Configuration.FunctionArn {
					continue
				}

				if currentVer == version {
					return fmt.Errorf("[warn] current version is %s no need update", version)
				}

				log.Printf("[info] starting update config version %s to %s", currentVer, version)
				fn.SetLambdaFunctionARN(fmt.Sprintf("%s:%s", curLbfn, version))
			}
		}
	}

	_, err = cloudFrontSvc.UpdateDistribution(&cloudfront.UpdateDistributionInput{DistributionConfig: distributionCfg.DistributionConfig, Id: aws.String(distributionID), IfMatch: distributionCfg.ETag})

	if err != nil {
		return err
	}

	log.Println("[info] setting lambda@edge with cloudfront successfull")

	return nil
}
