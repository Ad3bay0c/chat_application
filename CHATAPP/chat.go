package CHATAPP

import "net"

type Chat struct {
	name	string
	members	map[net.Addr]*User
}
