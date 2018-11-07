package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/lambda"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/urfave/cli"
)

var sess *session.Session

func init() {
	sess = session.Must(session.NewSession())
}

func setSession(region string) {
	sess = session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
}

func main() {
	app := cli.NewApp()

	app.Name = "Auto deploy lambda@edge function app"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) (err error) {

		if context.Bool("all") {
			fmt.Println("[info] update function code and configure cloudfront")

			ldname := context.Args().Get(0)
			srcPath := context.Args().Get(1)
			distributionID := context.Args().Get(2)
			pathPattern := context.Args().Get(3)
			evenType := context.Args().Get(4)

			region := context.Args().Get(5)
			if region != "" {
				setSession(region)
			}

			cfn, err := updateFunction(ldname, srcPath, true)

			if err != nil {
				log.Printf("[error] update function failure %s\n", err)
				return err
			}

			err = setup(distributionID, pathPattern, evenType, ldname, *cfn.Version)

			if err != nil {
				log.Printf("[error] setup distribution failure %s\n", err)
				return err
			}

			return nil
		}

		if context.Bool("update") {
			region := context.Args().Get(2)
			if region != "" {
				setSession(region)
			}

			_, err = updateFunction(context.Args().Get(0), context.Args().Get(1), true)

			if err != nil {
				log.Printf("[error] update cloudfront failure %s\n", err)
			}

			return err
		}

		if context.Bool("publish") {
			region := context.Args().Get(2)
			if region != "" {
				setSession(region)
			}

			_, err = updateFunction(context.Args().Get(0), context.Args().Get(1), true)

			if err != nil {
				log.Printf("[error] publish cloudfront failure %s\n", err)
			}

			return err
		}

		if context.Bool("setup") {
			region := context.Args().Get(5)
			if region != "" {
				setSession(region)
			}

			ldname := context.Args().Get(3)
			version := context.Args().Get(4)
			distributionID := context.Args().Get(0)
			pathPattern := context.Args().Get(1)
			evenType := context.Args().Get(2)

			err = setup(distributionID, pathPattern, evenType, ldname, version)

			if err != nil {
				log.Printf("[error] publish cloudfront failure %s\n", err)
			}

			return err
		}

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "all, a",
			Usage: `update lambda function and configure with cloudfront
			args[0]: lambda function name
			args[1]: source code path
			args[2]: distribution id
			args[3]: path pattern
			args[4]: event type
			args[5]: aws region`,
		},
		cli.BoolFlag{
			Name: "update, u",
			Usage: `update lambda function
			args[0]: lambda function name
			args[1]: source code path
			args[5]: aws region`,
		},
		cli.BoolFlag{
			Name: "publish, p",
			Usage: `update and publish new version
			args[0]: lambda function name
			args[1]: source code path
			args[5]: aws region`,
		},
		cli.BoolFlag{
			Name: "setup, s",
			Usage: `configure latest version of lambda for cloudfront
			args[0]: distribution id			
			args[1]: path pattern
			args[2]: event type
			args[3]: lambda function name
			args[4]: setting version
			args[5]: aws region`,
		},
	}

	app.Run(os.Args)
}

func updateFunction(ldname, srcPath string, publish bool) (*lambda.FunctionConfiguration, error) {
	svc := lambda.New(sess)

	log.Println("[info] starting update and publish function")
	log.Printf("[info] source code %s", srcPath)
	log.Printf("[info] lambda function name: %s", ldname)

	zip, err := ioutil.ReadFile(srcPath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(ldname),
		ZipFile:      zip,
		Publish:      aws.Bool(publish),
	}

	conf, err := svc.UpdateFunctionCode(input)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	log.Printf("[info] new function was be publish to version %s", *conf.Version)

	return conf, nil
}

func setup(distributionID, pathPattern, eventType, ldname, version string) error {
	svc := cloudfront.New(sess)
	lamd := lambda.New(sess)

	lamfn, err := lamd.GetFunction(&lambda.GetFunctionInput{FunctionName: aws.String(ldname)})

	if err != nil {
		return err
	}

	log.Printf("[info] starting setting  lambda %s with version %s", ldname, version)

	dis, err := svc.GetDistributionConfig(&cloudfront.GetDistributionConfigInput{Id: aws.String(distributionID)})

	for _, i := range dis.DistributionConfig.CacheBehaviors.Items {
		if *i.PathPattern == pathPattern {
			for _, fn := range i.LambdaFunctionAssociations.Items {

				ldarn := strings.Split(*fn.LambdaFunctionARN, ":")

				currentVer := ldarn[len(ldarn)-1]

				curLbfn := strings.Join(ldarn[:len(ldarn)-1], ":")

				if curLbfn != *lamfn.Configuration.FunctionArn {
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

	_, err = svc.UpdateDistribution(&cloudfront.UpdateDistributionInput{DistributionConfig: dis.DistributionConfig, Id: aws.String(distributionID), IfMatch: dis.ETag})

	if err != nil {
		return err
	}

	log.Println("[info] setting lambda@edge with cloudfront successfull")

	return nil
}
