
// Basic IAM role for Lambda function, allow collect logs in CloudWatch
resource "aws_iam_role" "restaurant_menu_role" {
  name = "restaurant_menu_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

// Attach AWS managed policy to the role
resource "aws_iam_role_policy_attachment" "restaurant_menu_lambda_policy" {
  role       = aws_iam_role.restaurant_menu_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
