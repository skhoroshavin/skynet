package api

type Auth interface {
	SignUp(id string, password string) (string, error)
	UserID(session string) (string, error)

	UpdatePassword(id string, oldPassword string, newPassword string) error
}
