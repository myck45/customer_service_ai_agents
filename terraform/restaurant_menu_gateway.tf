# Resource for API Gateway /api/v1/restaurant
resource "aws_api_gateway_resource" "restaurant_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "restaurant"
}

# Resource for API Gateway /api/v1/restaurant/{id}
resource "aws_api_gateway_resource" "restaurant_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.restaurant_resource.id
  path_part   = "{id}"
}

# Method POST for /api/v1/restaurant - Create restaurant
resource "aws_api_gateway_method" "restaurant_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.restaurant_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/restaurant/{id} - Delete restaurant
resource "aws_api_gateway_method" "restaurant_method_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.restaurant_id_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/restaurant - Get all restaurants
resource "aws_api_gateway_method" "restaurant_method_get_all" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.restaurant_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/restaurant/{id} - Get restaurant by id
resource "aws_api_gateway_method" "restaurant_method_get_by_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.restaurant_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method PUT for /api/v1/restaurant/{id} - Update restaurant
resource "aws_api_gateway_method" "restaurant_method_put" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.restaurant_id_resource.id
  http_method   = "PUT"
  authorization = "NONE"
}

# Integration for /api/v1/restaurant - Create restaurant
resource "aws_api_gateway_integration" "restaurant_integration_post" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.restaurant_resource.id
  http_method = aws_api_gateway_method.restaurant_method_post.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/restaurant/{id} - Delete restaurant
resource "aws_api_gateway_integration" "restaurant_integration_delete" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.restaurant_id_resource.id
  http_method = aws_api_gateway_method.restaurant_method_delete.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/restaurant - Get all restaurants
resource "aws_api_gateway_integration" "restaurant_integration_get_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.restaurant_resource.id
  http_method = aws_api_gateway_method.restaurant_method_get_all.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/restaurant/{id} - Get restaurant by id
resource "aws_api_gateway_integration" "restaurant_integration_get_by_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.restaurant_id_resource.id
  http_method = aws_api_gateway_method.restaurant_method_get_by_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/restaurant/{id} - Update restaurant
resource "aws_api_gateway_integration" "restaurant_integration_put" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.restaurant_id_resource.id
  http_method = aws_api_gateway_method.restaurant_method_put.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Resource for API Gateway /api/v1/menu
resource "aws_api_gateway_resource" "menu_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "menu"
}

# Resource for API Gateway /api/v1/menu/{id}
resource "aws_api_gateway_resource" "menu_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.menu_resource.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/menu/search
resource "aws_api_gateway_resource" "menu_search_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.menu_resource.id
  path_part   = "search"
}

# Method POST for /api/v1/menu - Create menu
resource "aws_api_gateway_method" "menu_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/menu/{id} - Delete menu
resource "aws_api_gateway_method" "menu_method_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_id_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/menu - Get all menus
resource "aws_api_gateway_method" "menu_method_get_all" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/menu/search - Search menu
resource "aws_api_gateway_method" "menu_method_search" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_search_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/menu/{id} - Get menu by id
resource "aws_api_gateway_method" "menu_method_get_by_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method PUT for /api/v1/menu/{id} - Update menu
resource "aws_api_gateway_method" "menu_method_put" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_id_resource.id
  http_method   = "PUT"
  authorization = "NONE"
}

# Integration for /api/v1/menu - Create menu
resource "aws_api_gateway_integration" "menu_integration_post" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_resource.id
  http_method = aws_api_gateway_method.menu_method_post.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu/{id} - Delete menu
resource "aws_api_gateway_integration" "menu_integration_delete" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_id_resource.id
  http_method = aws_api_gateway_method.menu_method_delete.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu - Get all menus
resource "aws_api_gateway_integration" "menu_integration_get_all" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_resource.id
  http_method = aws_api_gateway_method.menu_method_get_all.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu/search - Search menu
resource "aws_api_gateway_integration" "menu_integration_search" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_search_resource.id
  http_method = aws_api_gateway_method.menu_method_search.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu/{id} - Get menu by id
resource "aws_api_gateway_integration" "menu_integration_get_by_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_id_resource.id
  http_method = aws_api_gateway_method.menu_method_get_by_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu/{id} - Update menu
resource "aws_api_gateway_integration" "menu_integration_put" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_id_resource.id
  http_method = aws_api_gateway_method.menu_method_put.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Resource for API Gateway /api/v1/menu-files
resource "aws_api_gateway_resource" "menu_files_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.v1.id
  path_part   = "menu-files"
}

# Resource for API Gateway /api/v1/menu-files/{id}
resource "aws_api_gateway_resource" "menu_files_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.menu_files_resource.id
  path_part   = "{id}"
}

# Resource for API Gateway /api/v1/menu-files/restaurant
resource "aws_api_gateway_resource" "menu_files_restaurant_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.menu_files_resource.id
  path_part   = "restaurant"
}

# Resource for API Gateway /api/v1/menu-files/restaurant/{restaurant_id}
resource "aws_api_gateway_resource" "menu_files_restaurant_id_resource" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.menu_files_restaurant_resource.id
  path_part   = "{restaurant_id}"
}

# Method POST for /api/v1/menu-files - Upload menu file
resource "aws_api_gateway_method" "menu_files_method_post" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_files_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# Method DELETE for /api/v1/menu-files/{id} - Delete menu file
resource "aws_api_gateway_method" "menu_files_method_delete" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_files_id_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
}

# Method GET for /api/v1/menu-files/{id} - Get menu file by id
resource "aws_api_gateway_method" "menu_files_method_get_by_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_files_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method GET for /api/v1/menu-files/restaurant/{restaurant_id} - Get menu files by restaurant id
resource "aws_api_gateway_method" "menu_files_method_get_by_restaurant_id" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_files_restaurant_id_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

# Method PUT for /api/v1/menu-files/{id} - Update menu file
resource "aws_api_gateway_method" "menu_files_method_put" {
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id   = aws_api_gateway_resource.menu_files_id_resource.id
  http_method   = "PUT"
  authorization = "NONE"
}

# Integration for /api/v1/menu-files - Upload menu file
resource "aws_api_gateway_integration" "menu_files_integration_post" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_files_resource.id
  http_method = aws_api_gateway_method.menu_files_method_post.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu-files/{id} - Delete menu file
resource "aws_api_gateway_integration" "menu_files_integration_delete" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_files_id_resource.id
  http_method = aws_api_gateway_method.menu_files_method_delete.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu-files/{id} - Get menu file by id
resource "aws_api_gateway_integration" "menu_files_integration_get_by_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_files_id_resource.id
  http_method = aws_api_gateway_method.menu_files_method_get_by_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu-files/restaurant/{restaurant_id} - Get menu files by restaurant id
resource "aws_api_gateway_integration" "menu_files_integration_get_by_restaurant_id" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_files_restaurant_id_resource.id
  http_method = aws_api_gateway_method.menu_files_method_get_by_restaurant_id.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Integration for /api/v1/menu-files/{id} - Update menu file
resource "aws_api_gateway_integration" "menu_files_integration_put" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  resource_id = aws_api_gateway_resource.menu_files_id_resource.id
  http_method = aws_api_gateway_method.menu_files_method_put.http_method

  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.restaurant_menu_service.invoke_arn
}

# Invoke permission for API Gateway to call Lambda function
resource "aws_lambda_permission" "restaurant_menu_gateway_invoke_restaurant_menu_service" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.restaurant_menu_service.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.restaurant_menu_api_gateway.execution_arn}/*/*/*"
}
