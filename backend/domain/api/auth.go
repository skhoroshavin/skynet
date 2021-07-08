package api

type Auth interface {
	CreateSession(id string) (string, error)
	InvalidateSession(sid string)

	UserFromSession(sid string) (string, error)
}
