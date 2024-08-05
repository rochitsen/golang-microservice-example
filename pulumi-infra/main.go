package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// AWS ElasticBeanstalk resource
		// Application
		eBEnvironment, err := elasticbeanstalk.NewApplication(ctx, "dev", &elasticbeanstalk.ApplicationArgs{
			Name:        pulumi.String("Demo-service"),
			Description: pulumi.String("A Golang service"),
		})

		if err != nil {
			return err
		}

		// Environment
		_, err = elasticbeanstalk.NewEnvironment(ctx, "Demo-Serice-Env", &elasticbeanstalk.EnvironmentArgs{
			Name:              pulumi.String("Demo-Service"),
			Application:       eBEnvironment.Name,
			SolutionStackName: pulumi.String("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4"),
		})

		if err != nil {
			return err
		}

		return nil
	})
}
