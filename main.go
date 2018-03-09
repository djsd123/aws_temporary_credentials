package main

import (
	"fmt"
	"github.com/djsd123/aws_temporary_credentials/get_session_token"
	"os"
	"log"
	"html/template"
	"path/filepath"
)


func main() {

	creds := get_session_token.GetSessionToken()
	home := os.Getenv("HOME")
	awsDirectory := ".aws"
	path := fmt.Sprintf("%v" + string(filepath.Separator) + "%v", home, awsDirectory)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	filePath := path + "/config"

	const awsConfig = `
[default]
region=eu-west-1
aws_access_key_id={{ .AccessKeyID }}
aws_secret_access_key={{ .SecretAccessKey }}
aws_session_token={{ .SessionToken }}
`

	templ, err := template.New("config").Parse(awsConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	profile, err := os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil{
		log.Fatalf("Failed to open file: %s", err)
		panic(err)
	}


	err = templ.Execute(profile, creds)
	if err != nil {
		log.Fatal(err)
	}

	defer profile.Close()

}
