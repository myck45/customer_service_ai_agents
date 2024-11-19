package auth

type Auth interface {
	GenerateToken(id uint, email string, role string) (string, error)
}
