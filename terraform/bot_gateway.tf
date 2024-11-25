# Resource for API Gateway /api/v1/bot
resource "aws_api_gateway_resource" "bot_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "bot"
}

# Resource for API Gateway /api/v1/bot/{id}
resource "aws_api_gateway_resource" "bot_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_resource.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/bot/restaurant
resource "aws_api_gateway_resource" "bot_restaurant_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_resource.id
  path_part   = "restaurant"
}

# Resource for API Gateway /api/v1/bot/restaurant/{id}
resource "aws_api_gateway_resource" "bot_restaurant_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_restaurant_resource.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/bot/whatsapp
resource "aws_api_gateway_resource" "bot_whatsapp_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_resource.id
  path_part   = "whatsapp"
}

# Resource for API Gateway /api/v1/bot/whatsapp/{whatsapp}
resource "aws_api_gateway_resource" "bot_whatsapp_whatsapp_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_whatsapp_resource.id
  path_part   = "{whatsapp}"
}

# Resource for API Gateway /api/v1/bot/twilio
resource "aws_api_gateway_resource" "bot_twilio_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.bot_resource.id
  path_part   = "twilio"
}

# Method POST for /api/v1/bot - Create bot
resource "aws_api_gateway_method" "bot_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/bot/{id} - Delete bot
resource "aws_api_gateway_method" "bot_method_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_id_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/bot - Get all bots
resource "aws_api_gateway_method" "bot_method_get_all" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_resource.id
  http_method   = "GET"
  authorization = "NONE"

}

# Method GET for /api/v1/bot/{id} - Get bot by id
resource "aws_api_gateway_method" "bot_method_get_by_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/restaurant{id} - Get bot by restaurant id
resource "aws_api_gateway_method" "bot_method_get_by_restaurant_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_restaurant_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/whatsapp{whatsapp} - Get bot by whatsapp
resource "aws_api_gateway_method" "bot_method_get_by_whatsapp" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_whatsapp_whatsapp_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method PUT for /api/v1/bot/{id} - Update bot
resource "aws_api_gateway_method" "bot_method_put" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_id_resource.id
  http_method   = "PUT"
  authorization = "NONE"
}

# Method POST for /api/v1/bot/twilio - Bot response using Twilio
resource "aws_api_gateway_method" "bot_method_post_twilio" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.bot_twilio_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Integration for POST /api/v1/bot - Create bot
resource "aws_api_gateway_integration" "bot_integration_post" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_resource.id
  http_method = aws_api_gateway_method.bot_method_post.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for DELETE /api/v1/bot/{id} - Delete bot
resource "aws_api_gateway_integration" "bot_integration_delete" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_id_resource.id
  http_method = aws_api_gateway_method.bot_method_delete.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot - Get all bots
resource "aws_api_gateway_integration" "bot_integration_get_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_resource.id
  http_method = aws_api_gateway_method.bot_method_get_all.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/{id} - Get bot by id
resource "aws_api_gateway_integration" "bot_integration_get_by_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_id_resource.id
  http_method = aws_api_gateway_method.bot_method_get_by_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/restaurant{id} - Get bot by restaurant id
resource "aws_api_gateway_integration" "bot_integration_get_by_restaurant_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_restaurant_id_resource.id
  http_method = aws_api_gateway_method.bot_method_get_by_restaurant_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/whatsapp{whatsapp} - Get bot by whatsapp
resource "aws_api_gateway_integration" "bot_integration_get_by_whatsapp" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_whatsapp_whatsapp_resource.id
  http_method = aws_api_gateway_method.bot_method_get_by_whatsapp.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for PUT /api/v1/bot/{id} - Update bot
resource "aws_api_gateway_integration" "bot_integration_put" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.bot_id_resource.id
  http_method = aws_api_gateway_method.bot_method_put.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Invoke permission for  API Gateway to call Lambda function
resource "aws_lambda_permission" "restaurant_menu_gateway_invoke_bot_service" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.bot_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api_gateway.execution_arn}/*/*/*"
}
