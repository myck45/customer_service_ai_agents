# Resource for API Gateway /api/v1/user
resource "aws_api_gateway_resource" "user_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "user"
}

# Resource for API Gateway /api/v1/user/{id}
resource "aws_api_gateway_resource" "user_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.user_resource.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/user/email
resource "aws_api_gateway_resource" "user_email_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.user_resource.id
  path_part   = "email"
}


# Resource for API Gateway /api/v1/user/email/{email}
resource "aws_api_gateway_resource" "user_email_email_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.user_email_resource.id
  path_part   = "{email}"
}

# Resource for API Gateway /api/v1/user/login
resource "aws_api_gateway_resource" "user_login_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.user_resource.id
  path_part   = "login"
}

# Method POST for /api/v1/user - Create user
resource "aws_api_gateway_method" "user_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/user/{id} - Delete user
resource "aws_api_gateway_method" "user_method_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_id_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/user - Get all users
resource "aws_api_gateway_method" "user_method_get_all" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/user/{id} - Get user by id
resource "aws_api_gateway_method" "user_method_get_by_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/user/email/{email} - Get user by email
resource "aws_api_gateway_method" "user_method_get_by_email" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_email_email_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method PUT for /api/v1/user/{id} - Update user
resource "aws_api_gateway_method" "user_method_put" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_id_resource.id
  http_method   = "PUT"
  authorization = "NONE"
}

# Method POST for /api/v1/user/login - Login user
resource "aws_api_gateway_method" "user_method_post_login" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.user_login_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Integration for POST /api/v1/user - Create user
resource "aws_api_gateway_integration" "user_integration_post" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_resource.id
  http_method = aws_api_gateway_method.user_method_post.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for DELETE /api/v1/user/{id} - Delete user
resource "aws_api_gateway_integration" "user_integration_delete" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_id_resource.id
  http_method = aws_api_gateway_method.user_method_delete.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user - Get all users
resource "aws_api_gateway_integration" "user_integration_get_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_resource.id
  http_method = aws_api_gateway_method.user_method_get_all.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user/{id} - Get user by id
resource "aws_api_gateway_integration" "user_integration_get_by_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_id_resource.id
  http_method = aws_api_gateway_method.user_method_get_by_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user/email/{email} - Get user by email
resource "aws_api_gateway_integration" "user_integration_get_by_email" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_email_email_resource.id
  http_method = aws_api_gateway_method.user_method_get_by_email.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for PUT /api/v1/user/{id} - Update user
resource "aws_api_gateway_integration" "user_integration_put" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_id_resource.id
  http_method = aws_api_gateway_method.user_method_put.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for POST /api/v1/user/login - Login user
resource "aws_api_gateway_integration" "user_integration_post_login" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.user_login_resource.id
  http_method = aws_api_gateway_method.user_method_post_login.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Invoke permission for API Gateway to call Lambda function
resource "aws_lambda_permission" "restaurant_menu_gateway_invoke_lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.user_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api_gateway.execution_arn}/*/*/*"
}
