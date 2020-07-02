package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	profile string
	region  string
)

func init() {
	flag.StringVar(&profile, "p", "default", "aws shared credential profile name")
	flag.StringVar(&region, "r", "ap-northeast-1", "aws region")
	flag.Parse()
}

func main() {
	sess := session.Must(session.NewSession())
	cred := credentials.NewSharedCredentials("", profile)
	svc := ec2.New(sess, &aws.Config{Credentials: cred,
		Region: aws.String(region)})
	res, _ := svc.DescribeInstances(nil)

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			var instanceName string
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					instanceName = *t.Value
				}
			}
			fmt.Printf("%s,%s,%s,%s,%s\n",
				*i.InstanceId,
				instanceName,
				*i.InstanceType,
				*i.Placement.AvailabilityZone,
				*i.State.Name,
			)
		}
	}
}
