package main


import (
        
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/jmespath/go-jmespath"
	
)

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
}))

input := &dynamodb.UpdateItemInput{
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
        ":r": {
            N: aws.String(movies),
        },
    },
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
        "Year": {
            N: aws.String(Year),
        },
        "Movie": {
            S: aws.String(Movie),
        },
    },
    ReturnValues:     aws.String("UPDATED_NEW"),
    UpdateExpression: aws.String("set Rating = :r"),
}

_, err := svc.UpdateItem(input)
if err != nil {
    fmt.Println(err.Error())
    return
}

// Create DynamoDB client
svc := dynamodb.New(sess)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Table updated!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}


// // main.go
// package main

// import (
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// func Handler(request Request) (Response, error) {
// 	return "Hello!", nil
// }

// func main() {
// 	// Make the handler available for Remote Procedure Call by AWS Lambda
// 	lambda.Start(Handler)
// }