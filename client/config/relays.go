package config

type Relay struct {
	Name string
	Addr string
}

func DefaultRelays() []Relay {
	return []Relay{
		{Name: "Relay-A", Addr: "103.102.46.130:9999"},
		{Name: "Relay-B", Addr: "127.0.0.1:9000"},
	}
}
