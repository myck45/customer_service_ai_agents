variable "db_host" {
  description = "The hostname of the database"
  type        = string
  sensitive   = true
}

variable "db_port" {
  description = "The port of the database"
  type        = string
  sensitive   = true
}

variable "db_user" {
  description = "The username of the database"
  type        = string
  sensitive   = true
}

variable "db_name" {
  description = "The name of the database"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "The password of the database"
  type        = string
  sensitive   = true
}

variable "openai_api_key" {
  description = "The API key for OpenAI"
  type        = string
  sensitive   = true
}

variable "jwt_secret" {
  description = "The secret for JWT"
  type        = string
  sensitive   = true
}

variable "twilio_account_sid" {
  description = "The account SID for Twilio"
  type        = string
  sensitive   = true
}

variable "twilio_auth_token" {
  description = "The auth token for Twilio"
  type        = string
  sensitive   = true
}

variable "supabase_url" {
  description = "The URL for Supabase"
  type        = string
  sensitive   = true
}

variable "supabase_key" {
  description = "The key for Supabase"
  type        = string
  sensitive   = true
}
