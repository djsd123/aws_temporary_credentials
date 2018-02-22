package main

import (
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"fmt"
	"bufio"
	"os"
)

func main() {

	var (
		serialNumber = "Device-ARN"
	)


	mfa_code := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter MFA Code: ")
	mfa_code.Scan()

	svc := sts.New(session.New())

	input := &sts.GetSessionTokenInput {
		DurationSeconds: aws.Int64(900),
		SerialNumber:    aws.String(serialNumber),
		TokenCode:       aws.String(mfa_code.Text()),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}