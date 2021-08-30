package CHATAPP

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type User struct {
	username	string
	chat	*Chat
	conn	net.Conn
}

func (user *User) readInput(s *Server) {
	user.writeMessage(fmt.Sprintf("Welcome to Adebayo Chat App, " +
		"Please Update your username and continue with other operations...\n"))
	for {
		input, err := bufio.NewReader(user.conn).ReadString('\n')
		if err != nil {
			user.writeMessage(fmt.Sprintf("Error Reading strings: %v", err))
			continue
		}

		input = strings.Trim(input, "\n")
		args := strings.Split(input, " ")
		command := strings.TrimSpace(args[0])

		switch command {
		case "*username":
			s.instructions <- &Instruction{
				command: USERNAME,
				user:    user,
				input:   args,
			}
		case "*join":
			s.instructions <- &Instruction{
				command: JOIN,
				user:    user,
				input:   args,
			}
		case "*chats":
			s.instructions <- &Instruction{
				command: CHATS,
				user:    user,
				input:   args,
			}
		case "*send":
			s.instructions <- &Instruction{
				command: SEND,
				user:    user,
				input:   args,
			}
		case "*quit":
			s.instructions <- &Instruction{
				command: QUIT,
				user:    user,
				input:   args,
			}
		default:
			user.errorMessage(fmt.Sprintf("invalid command; Choose from the commands below: " +
				"\n\t *username 'your name'-> to set your username," +
				"\n\t *chats -> to list all available chat chats," +
				"\n\t *join 'chats'-> to join/create a chat," +
				"\n\t *send 'message'-> send a message to a chat," +
				"\n\t *quit -> to disconnect from a connection",
			))
		}
	}
}

func (user *User) writeMessage(msg string) {
	user.conn.Write([]byte(msg))
}

func(user *User) errorMessage(msg string) {
	user.conn.Write([]byte("OOPS!!! An Error Occured: \n\t"+msg))
}