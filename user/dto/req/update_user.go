package req

type UpdateUserReq struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100"`
	BirthDate string `json:"birth_date" binding:"required"`
	UserEmail string `json:"user_email" binding:"required,email"`
	PhoneNum  string `json:"phone_num" binding:"required,min=10,max=20"`
	Password  string `json:"password" binding:"required,min=8,max=255"`
}
