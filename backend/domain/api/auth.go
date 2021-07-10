package api

type Auth interface {
	CreateUser(id string, password string) error
	UpdatePassword(id string, oldPassword string, newPassword string) error

	//CreateSession(id string) (string, error)
	//InvalidateSession(sid string)
	//
	//UserFromSession(sid string) (string, error)
}
