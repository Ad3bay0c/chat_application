package CHATAPP

type Server struct {
	chats			map[string]*Chat
	instructions	chan Instruction
}