# resource "aws_db_instance" "res_menu_db" {
#   storage_encrypted       = true
#   allocated_storage       = 20
#   storage_type            = "gp2"
#   engine                  = "postgres"
#   engine_version          = "15.0"
#   instance_class          = "db.t2.micro"
#   db_name                 = var.db_name
#   username                = var.db_user
#   password                = var.db_password
#   skip_final_snapshot     = true
#   backup_retention_period = 7
#   multi_az                = false
#   publicly_accessible     = false


#   tags = {
#     Name        = "res_menu_db"
#     Environment = "dev"
#   }
# }

# output "db_endpoint" {
#   value = aws_db_instance.res_menu_db.endpoint
# }

# output "db_instance_id" {
#   value = aws_db_instance.res_menu_db.id
# }
