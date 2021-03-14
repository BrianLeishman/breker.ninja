package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-playground/validator/v10"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/BrianLeishman/breker.ninja/assets/go/validators"
)

var ginLambda *ginadapter.GinLambda

var r = gin.Default()

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("email", validators.Email)
		v.RegisterValidation("name", validators.Name)
	}

	if len(os.Getenv("AWS_EXECUTION_ENV")) != 0 {
		gin.SetMode(gin.ReleaseMode)

		ginLambda = ginadapter.New(r)
		lambda.Start(handler)
	} else {
		log.Println("running locally http://localhost:8085")
		r.Run(":8085")
	}
}
