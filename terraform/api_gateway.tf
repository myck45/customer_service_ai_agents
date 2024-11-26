# API GATEWAY
resource "aws_api_gateway_rest_api" "restaurant_menu_api_gateway" {
  name        = "restaurant_menu_api_gateway"
  description = "API Gateway for Restaurant Menu"
}

# Base path mapping /api
resource "aws_api_gateway_resource" "api" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.root_resource_id
  path_part   = "api"
}

# Base path mapping /api/v1
resource "aws_api_gateway_resource" "v1" {
  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  parent_id   = aws_api_gateway_resource.api.id
  path_part   = "v1"
}

# Gateway Deployment
resource "aws_api_gateway_deployment" "restaurant_menu_api_gateway_deployment" {
  depends_on = [
    aws_api_gateway_integration.user_integration_post,
    aws_api_gateway_integration.user_integration_delete,
    aws_api_gateway_integration.user_integration_get_all,
    aws_api_gateway_integration.user_integration_get_by_id,
    aws_api_gateway_integration.user_integration_get_by_email,
    aws_api_gateway_integration.user_integration_put,
    aws_api_gateway_integration.user_integration_post_login,
    aws_api_gateway_integration.restaurant_integration_post,
    aws_api_gateway_integration.restaurant_integration_delete,
    aws_api_gateway_integration.restaurant_integration_get_all,
    aws_api_gateway_integration.restaurant_integration_get_by_id,
    aws_api_gateway_integration.restaurant_integration_put,
    aws_api_gateway_integration.menu_integration_delete,
    aws_api_gateway_integration.menu_integration_get_all,
    aws_api_gateway_integration.menu_integration_search,
    aws_api_gateway_integration.menu_integration_get_by_id,
    aws_api_gateway_integration.menu_integration_put,
    aws_api_gateway_integration.bot_integration_post,
    aws_api_gateway_integration.bot_integration_delete,
    aws_api_gateway_integration.bot_integration_get_all,
    aws_api_gateway_integration.bot_integration_get_by_id,
    aws_api_gateway_integration.bot_integration_get_by_restaurant_id,
    aws_api_gateway_integration.bot_integration_get_by_whatsapp,
    aws_api_gateway_integration.bot_integration_put,
  ]

  rest_api_id = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  description = "Deployment for Restaurant Menu API Gateway"

  triggers = {
    redeployment = sha1(jsonencode({
      user_integration_post                = aws_api_gateway_integration.user_integration_post.id,
      user_integration_delete              = aws_api_gateway_integration.user_integration_delete.id,
      user_integration_get_all             = aws_api_gateway_integration.user_integration_get_all.id,
      user_integration_get_by_id           = aws_api_gateway_integration.user_integration_get_by_id.id,
      user_integration_get_by_email        = aws_api_gateway_integration.user_integration_get_by_email.id,
      user_integration_put                 = aws_api_gateway_integration.user_integration_put.id,
      user_integration_post_login          = aws_api_gateway_integration.user_integration_post_login.id,
      restaurant_integration_post          = aws_api_gateway_integration.restaurant_integration_post.id,
      restaurant_integration_delete        = aws_api_gateway_integration.restaurant_integration_delete.id,
      restaurant_integration_get_all       = aws_api_gateway_integration.restaurant_integration_get_all.id,
      restaurant_integration_get_by_id     = aws_api_gateway_integration.restaurant_integration_get_by_id.id,
      restaurant_integration_put           = aws_api_gateway_integration.restaurant_integration_put.id,
      menu_integration_delete              = aws_api_gateway_integration.menu_integration_delete.id,
      menu_integration_get_all             = aws_api_gateway_integration.menu_integration_get_all.id,
      menu_integration_search              = aws_api_gateway_integration.menu_integration_search.id,
      menu_integration_get_by_id           = aws_api_gateway_integration.menu_integration_get_by_id.id,
      menu_integration_put                 = aws_api_gateway_integration.menu_integration_put.id,
      bot_integration_post                 = aws_api_gateway_integration.bot_integration_post.id,
      bot_integration_delete               = aws_api_gateway_integration.bot_integration_delete.id,
      bot_integration_get_all              = aws_api_gateway_integration.bot_integration_get_all.id,
      bot_integration_get_by_id            = aws_api_gateway_integration.bot_integration_get_by_id.id,
      bot_integration_get_by_restaurant_id = aws_api_gateway_integration.bot_integration_get_by_restaurant_id.id,
      bot_integration_get_by_whatsapp      = aws_api_gateway_integration.bot_integration_get_by_whatsapp.id,
      bot_integration_put                  = aws_api_gateway_integration.bot_integration_put.id,
    }))
  }

  lifecycle {
    create_before_destroy = true
  }
}

# Gateway Stage
resource "aws_api_gateway_stage" "restaurant_menu_api_gateway_stage" {
  deployment_id = aws_api_gateway_deployment.restaurant_menu_api_gateway_deployment.id
  rest_api_id   = aws_api_gateway_rest_api.restaurant_menu_api_gateway.id
  stage_name    = "dev"
}

# Invocation URL
output "restaurant_menu_api_gateway_url" {
  value = "https://${aws_api_gateway_rest_api.restaurant_menu_api_gateway.id}.execute-api.sa-east-1.amazonaws.com/${aws_api_gateway_stage.restaurant_menu_api_gateway_stage.stage_name}/api/v1"
}
