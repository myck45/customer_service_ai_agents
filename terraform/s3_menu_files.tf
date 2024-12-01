resource "aws_s3_bucket" "restaurant_menu_bucket" {
  bucket = "restaurant-menu-files-${random_string.bucket_suffix.result}"

  tags = {
    Name        = "Restaurant Menu Files Bucket"
    Environment = "Development"
  }
}

resource "random_string" "bucket_suffix" {
  length  = 8
  special = false
  upper   = false
}
