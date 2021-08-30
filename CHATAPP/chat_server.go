package CHATAPP

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	chats			map[string]*Chat
	instructions	chan *Instruction
}

func StartServer() {
	s := &Server{
		chats: make(map[string]*Chat),
		instructions: make(chan *Instruction),
	}

	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		panic(err)
	}

	for {
		log.Printf("Server Started at localhost:3333")

		conn, err := listener.Accept()
		checkError(err, fmt.Sprintf("Error Accepting Request: %v", err))

		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	log.Printf("A new Connection to the Server : %s", conn.RemoteAddr().String())

	newUser := &User{
		conn: conn,
		username: "anonymous",
	}

	newUser.readInput(s)

}
func checkError(err error, msg string) {
	if err != nil {
		log.Printf(msg)
	}
}