package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

//Handler is the method to be invoked
func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
		OK:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
