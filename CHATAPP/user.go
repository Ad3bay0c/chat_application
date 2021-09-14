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
	user.writeMessage(user, fmt.Sprintf("Welcome to Adebayo Chat App, " +
		"Please Update your username and continue with other operations..."))
	for {
		input, err := bufio.NewReader(user.conn).ReadString('\n')
		if err != nil {
			user.writeMessage(user, fmt.Sprintf("Error Reading strings: %v", err))
			continue
		}

		input = strings.Trim(input, "\n")
		args := strings.Split(input, " ")
		command := strings.ToLower(strings.TrimSpace(args[0]))

		switch command {
		case "*username":
			s.Instructions <- &Instruction{
				command: USERNAME,
				user:    user,
				input:   args,
			}
		case "*join":
			s.Instructions <- &Instruction{
				command: JOIN,
				user:    user,
				input:   args,
			}
		case "*chats":
			s.Instructions <- &Instruction{
				command: CHATS,
				user:    user,
				input:   args,
			}
		case "*send":
			s.Instructions <- &Instruction{
				command: SEND,
				user:    user,
				input:   args,
			}
		case "*quit":
			s.Instructions <- &Instruction{
				command: QUIT,
				user:    user,
				input:   args,
			}
		default:
			user.errorMessage(fmt.Sprintf("invalid command; Choose from the commands below: " +
				"\n\t *username 'your name'-> to set your username," +
				"\n\t *Chats -> to list all available chat Chats," +
				"\n\t *join 'Chats'-> to join/create a chat," +
				"\n\t *send 'message'-> send a message to a chat," +
				"\n\t *quit -> to disconnect from a connection",
			))
		}
	}
}

func (user *User) quitGroup() {
	user.chat.broadcast(user, fmt.Sprintf("%v left the chat group", user.username))

	delete(user.chat.members, user.conn.RemoteAddr())
	user.chat = nil
}

func (user *User) writeMessage(u *User, msg string) {
	user.conn.Write([]byte(fmt.Sprintf("$%v: %s\n", u.username, msg)))
}

func(user *User) errorMessage(msg string) {
	user.conn.Write([]byte(fmt.Sprintf("OOPS!!! An Error Occured: \n\t%v\n",msg)))
}