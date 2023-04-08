package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/tarantool/go-tarantool"
)

const VinylSpaceName = "test_space_vinyl"

func BenchmarkWrite10kToSpace(b *testing.B) {
	conn, teardown := setupTesting(b)
	b.Cleanup(teardown)

	for i := 0; i < 10000; i++ {
		conn.InsertAsync(VinylSpaceName, []interface{}{
			fmt.Sprintf("google%d", i),
			"new_type",
			1,
		})
	}
}

func BenchmarkRead10kFromSpace(b *testing.B) {
	conn, teardown := setupTesting(b)
	b.Cleanup(teardown)

	for i := 0; i < 10000; i++ {
		_, _ = conn.Select(VinylSpaceName, "domain", 0, 1, 0, []interface{}{
			fmt.Sprintf("google%d", i),
		})
	}
}

func setupTesting(b *testing.B) (*tarantool.Connection, func()) {

	b.Helper()

	connect := func() (*tarantool.Connection, error) {
		return tarantool.Connect("localhost:3306", tarantool.Opts{
			Reconnect: 1 * time.Second,
			User:      "user",
			Pass:      "password",
		})
	}

	conn, err := connect()

	if err != nil {
		fmt.Println(err.Error())
	}

	//_, _ = conn.Call("box.space."+VinylSpaceName+":drop", []interface{}{})

	script, err := os.ReadFile("tarantool/init.lua")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = conn.Eval(string(script), []interface{}{})
	if err != nil {
		fmt.Println(err.Error())
	}

	return conn, func() {
		_ = conn.Close()
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
