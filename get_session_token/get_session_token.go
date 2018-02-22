package get_session_token

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sess = session.Must(session.NewSession())
)
func GetSessionToken() credentials.Value {

	creds := ec2rolecreds.NewCredentials(sess)
	creds.Expire()
	credsValue, err := creds.Get()
	if err != nil {
		panic(err)
	}

	return credsValue
}
