package req

type CreateUserReq struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100"`
	BirthDate string `json:"birth_date" binding:"required"`
	UserEmail string `json:"user_email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=255"`
	PhoneNum  string `json:"phone_num" binding:"required,min=10,max=20"`
	Role      string `json:"role" binding:"required,oneof=admin user"`
}
