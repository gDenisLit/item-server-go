package models

type ServerErr struct {
	Message string
}

type ClientErr struct {
	Message string
}

func (e *ServerErr) Error() string {
	return e.Message
}

func (e *ClientErr) Error() string {
	return e.Message
}
