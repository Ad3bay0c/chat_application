package CHATAPP

import "net"

type User struct {
	username	string
	chat	*Chat
	conn	net.Conn
}

func (user *User) readInput(s *Server) {

}
