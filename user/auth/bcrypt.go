package auth

type Bcrypt interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
