package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type client struct {
	conn     *net.Conn
	commands chan Command
}

func (c *client) readUserInput() {
	for {
		msg, err := bufio.NewReader(*c.conn).ReadString('\n')
		if err != nil {
			log.Println("Error with message:", msg, err)
			break
		}
		msg = strings.Trim(msg, "\r\n")
		splittedMsg := strings.Split(msg, " ")
		log.Println(splittedMsg, len(splittedMsg))

		switch splittedMsg[0] {
		case "get":
			c.commands <- Command{
				id:     CMD_GET,
				flags:  splittedMsg[1:],
				client: c,
			}
		case "set":
			c.commands <- Command{
				id:     CMD_SET,
				flags:  splittedMsg[1:],
				client: c,
			}
		case "exists":
			c.commands <- Command{
				id:     CMD_EXISTS,
				flags:  splittedMsg[1:],
				client: c,
			}
		case "getset":
			c.commands <- Command{
				id:     CMD_GETSET,
				flags:  splittedMsg[1:],
				client: c,
			}
		default:
			c.err(&err)
		}

	}
}

func (c *client) write(message *string) {
	(*c.conn).Write([]byte(*message + "\n"))
}
func (c *client) err(err *error) {
	(*c.conn).Write([]byte("> " + (*err).Error() + "\n"))
}
