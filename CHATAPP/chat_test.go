package CHATAPP

import (
	"net"
	"testing"
)

func TestServer_StartServer(t *testing.T) {
	StartServer()
}

func TestStartServer(t *testing.T) {
	tables := []struct{
		output	int
		word	string
	}{
		{
			output: 1,
			word: "First Client",
		},
		{
			output: 2,
			word: "Second Client",
		},
	}


	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			conn, err := net.Dial("tcp", "localhost:3333")
			if err != nil {
				t.Errorf("Error")
				conn.Close()
			}
		})
	}
}
