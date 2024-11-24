resource "aws_api_gateway_rest_api" "restaurant_menu_api" {
  name        = "restaurant_menu_api"
  description = "Restaurant Menu API"
}

# Resource for API Gateway /api endpoint
resource "aws_api_gateway_resource" "api" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_rest_api.restaurant_menu_api.root_resource_id
  path_part   = "api"
}

# Resource for API Gateway /api/v1 endpoint
resource "aws_api_gateway_resource" "v1" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.api.id
  path_part   = "v1"
}

############################################
############################################
##     USER SERVICE GATEWAY ENDPOINTS     ##
############################################
############################################

# Resource for API Gateway /api/v1/user endpoint
resource "aws_api_gateway_resource" "user" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "user"
}

# Resource for API Gateway /api/v1/user/{id} endpoint
resource "aws_api_gateway_resource" "user_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.user.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/user/all endpoint
resource "aws_api_gateway_resource" "user_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.user.id
  path_part   = "all"
}

# Resource for API Gateway /api/v1/user/email endpoint
resource "aws_api_gateway_resource" "user_email" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_rest_api.restaurant_menu_api.root_resource_id
  path_part   = "email"
}


# Resource for API Gateway /api/v1/user/email/{email} endpoint
resource "aws_api_gateway_resource" "user_email_email" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.user_email.id
  path_part   = "{email}"
}

# Resource for API Gateway /api/v1/user/update endpoint
resource "aws_api_gateway_resource" "user_update" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.user.id
  path_part   = "update"
}

# Resource for API Gateway /api/v1/user/update/{id} endpoint
resource "aws_api_gateway_resource" "user_update_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.user_update.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/user/login endpoint
resource "aws_api_gateway_resource" "user_login" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_rest_api.restaurant_menu_api.root_resource_id
  path_part   = "login"
}

############################################
# HTTP METHODS FOR USER SERVICE ENDPOINTS  #
############################################

# Method POST for /api/v1/user endpoint
resource "aws_api_gateway_method" "user_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/user/{id} endpoint
resource "aws_api_gateway_method" "user_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_id.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/user/all endpoint
resource "aws_api_gateway_method" "user_all_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_all.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/user/{id} endpoint
resource "aws_api_gateway_method" "user_id_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_id.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/user/email/{email} endpoint
resource "aws_api_gateway_method" "user_email_email_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_email_email.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method POST for /api/v1/user/update/{id} endpoint
resource "aws_api_gateway_method" "user_update_id_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_update_id.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method POST for /api/v1/user/login endpoint
resource "aws_api_gateway_method" "user_login_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.user_login.id
  http_method   = "POST"
  authorization = "NONE"
}

############################################
# Integration for USER SERVICE ENDPOINTS   #
############################################

# Integration for POST /api/v1/user endpoint
resource "aws_api_gateway_integration" "user_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user.id
  http_method             = aws_api_gateway_method.user_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for DELETE /api/v1/user/{id} endpoint
resource "aws_api_gateway_integration" "user_delete" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_id.id
  http_method             = aws_api_gateway_method.user_delete.http_method
  integration_http_method = "DELETE"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user/all endpoint
resource "aws_api_gateway_integration" "user_all_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_all.id
  http_method             = aws_api_gateway_method.user_all_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user/{id} endpoint
resource "aws_api_gateway_integration" "user_id_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_id.id
  http_method             = aws_api_gateway_method.user_id_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for GET /api/v1/user/email/{email} endpoint
resource "aws_api_gateway_integration" "user_email_email_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_email_email.id
  http_method             = aws_api_gateway_method.user_email_email_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for POST /api/v1/user/update/{id} endpoint
resource "aws_api_gateway_integration" "user_update_id_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_update_id.id
  http_method             = aws_api_gateway_method.user_update_id_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

# Integration for POST /api/v1/user/login endpoint
resource "aws_api_gateway_integration" "user_login_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.user_login.id
  http_method             = aws_api_gateway_method.user_login_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_service.invoke_arn
}

#############################################
#############################################
##      RESTAURANT MENU SERVICE ENDPOINTS  ##
#############################################
#############################################

# Resource for API Gateway /api/v1/restaurant endpoint
resource "aws_api_gateway_resource" "restaurant" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "restaurant"
}

# Resource for API Gateway /api/v1/menu endpoint
resource "aws_api_gateway_resource" "menu" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "menu"
}


# Resource for API Gateway /api/v1/restaurant/{id} endpoint
resource "aws_api_gateway_resource" "restaurant_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.restaurant.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/restaurant/all endpoint
resource "aws_api_gateway_resource" "restaurant_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.restaurant.id
  path_part   = "all"
}

# Resource for API Gateway /api/v1/restaurant/update endpoint
resource "aws_api_gateway_resource" "restaurant_update" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.restaurant.id
  path_part   = "update"
}

# Resource for API Gateway /api/v1/restaurant/update/{id} endpoint
resource "aws_api_gateway_resource" "restaurant_update_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.restaurant_update.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/menu/{id} endpoint
resource "aws_api_gateway_resource" "menu_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.menu.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/menu/all endpoint
resource "aws_api_gateway_resource" "menu_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.menu.id
  path_part   = "all"
}

# Resource for API Gateway /api/v1/menu/search endpoint
resource "aws_api_gateway_resource" "menu_search" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.menu.id
  path_part   = "search"
}

# Resource for API Gateway /api/v1/menu/update endpoint
resource "aws_api_gateway_resource" "menu_update" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.menu.id
  path_part   = "update"
}

# Resource for API Gateway /api/v1/menu/update/{id} endpoint
resource "aws_api_gateway_resource" "menu_update_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.menu_update.id
  path_part   = "{id}"
}

############################################
# HTTP METHODS FOR RESTAURANT MENU SERVICE #
############################################

# Method POST for /api/v1/restaurant endpoint
resource "aws_api_gateway_method" "restaurant_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.restaurant.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/restaurant/{id} endpoint
resource "aws_api_gateway_method" "restaurant_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.restaurant_id.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/restaurant/all endpoint
resource "aws_api_gateway_method" "restaurant_all_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.restaurant_all.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/restaurant/{id} endpoint
resource "aws_api_gateway_method" "restaurant_id_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.restaurant_id.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method POST for /api/v1/restaurant/update/{id} endpoint
resource "aws_api_gateway_method" "restaurant_update_id_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.restaurant_update_id.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method POST for /api/v1/menu endpoint
resource "aws_api_gateway_method" "menu_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/menu/{id} endpoint
resource "aws_api_gateway_method" "menu_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu_id.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/menu/all endpoint
resource "aws_api_gateway_method" "menu_all_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu_all.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/menu/search endpoint
resource "aws_api_gateway_method" "menu_search_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu_search.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/menu/{id} endpoint
resource "aws_api_gateway_method" "menu_id_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu_id.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method POST for /api/v1/menu/update/{id} endpoint
resource "aws_api_gateway_method" "menu_update_id_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.menu_update_id.id
  http_method   = "POST"
  authorization = "NONE"
}

############################################
# Integration for RESTAURANT MENU SERVICE  #
############################################

# Integration for POST /api/v1/restaurant endpoint
resource "aws_api_gateway_integration" "restaurant_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.restaurant.id
  http_method             = aws_api_gateway_method.restaurant_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for DELETE /api/v1/restaurant/{id} endpoint
resource "aws_api_gateway_integration" "restaurant_delete" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.restaurant_id.id
  http_method             = aws_api_gateway_method.restaurant_delete.http_method
  integration_http_method = "DELETE"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for GET /api/v1/restaurant/all endpoint
resource "aws_api_gateway_integration" "restaurant_all_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.restaurant_all.id
  http_method             = aws_api_gateway_method.restaurant_all_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for GET /api/v1/restaurant/{id} endpoint
resource "aws_api_gateway_integration" "restaurant_id_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.restaurant_id.id
  http_method             = aws_api_gateway_method.restaurant_id_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for POST /api/v1/restaurant/update/{id} endpoint
resource "aws_api_gateway_integration" "restaurant_update_id_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.restaurant_update_id.id
  http_method             = aws_api_gateway_method.restaurant_update_id_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for POST /api/v1/menu endpoint
resource "aws_api_gateway_integration" "menu_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu.id
  http_method             = aws_api_gateway_method.menu_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for DELETE /api/v1/menu/{id} endpoint
resource "aws_api_gateway_integration" "menu_delete" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu_id.id
  http_method             = aws_api_gateway_method.menu_delete.http_method
  integration_http_method = "DELETE"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for GET /api/v1/menu/all endpoint
resource "aws_api_gateway_integration" "menu_all_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu_all.id
  http_method             = aws_api_gateway_method.menu_all_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for GET /api/v1/menu/search endpoint
resource "aws_api_gateway_integration" "menu_search_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu_search.id
  http_method             = aws_api_gateway_method.menu_search_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for GET /api/v1/menu/{id} endpoint
resource "aws_api_gateway_integration" "menu_id_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu_id.id
  http_method             = aws_api_gateway_method.menu_id_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for POST /api/v1/menu/update/{id} endpoint
resource "aws_api_gateway_integration" "menu_update_id_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.menu_update_id.id
  http_method             = aws_api_gateway_method.menu_update_id_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

############################################
############################################
##         BOT SERVICE ENDPOINTS          ##
############################################
############################################

# Resource for API Gateway /api/v1/bot endpoint
resource "aws_api_gateway_resource" "bot" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "bot"
}

# Resource for API Gateway /api/v1/bot-response endpoint
resource "aws_api_gateway_resource" "bot_response" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "bot-response"
}

# Resource for API Gateway /api/v1/bot/{id} endpoint
resource "aws_api_gateway_resource" "bot_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/bot/all endpoint
resource "aws_api_gateway_resource" "bot_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot.id
  path_part   = "all"
}

# Resource for API Gateway /api/v1/bot/restaurant endpoint
resource "aws_api_gateway_resource" "bot_restaurant" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot.id
  path_part   = "restaurant"
}

# Resource for API Gateway /api/v1/bot/restaurant/{id} endpoint
resource "aws_api_gateway_resource" "bot_restaurant_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot_restaurant.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/bot/whatsapp endpoint
resource "aws_api_gateway_resource" "bot_whatsapp" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot.id
  path_part   = "whatsapp"
}

# Resource for API Gateway /api/v1/bot/whatsapp/{whatsapp} endpoint
resource "aws_api_gateway_resource" "bot_whatsapp_whatsapp" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot_whatsapp.id
  path_part   = "{whatsapp}"
}

# Resource for API Gateway /api/v1/bot/update endpoint
resource "aws_api_gateway_resource" "bot_update" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot.id
  path_part   = "update"
}

# Resource for API Gateway /api/v1/bot/update/{id} endpoint
resource "aws_api_gateway_resource" "bot_update_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot_update.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/bot-response/twilio endpoint
resource "aws_api_gateway_resource" "bot_response_twilio" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot_response.id
  path_part   = "twilio"
}

# Resource for API Gateway /api/v1/bot-response/twilio/webhook endpoint
resource "aws_api_gateway_resource" "bot_response_twilio_webhook" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  parent_id   = aws_api_gateway_resource.bot_response_twilio.id
  path_part   = "webhook"
}

############################################
# HTTP METHODS FOR BOT SERVICE ENDPOINTS   #
############################################

# Method POST for /api/v1/bot endpoint
resource "aws_api_gateway_method" "bot_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/bot/{id} endpoint
resource "aws_api_gateway_method" "bot_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_id.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/all endpoint
resource "aws_api_gateway_method" "bot_all_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_all.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/{id} endpoint
resource "aws_api_gateway_method" "bot_id_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_id.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/restaurant/{id} endpoint
resource "aws_api_gateway_method" "bot_restaurant_id_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_restaurant_id.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/bot/whatsapp/{whatsapp} endpoint
resource "aws_api_gateway_method" "bot_whatsapp_whatsapp_get" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_whatsapp_whatsapp.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method POST for /api/v1/bot/update/{id} endpoint
resource "aws_api_gateway_method" "bot_update_id_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_update_id.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method POST for /api/v1/bot-response/twilio/webhook endpoint
resource "aws_api_gateway_method" "bot_response_twilio_webhook_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_response_twilio_webhook.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method OPTIONS for /api/v1/bot-response/twilio/webhook endpoint
resource "aws_api_gateway_method" "bot_response_twilio_webhook_options" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id   = aws_api_gateway_resource.bot_response_twilio_webhook.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

############################################
# Integration for BOT SERVICE ENDPOINTS    #
############################################

# Integration for POST /api/v1/bot endpoint
resource "aws_api_gateway_integration" "bot_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot.id
  http_method             = aws_api_gateway_method.bot_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for DELETE /api/v1/bot/{id} endpoint
resource "aws_api_gateway_integration" "bot_delete" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_id.id
  http_method             = aws_api_gateway_method.bot_delete.http_method
  integration_http_method = "DELETE"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/all endpoint
resource "aws_api_gateway_integration" "bot_all_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_all.id
  http_method             = aws_api_gateway_method.bot_all_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/{id} endpoint
resource "aws_api_gateway_integration" "bot_id_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_id.id
  http_method             = aws_api_gateway_method.bot_id_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/restaurant/{id} endpoint
resource "aws_api_gateway_integration" "bot_restaurant_id_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_restaurant_id.id
  http_method             = aws_api_gateway_method.bot_restaurant_id_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for GET /api/v1/bot/whatsapp/{whatsapp} endpoint
resource "aws_api_gateway_integration" "bot_whatsapp_whatsapp_get" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_whatsapp_whatsapp.id
  http_method             = aws_api_gateway_method.bot_whatsapp_whatsapp_get.http_method
  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for POST /api/v1/bot/update/{id} endpoint
resource "aws_api_gateway_integration" "bot_update_id_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_update_id.id
  http_method             = aws_api_gateway_method.bot_update_id_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for POST /api/v1/bot-response/twilio/webhook endpoint
resource "aws_api_gateway_integration" "bot_response_twilio_webhook_post" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_response_twilio_webhook.id
  http_method             = aws_api_gateway_method.bot_response_twilio_webhook_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.bot_service.invoke_arn
}

# Integration for OPTIONS /api/v1/bot-response/twilio/webhook endpoint
resource "aws_api_gateway_integration" "bot_response_twilio_webhook_options" {
  rest_api_id             = aws_api_gateway_rest_api.restaurant_menu_api.id
  resource_id             = aws_api_gateway_resource.bot_response_twilio_webhook.id
  http_method             = aws_api_gateway_method.bot_response_twilio_webhook_options.http_method
  integration_http_method = "OPTIONS"
  type                    = "MOCK"
  request_templates = {
    "application/json" = jsonencode({
      statusCode = 200
    })
  }
}


#############################
# LAMBDA INVOKE PERMISSIONS #
#############################

# Invoke permission for user_service
resource "aws_lambda_permission" "user_service" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.user_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api.execution_arn}/*/*/*"
}

# Invoke permission for restaurant_menu_service
resource "aws_lambda_permission" "restaurant_menu_service" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.restaurant_menu_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api.execution_arn}/*/*/*"
}

# Invoke permission for bot_service
resource "aws_lambda_permission" "bot_service" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.bot_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api.execution_arn}/*/*/*"
}

##################################################
# DEPLOYMENT FOR API GATEWAY RESTAURANT MENU API #
##################################################
resource "aws_api_gateway_deployment" "restaurant_menu_api_deployment" {
  depends_on = [
    aws_api_gateway_integration.user_post,
    aws_api_gateway_integration.user_delete,
    aws_api_gateway_integration.user_all_get,
    aws_api_gateway_integration.user_id_get,
    aws_api_gateway_integration.user_email_email_get,
    aws_api_gateway_integration.user_update_id_post,
    aws_api_gateway_integration.user_login_post,
    aws_api_gateway_integration.restaurant_post,
    aws_api_gateway_integration.restaurant_delete,
    aws_api_gateway_integration.restaurant_all_get,
    aws_api_gateway_integration.restaurant_id_get,
    aws_api_gateway_integration.restaurant_update_id_post,
    aws_api_gateway_integration.menu_post,
    aws_api_gateway_integration.menu_delete,
    aws_api_gateway_integration.menu_all_get,
    aws_api_gateway_integration.menu_search_get,
    aws_api_gateway_integration.menu_id_get,
    aws_api_gateway_integration.menu_update_id_post,
    aws_api_gateway_integration.bot_post,
    aws_api_gateway_integration.bot_delete,
    aws_api_gateway_integration.bot_all_get,
    aws_api_gateway_integration.bot_id_get,
    aws_api_gateway_integration.bot_restaurant_id_get,
    aws_api_gateway_integration.bot_whatsapp_whatsapp_get,
    aws_api_gateway_integration.bot_update_id_post,
    aws_api_gateway_integration.bot_response_twilio_webhook_post,
    aws_api_gateway_integration.bot_response_twilio_webhook_options,
  ]
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api.id
  description = "Deployment for the restaurant menu API"

  triggers = {
    redeployment = sha1(join(",", [
      jsonencode(aws_api_gateway_integration.user_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_delete.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_all_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_id_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_email_email_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_update_id_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.user_login_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.restaurant_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.restaurant_delete.integration_http_method),
      jsonencode(aws_api_gateway_integration.restaurant_all_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.restaurant_id_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.restaurant_update_id_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_delete.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_all_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_search_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_id_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.menu_update_id_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_delete.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_all_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_id_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_restaurant_id_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_whatsapp_whatsapp_get.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_update_id_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_response_twilio_webhook_post.integration_http_method),
      jsonencode(aws_api_gateway_integration.bot_response_twilio_webhook_options.integration_http_method)
    ]))
  }

  lifecycle {
    create_before_destroy = true
  }
}

#####################################################
# GATEWAY STAGE FOR API GATEWAY RESTAURANT MENU API #
#####################################################

resource "aws_api_gateway_stage" "restaurant_menu_api_stage" {
  stage_name    = "dev"
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api.id
  deployment_id = aws_api_gateway_deployment.restaurant_menu_api_deployment.id
  description   = "dev stage for the restaurant menu API"
  depends_on    = [aws_api_gateway_deployment.restaurant_menu_api_deployment]
}

######################################################
# OUTPUTS FOR API GATEWAY RESTAURANT MENU API        #
######################################################
output "restaurant_menu_api_gateway_invoke_url" {
  value = aws_api_gateway_stage.restaurant_menu_api_stage.invoke_url
}
