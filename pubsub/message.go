package pubsub

type Message interface {
	GetData() ([]byte, error)
	GetAttributes() (map[string]string, error)
}
