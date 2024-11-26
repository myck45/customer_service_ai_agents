resource "aws_lambda_function" "user_service" {
  filename      = "user_service.zip"
  function_name = "user_service"
  role          = aws_iam_role.restaurant_menu_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  memory_size   = 128
  timeout       = 160

  source_code_hash = filebase64sha256("user_service.zip")

  environment {
    variables = {
      DB_HOST            = var.db_host
      DB_PORT            = var.db_port
      DB_USER            = var.db_user
      DB_NAME            = var.db_name
      DB_PASSWORD        = var.db_password
      OPENAI_API_KEY     = var.openai_api_key
      JWT_SECRET         = var.jwt_secret
      TWILIO_ACCOUNT_SID = var.twilio_account_sid
      TWILIO_AUTH_TOKEN  = var.twilio_auth_token
      SUPABASE_URL       = var.supabase_url
      SUPABASE_KEY       = var.supabase_key
    }
  }
}

resource "aws_lambda_function" "restaurant_menu_service" {
  filename      = "restaurant_menu_service.zip"
  function_name = "restaurant_menu_service"
  role          = aws_iam_role.restaurant_menu_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  memory_size   = 128
  timeout       = 160

  source_code_hash = filebase64sha256("restaurant_menu_service.zip")

  environment {
    variables = {
      DB_HOST            = var.db_host
      DB_PORT            = var.db_port
      DB_USER            = var.db_user
      DB_NAME            = var.db_name
      DB_PASSWORD        = var.db_password
      OPENAI_API_KEY     = var.openai_api_key
      JWT_SECRET         = var.jwt_secret
      TWILIO_ACCOUNT_SID = var.twilio_account_sid
      TWILIO_AUTH_TOKEN  = var.twilio_auth_token
      SUPABASE_URL       = var.supabase_url
      SUPABASE_KEY       = var.supabase_key
    }
  }
}

resource "aws_lambda_function" "bot_service" {
  filename      = "bot_service.zip"
  function_name = "bot_service"
  role          = aws_iam_role.restaurant_menu_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  memory_size   = 128
  timeout       = 160

  source_code_hash = filebase64sha256("bot_service.zip")

  environment {
    variables = {
      DB_HOST            = var.db_host
      DB_PORT            = var.db_port
      DB_USER            = var.db_user
      DB_NAME            = var.db_name
      DB_PASSWORD        = var.db_password
      OPENAI_API_KEY     = var.openai_api_key
      JWT_SECRET         = var.jwt_secret
      TWILIO_ACCOUNT_SID = var.twilio_account_sid
      TWILIO_AUTH_TOKEN  = var.twilio_auth_token
      SUPABASE_URL       = var.supabase_url
      SUPABASE_KEY       = var.supabase_key
    }
  }
}
