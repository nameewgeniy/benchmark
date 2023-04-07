package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/tarantool/go-tarantool"
)

type Tarantool struct {
	Host     string
	User     string
	Password string
}

func initTarantoolClient(cfg *Tarantool) (*tarantool.Connection, error) {
	connect := func() (*tarantool.Connection, error) {
		return tarantool.Connect(cfg.Host, tarantool.Opts{
			Reconnect: 1 * time.Second,
			User:      cfg.User,
			Pass:      cfg.Password,
		})
	}

	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("connect to tarantool: %v", err)
	}

	return conn, nil
}

func BenchmarkReadWrite(b *testing.B) {

	tarantoolConn, err := initTarantoolClient(&Tarantool{
		Host:     "localhost:3306",
		User:     "user",
		Password: "password",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() { _ = tarantoolConn.Close() }()

}
