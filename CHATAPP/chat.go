package CHATAPP

import "net"

type Chat struct {
	name	string
	members	map[net.Addr]*User
}

func (chat *Chat) broadcast(user *User, msg string) {
	for userAddr, u := range chat.members {
		if userAddr != user.conn.RemoteAddr() {
			if user == nil {
				u.writeMessage(nil, msg)
			} else {
				u.writeMessage(user, msg)
			}

		}
	}
}
