package CHATAPP

import "net"

type User struct {
	name	string
	chat	*Chat
	conn	net.Conn
}
