package CHATAPP

import (
	"fmt"
	"log"
	"net"
	"strings"
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
			s.updateUsername(v.user, v.input)
		case JOIN:
		case SEND:
		case CHATS:
		case QUIT:
			s.quitConnection(v.user)
		}
	}
}

func (s *Server) updateUsername(user *User, args []string) {
	if len(args) < 2 {
		user.writeMessage(user, fmt.Sprintf("Enter a New Username; (*username doe)"))
		return
	}
	username := strings.TrimSpace(args[1])

	user.username = username

	user.writeMessage(user, fmt.Sprintf("Username Updated to %s", username))

}

func (s *Server) joinGroup(user *User, args []string) {
	if len(args) < 2 {
		user.writeMessage(user, fmt.Sprintf("Enter a Name of the group to join or create new one; (*join sport)"))
		return
	}
	if user != nil {
		user.quitGroup()
	}

	groupName := strings.TrimSpace(args[0])

	grp, ok := s.chats[groupName]

	if !ok {
		grp = &Chat{
			name:    groupName,
			members: make(map[net.Addr]*User),
		}
		s.chats[groupName] = grp
	}

	grp.members[user.conn.RemoteAddr()] = user

	user.chat = grp

	user.chat.broadcast(user, fmt.Sprintf("%v joined the group", user.username))

	user.writeMessage(user, fmt.Sprintf("%v welcome to the group", user.username))
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