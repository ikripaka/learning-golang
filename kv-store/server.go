package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

var INCORRECTARGSQUANTITY = errors.New("Incorrect args quantity")
var NOSUCHELEMENTWITHTHISKEY = errors.New("No element with such key")

const idleTimeout = time.Duration(2) //minutes

type server struct {
	commands chan Command
	mutex    sync.Mutex
	hashMap  map[string]string
}

func newServer() *server {
	return &server{
		commands: make(chan Command),
		mutex:    sync.Mutex{},
		hashMap:  make(map[string]string),
	}
}

func (s *server) handleClientRequest(conn *net.Conn) {
	client := &client{
		conn:     conn,
		commands: s.commands,
	}
	log.Println("new user!!")
	client.readUserInput()
}

func (s *server) processCommandsOnServer() {
	for v := range s.commands {
		switch v.id {
		case CMD_SET:
			s.set(&v.flags, v.client)
		case CMD_GET:
			s.get(&v.flags, v.client)
		case CMD_GETSET:
			s.getset(&v.flags, v.client)
		case CMD_EXISTS:
			s.exists(&v.flags, v.client)
		}
	}
}

func (s *server) handleRequest(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("unable to start s: %s", err.Error())
		}

		log.Print(msg)
	}

	conn.Close()
}

func (s *server) get(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_GET_ARGS_NUM, *commandArgs); err != nil {
		c.err(INCORRECTARGSQUANTITY.Error())
		return
	}
	s.mutex.Lock()
	element, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	if ok {
		c.write(element)
	} else {
		c.err(NOSUCHELEMENTWITHTHISKEY.Error())
	}
}

func (s *server) set(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_SET_ARGS_NUM, *commandArgs); err != nil {
		c.err(INCORRECTARGSQUANTITY.Error())
		return
	}
	s.mutex.Lock()
	s.hashMap[(*commandArgs)[0]] = (*commandArgs)[1]
	s.mutex.Unlock()

	c.write((*commandArgs)[1])
}

func (s *server) getset(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_GETSET_ARGS_NUM, *commandArgs); err != nil {
		c.err(INCORRECTARGSQUANTITY.Error())
		return
	}

	s.mutex.Lock()
	oldElement, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	if ok {
		s.mutex.Lock()
		s.hashMap[(*commandArgs)[0]] = (*commandArgs)[1]
		s.mutex.Unlock()

		c.write(oldElement)
	} else {
		c.err(NOSUCHELEMENTWITHTHISKEY.Error())
		s.set(commandArgs, c)
	}
}

func (s *server) exists(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_EXISTS_ARGS_NUM, *commandArgs); err != nil {
		c.err(INCORRECTARGSQUANTITY.Error())
		return
	}
	s.mutex.Lock()
	_, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	if ok {
		c.write("true")
	} else {
		c.write("false")
	}
}

func validateArgs(argsNum int, flags []string) error {
	if len(flags) < argsNum || len(flags) > argsNum {
		return INCORRECTARGSQUANTITY
	}
	return nil
}
