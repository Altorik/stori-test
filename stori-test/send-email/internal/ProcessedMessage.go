package internal

type MessageId struct {
	value string
}

type DateEvent struct {
	value string
}

func (date DateEvent) String() string {
	return date.value
}
func (id MessageId) String() string {
	return id.value
}

type ProcessMessage struct {
	MessageId string
	DateEvent string
}

type ProcessMessageRepositorio interface {
	Guardar(message ProcessMessage) error
	Buscar(message ProcessMessage) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/mocks --name=ProcessMessageRepositorio
