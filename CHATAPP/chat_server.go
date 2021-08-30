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

func (s *Server) readInstruction() {
	for v := range s.instructions {
		switch v.command {
		case USERNAME:
		case JOIN:
		case SEND:
		case CHATS:
		case QUIT:
			s.quitConnection(v.user)
		}
	}
}

func (s *Server) quitConnection(user *User) {

	if user.chat != nil {
		user.quitGroup()
	}
	log.Printf("A Connection Disconnected: %s", user.conn.RemoteAddr().String())

	user.conn.Close()
}
func checkError(err error, msg string) {
	if err != nil {
		log.Printf(msg)
	}
}