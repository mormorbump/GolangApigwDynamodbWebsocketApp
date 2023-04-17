package main

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/infrastructure"
)

func main() {
	infrastructure.LambdaStart(infrastructure.DisconnectionHandler)
}
