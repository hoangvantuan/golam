package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Auto deploy lambda@edge function cmd"
	app.Version = "0.0.1"

	// register flag
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Usage: "name of lambda function",
		},
		cli.StringFlag{
			Name:  "distribution, d",
			Usage: "cloudfront distribution id",
		},
		cli.StringFlag{
			Name:  "path, p",
			Usage: "lambda source code path directory",
		},
		cli.StringFlag{
			Name:  "lambda-version, lv",
			Usage: "version of lambda function",
		},
		cli.StringFlag{
			Name:  "path-pattern, pt",
			Usage: "path pattern of cloudfront distribution",
		},
		cli.StringFlag{
			Name:  "even-type, et",
			Usage: "event type of cloudfront distribution",
		},
		cli.StringFlag{
			Name:  "region, r",
			Usage: "aws region",
		},
		cli.BoolFlag{
			Name:  "publish-new-version, pnv",
			Usage: "event type of cloudfront distribution",
		},
	}

	// register command
	app.Commands = []cli.Command{
		{
			Name:  "deploy",
			Usage: "publish new lambda@edge version and connect to cloudfront distribution",
			Action: func(c *cli.Context) error {
				return Deploy(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of lambda function",
				},
				cli.StringFlag{
					Name:  "distribution, d",
					Usage: "cloudfront distribution id",
				},
				cli.StringFlag{
					Name:  "path, p",
					Usage: "lambda source code path directory",
				},
				cli.StringFlag{
					Name:  "lambda-version, lv",
					Usage: "version of lambda function",
				},
				cli.StringFlag{
					Name:  "path-pattern, pt",
					Usage: "path pattern of cloudfront distribution",
				},
				cli.StringFlag{
					Name:  "event-type, et",
					Usage: "event type of cloudfront distribution",
				},
				cli.StringFlag{
					Name:  "region, r",
					Usage: "aws region",
				},
			},
		},
		{
			Name:  "update",
			Usage: "update lambda function, publish new version from source",
			Action: func(c *cli.Context) error {
				return Update(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of lambda function",
				},
				cli.StringFlag{
					Name:  "path, p",
					Usage: "lambda source code path directory",
				},
				cli.StringFlag{
					Name:  "region, r",
					Usage: "aws region",
				},
				cli.BoolFlag{
					Name:  "publish-new-version, pnv",
					Usage: "event type of cloudfront distribution",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
