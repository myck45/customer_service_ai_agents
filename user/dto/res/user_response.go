package res

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	UserEmail string `json:"user_email"`
	PhoneNum  string `json:"phone_num"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
