resource "aws_iam_role" "iam_for_lambda" {
  name = replace("${var.system_name_prefix}_iam_for_lambda", "_", "-")

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action : "sts:AssumeRole",
        Principal : {
          Service : "lambda.amazonaws.com"
        },
        Effect : "Allow",
        Sid : ""
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "dynamodb_lambda_policy" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy" "dynamodb_policy" {
  name = replace("${var.system_name_prefix}_dynamodb_policy", "_", "-")
  role = aws_iam_role.iam_for_lambda.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:*",
        ]
        Effect = "Allow"
        // dynamodbのindexのiam policyまでのpathもかく
        // https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/iam-policy-specific-table-indexes.html
        Resource = [
          "arn:aws:dynamodb:*:*:table/${aws_dynamodb_table.connection.name}",
          "arn:aws:dynamodb:*:*:table/${aws_dynamodb_table.connection.name}/index/*"
        ]
      },
    ]
  })
}


resource "aws_iam_role_policy" "mng_api_policy" {
  name = replace("${var.system_name_prefix}_mng_api_policy", "_", "-")
  role = aws_iam_role.iam_for_lambda.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "execute-api:ManageConnections",
        ]
        Effect   = "Allow"
        Resource = [
          "${aws_apigatewayv2_api.websocket_gw.execution_arn}/${aws_apigatewayv2_stage.websocket_stage.name}/POST/@connections/{connectionId}",
          "${aws_apigatewayv2_api.websocket_gw.execution_arn}/${aws_apigatewayv2_stage.websocket_stage.name}/GET/@connections/{connectionId}"
        ]
      },
    ]
  })
}