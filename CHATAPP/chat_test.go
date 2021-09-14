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
		word	string
	}{
		{
			word: "First Client",
		},
		{
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
