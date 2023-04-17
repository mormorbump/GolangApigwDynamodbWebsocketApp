resource "aws_api_gateway_rest_api" "rest_api" {
  name        = replace("${var.system_name_prefix}_restapi_gw", "_", "-")
  description = "API server to WebSocket sendMessage"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_stage" "env" {
  stage_name    = var.aws_env
  rest_api_id   = aws_api_gateway_rest_api.rest_api.id
  deployment_id = aws_api_gateway_deployment.default.id
}



resource "aws_api_gateway_deployment" "default" {
  rest_api_id = aws_api_gateway_rest_api.rest_api.id

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_rest_api.rest_api.body,
      aws_api_gateway_resource.send_message,
      aws_api_gateway_method.send_message_post,
      aws_api_gateway_integration.send_message_post,
      aws_api_gateway_method_response.send_message_response_200,

      aws_api_gateway_resource.send_gift,
      aws_api_gateway_method.send_gift_post,
      aws_api_gateway_integration.send_gift_post,
      aws_api_gateway_method_response.send_gift_response_200,
    ]))
  }

  lifecycle {
    create_before_destroy = true
  }
}

################################################################################
# Path  : /send_message                                                       #
# Method: POST                                                                 #
################################################################################
resource "aws_api_gateway_resource" "send_message" {
  path_part   = "send_message"
  parent_id   = aws_api_gateway_rest_api.rest_api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
}

resource "aws_api_gateway_method" "send_message_post" {
  rest_api_id   = aws_api_gateway_rest_api.rest_api.id
  resource_id   = aws_api_gateway_resource.send_message.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "send_message_post" {
  rest_api_id             = aws_api_gateway_rest_api.rest_api.id
  resource_id             = aws_api_gateway_resource.send_message.id
  http_method             = aws_api_gateway_method.send_message_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.send_message.invoke_arn
}

resource "aws_api_gateway_method_response" "send_message_response_200" {
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
  resource_id = aws_api_gateway_resource.send_message.id
  http_method = aws_api_gateway_method.send_message_post.http_method
  status_code = "200"
}


################################################################################
# Path  : /send_gift                                                           #
# Method: POST                                                                 #
################################################################################
resource "aws_api_gateway_resource" "send_gift" {
  path_part   = "send_gift"
  parent_id   = aws_api_gateway_rest_api.rest_api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
}

resource "aws_api_gateway_method" "send_gift_post" {
  rest_api_id   = aws_api_gateway_rest_api.rest_api.id
  resource_id   = aws_api_gateway_resource.send_gift.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "send_gift_post" {
  rest_api_id             = aws_api_gateway_rest_api.rest_api.id
  resource_id             = aws_api_gateway_resource.send_gift.id
  http_method             = aws_api_gateway_method.send_gift_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.send_gift.invoke_arn
}

resource "aws_api_gateway_method_response" "send_gift_response_200" {
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
  resource_id = aws_api_gateway_resource.send_gift.id
  http_method = aws_api_gateway_method.send_gift_post.http_method
  status_code = "200"
}