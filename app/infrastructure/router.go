//EntryPoint(Routing)
package infrastructure

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/controllers"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/protocols"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/infrastructure/handlers"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
	"net/http"
	"os"
)

var (
	webSocketController *controllers.WebSocketController
)

type Response = events.APIGatewayProxyResponse

func LambdaStart(handler interface{}) {
	lambda.Start(handler)
}

func init() {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dbHandler := handlers.NewDBHandler(
		cfg,
		os.Getenv("DB_PRIMARY_KEY"),
		os.Getenv("DB_SORT_KEY"),
		os.Getenv("CONNECTION_TABLE_NAME"),
		os.Getenv("DB_GSI_NAME"),
		os.Getenv("DB_TTL_KEY"),
		os.Getenv("DB_ATTR_KEY1"),
	)

	clientHandler := handlers.NewClientHandler(
		cfg,
		"https",
		os.Getenv("AWS_ENV"),
		os.Getenv("APIGW_HOST"),
	)

	webSocketController = controllers.NewWebSocketController(dbHandler, clientHandler)
}

func ConnectionHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (Response, error) {
	conn, err := webSocketController.Create(req.RequestContext.ConnectionID, req.QueryStringParameters)
	if err != nil {
		log.Println(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	connByte, err := json.Marshal(conn)
	return Response{StatusCode: http.StatusOK, Body: string(connByte), Headers: map[string]string{"connectionId": conn.ConnectionID}}, nil
}

func DisconnectionHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (Response, error) {
	if err := webSocketController.Destroy(req.RequestContext.ConnectionID); err != nil {
		log.Println(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	return Response{StatusCode: http.StatusOK}, nil
}

func HTTPRequestSendMessageHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (Response, error) {
	webhookMessage := protocols.WebhookMessage{}
	err := json.Unmarshal([]byte(req.Body), &webhookMessage)
	if err != nil {
		log.Println(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	if err := webSocketController.SendMessage(&webhookMessage); err != nil {
		log.Println(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	return Response{StatusCode: http.StatusOK, Body: "success! send message"}, nil
}
