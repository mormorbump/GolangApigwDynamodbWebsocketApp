package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"log"
	"net/url"
)

var endpoint url.URL

type ClientHandler struct {
	Client *apigatewaymanagementapi.Client
}

func NewClientHandler(cfg aws.Config, scheme string, path string, host string) *ClientHandler {
	endpoint.Scheme = scheme
	endpoint.Path = path
	endpoint.Host = host

	endpointResolver := apigatewaymanagementapi.EndpointResolverFromURL(endpoint.String())
	client := apigatewaymanagementapi.NewFromConfig(
		cfg,
		apigatewaymanagementapi.WithEndpointResolver(endpointResolver),
	)
	return &ClientHandler{Client: client}
}

func (ch *ClientHandler) setPostToConnectionInput(connectionId string, message []byte) *apigatewaymanagementapi.PostToConnectionInput {
	return &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionId),
		Data:         message,
	}
}

func (ch *ClientHandler) GetConnection(connectionId string) (bool, error) {
	out, err := ch.Client.GetConnection(context.TODO(), &apigatewaymanagementapi.GetConnectionInput{ConnectionId: aws.String(connectionId)})
	if out == nil {
		log.Println("errないかな", err)
		return false, err
	}
	if err != nil {
		//log.Println("connectionがきれています", err)
		return false, err
	}
	return true, nil
}

func (ch *ClientHandler) PostToConnection(connectionId string, response interface{}) error {
	platformByte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = ch.Client.PostToConnection(context.TODO(), ch.setPostToConnectionInput(connectionId, platformByte))
	if err != nil {
		return err
	}
	return nil
}
