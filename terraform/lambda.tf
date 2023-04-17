data "aws_region" "current" {}

// Archive Function
// Goのビルドファイルをデプロイ用にまとめてzip
data "archive_file" "function_zip" {
  type        = "zip"
  source_dir  = "../../app/bin"
  output_path = "function.zip"
}

resource "aws_lambda_function" "connection" {
  function_name = replace("${var.system_name_prefix}_connection", "_", "-")
  filename      = data.archive_file.function_zip.output_path
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "connection"

  publish          = true
  source_code_hash = data.archive_file.function_zip.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      CONNECTION_TABLE_NAME = aws_dynamodb_table.connection.name
      AWS_ENV               = var.aws_env
      DB_PRIMARY_KEY        = var.dynamodb_hash_key
      DB_SORT_KEY           = var.dynamodb_range_key
      API_SERVER_BASE_URL   = var.api_server_base_url
      APIGW_HOST            = "${aws_apigatewayv2_api.websocket_gw.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
      DB_GSI_NAME           = var.dynamodb_gsi_name
      DB_TTL_KEY            = var.dynamodb_ttl_key
      DB_ATTR_KEY1          = var.dynamodb_attr_key1
    }
  }
}

resource "aws_lambda_permission" "connection" {
  statement_id  = "AlowApiGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.connection.function_name
  principal     = "apigateway.amazonaws.com"
}

resource "aws_lambda_function" "disconnection" {
  function_name = replace("${var.system_name_prefix}_disconnection", "_", "-")
  filename      = data.archive_file.function_zip.output_path
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "disconnection"

  publish          = true
  source_code_hash = data.archive_file.function_zip.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      CONNECTION_TABLE_NAME = aws_dynamodb_table.connection.name
      AWS_ENV               = var.aws_env
      DB_PRIMARY_KEY        = var.dynamodb_hash_key
      DB_SORT_KEY           = var.dynamodb_range_key
      API_SERVER_BASE_URL   = var.api_server_base_url
      APIGW_HOST            = "${aws_apigatewayv2_api.websocket_gw.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
      DB_GSI_NAME           = var.dynamodb_gsi_name
      DB_TTL_KEY            = var.dynamodb_ttl_key
      DB_ATTR_KEY1          = var.dynamodb_attr_key1
    }
  }
}

resource "aws_lambda_permission" "disconnection" {
  statement_id  = "AlowApiGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.disconnection.function_name
  principal     = "apigateway.amazonaws.com"
}

resource "aws_lambda_function" "send_message" {
  function_name = replace("${var.system_name_prefix}_send_message", "_", "-")
  filename      = data.archive_file.function_zip.output_path
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "send_message"

  publish          = true
  source_code_hash = data.archive_file.function_zip.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      CONNECTION_TABLE_NAME = aws_dynamodb_table.connection.name
      AWS_ENV               = var.aws_env
      DB_PRIMARY_KEY        = var.dynamodb_hash_key
      DB_SORT_KEY           = var.dynamodb_range_key
      API_SERVER_BASE_URL   = var.api_server_base_url
      APIGW_HOST            = "${aws_apigatewayv2_api.websocket_gw.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
      DB_GSI_NAME           = var.dynamodb_gsi_name
      DB_TTL_KEY            = var.dynamodb_ttl_key
      DB_ATTR_KEY1          = var.dynamodb_attr_key1
    }
  }
}

resource "aws_lambda_permission" "send_message" {
  statement_id  = "AlowApiGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.send_message.function_name
  principal     = "apigateway.amazonaws.com"
}





###

resource "aws_lambda_function" "send_gift" {
  function_name = replace("${var.system_name_prefix}_send_gift", "_", "-")
  filename      = data.archive_file.function_zip.output_path
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "send_gift"

  publish          = true
  source_code_hash = data.archive_file.function_zip.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      CONNECTION_TABLE_NAME = aws_dynamodb_table.connection.name
      AWS_ENV               = var.aws_env
      DB_PRIMARY_KEY        = var.dynamodb_hash_key
      DB_SORT_KEY           = var.dynamodb_range_key
      API_SERVER_BASE_URL   = var.api_server_base_url
      APIGW_HOST            = "${aws_apigatewayv2_api.websocket_gw.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
      DB_GSI_NAME           = var.dynamodb_gsi_name
      DB_TTL_KEY            = var.dynamodb_ttl_key
      DB_ATTR_KEY1          = var.dynamodb_attr_key1
    }
  }
}

resource "aws_lambda_permission" "send_gift" {
  statement_id  = "AlowApiGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.send_gift.function_name
  principal     = "apigateway.amazonaws.com"
}