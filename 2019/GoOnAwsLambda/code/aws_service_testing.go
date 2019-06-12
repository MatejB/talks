package code

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func TestX(t *testing.T) {
	//....

	dynamoDBServer := httptest.NewServer(func(w http.ResponseWriter, req *http.Request) {
		// mock response and assert request
	})
	defer dynamoDBServer.Close()

	cfg, _ := external.LoadDefaultAWSConfig()
	resolver := func(string, string) (aws.Endpoint, error) {
		return aws.Endpoint{URL: dynamoDBServer.URL}, nil
	}
	cfg.EndpointResolver = aws.EndpointResolverFunc(resolver)
	db := dynamodb.New(cfg)

	//....
}
