package main

import "fmt"

type Postgres struct {
	host string
	port string
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		fmt.Println(opt)
		opt(p)
	}
	return p, nil
}

func main() {
	p, err := New(WithHost("1"), WithPort("2"))
	if err != nil {
		return
	}

	fmt.Println(p)
}

type Option func(*Postgres)

func WithHost(p Postgres, host string) {
	p.host = host
	return
}

func WithPort(port string) Option {
	return func(postgres *Postgres) {
		postgres.port = port
	}
}
