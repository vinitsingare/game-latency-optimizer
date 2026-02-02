package config

type Relay struct {
	Name string
	Addr string
}

func DefaultRelays() []Relay {
	return []Relay{
		{Name: "Relay-A", Addr: "127.0.0.1:9000"},
		{Name: "Relay-B", Addr: "127.0.0.1:9001"},
	}
}
