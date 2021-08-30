package CHATAPP

import (
	"fmt"
	"net"
)

type User struct {
	username	string
	chat	*Chat
	conn	net.Conn
}

func (user *User) readInput(s *Server) {
	user.writeMessage(fmt.Sprintf("Welcome to Adebayo Chat App, " +
		"Please Update your username and continue with other operations...\n"))
	//for {
	//	input, err := bufio.NewReader(user.conn).ReadString('\n')
	//	checkError(err, fmt.Sprintf("Error Reading strings: %v", err))
	//}
}

func (user *User) writeMessage(msg string) {
	user.conn.Write([]byte(msg))
}