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

	go s.readInstruction()

	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		panic(err)
	}
	log.Printf("Server Started at localhost:3333")
	for {

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
			s.joinGroup(v.user, v.input)
		case SEND:
			s.sendMessage(v.user, v.input)
		case CHATS:
			s.chatsList(v.user)
		case QUIT:
			s.quitConnection(v.user)
		}
	}
}

func (s *Server) updateUsername(user *User, args []string) {
	if len(args) < 2 {
		user.errorMessage(fmt.Sprintf("Enter a New Username; (*username doe)"))
		return
	}
	username := strings.TrimSpace(args[1])

	user.username = username

	user.writeMessage(user, fmt.Sprintf("Username Updated to %s", username))

}

func (s *Server) joinGroup(user *User, args []string) {
	if len(args) < 2 {
		user.errorMessage(fmt.Sprintf("Enter a group Name to join or create new one; (*join sport)"))
		return
	}
	if user.chat != nil {
		user.quitGroup()
	}

	groupName := strings.TrimSpace(args[1])

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

	user.writeMessage(user, fmt.Sprintf("welcome to the group %v", groupName))
}

func (s *Server) sendMessage(user *User, args []string) {
	if len(args) < 2 {
		user.errorMessage(fmt.Sprintf("Type in a message ; (*send Hi, How are you doing)"))
		return
	}
	if user.chat == nil {
		user.errorMessage(fmt.Sprintf("You belong to no group, Join a group first; (*join sport)"))
		return
	}

	msg := strings.TrimSpace(strings.Join(args[1:], " "))
	user.chat.broadcast(user, msg)
}

func(s *Server) chatsList(user *User) {
	if len(s.chats) == 0 {
		user.writeMessage(user, fmt.Sprintf("Empty, create new chat group (*join sport)"))
		return
	}
	list := make([]string, 0)
	for name, _ := range s.chats {
		list = append(list, name)
	}
	user.writeMessage(user, fmt.Sprintf("Chat Groups: %s", strings.Join(list, ", ")))
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