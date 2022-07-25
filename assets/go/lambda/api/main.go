package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/Smerity/govarint"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/guregu/dynamo"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/BrianLeishman/breker.ninja/assets/go/validators"
	_ "github.com/joho/godotenv/autoload"
)

var ginLambda *ginadapter.GinLambda

var r = gin.Default()

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

var table dynamo.Table

func main() {
	db := dynamo.New(session.New())
	table = db.Table("breker")

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

func init() {
	b, _ := hex.DecodeString("0800000543a80092020f58afc80005b8d8200002f9b8")
	dec := govarint.NewU32GroupVarintDecoder(bytes.NewReader(b))

	for i := 0; i < 10; i++ {
		fmt.Println(dec.GetU32())
	}
}
