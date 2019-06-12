package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	storage := dynamo.NewStorage(db, tableName)

	svc := business.NewService(storage)
	svc = app.NewAuthMiddleware(svc)
	svc = app.NewLoggingMiddleware(svc, logger)

	handler := apigw.MakeHandler(svc)

	lambda.Start(handler)
}
