package main

import (
	"fmt"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cmd := exec.Command("/opt/headless-chromium/headless-chromium", "--headless", "--disable-dev-shm-usage", "--ignore-certificate-errors", "--no-sandbox", "--single-process", "--disable-gpu", "--dump-dom", "https://www.chromestatus.com/")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Hello, %v", string(output)),
			StatusCode: 400,
		}, nil
	}

	fmt.Println(output)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(output)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
