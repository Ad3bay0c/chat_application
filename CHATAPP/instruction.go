package CHATAPP

type Instruction struct {
	command		int
	user		*User
	input		[]string
}

const (
	USERNAME = iota + 1
	JOIN
	CHATS
	SEND
	QUIT
)
