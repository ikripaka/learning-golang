package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var INCORRECTARGSQUANTITY = errors.New("Incorrect args quantity")
var NOSUCHELEMENTWITHTHISKEY = errors.New("No Element with such key")
var UNKNOWNERROR = errors.New("I don't know this command")

const idleTimeout = time.Duration(2) //minutes

type server struct {
	commands chan Command
	mutex    sync.Mutex
	hashMap  map[string]Element
}

func newServer() *server {
	s := &server{
		commands: make(chan Command),
		mutex:    sync.Mutex{},
		hashMap:  make(map[string]Element),
	}
	return s
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
		c.err(&INCORRECTARGSQUANTITY)
		return
	}
	s.mutex.Lock()
	element, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	if ok {
		c.write(element.data)
	} else {
		c.err(&NOSUCHELEMENTWITHTHISKEY)
	}
}

func (s *server) set(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_SET_ARGS_NUM, *commandArgs); err != nil {
		c.err(&INCORRECTARGSQUANTITY)
		return
	}

	s.mutex.Lock()
	s.hashMap[(*commandArgs)[0]] = Element{
		mutex: &sync.Mutex{},
		data:  &(*commandArgs)[1],
	}
	s.mutex.Unlock()

Save("/home/ikripaka/go/src/github.com/ikripaka/learning-golang/kv-store/data.json", s, s.hashMap)
	c.write(&(*commandArgs)[1])
}

func (s *server) getset(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_GETSET_ARGS_NUM, *commandArgs); err != nil {
		c.err(&INCORRECTARGSQUANTITY)
		return
	}

	s.mutex.Lock()
	oldElement, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	if ok {
		s.hashMap[(*commandArgs)[0]].mutex.Lock()
		v, _ := s.hashMap[(*commandArgs)[0]]
		v.data = &(*commandArgs)[1]
		s.hashMap[(*commandArgs)[0]].mutex.Unlock()

		c.write(oldElement.data)
	} else {
		c.err(&NOSUCHELEMENTWITHTHISKEY)
		s.set(commandArgs, c)
	}
}

func (s *server) exists(commandArgs *[]string, c *client) {
	if err := validateArgs(CMD_EXISTS_ARGS_NUM, *commandArgs); err != nil {
		c.err(&INCORRECTARGSQUANTITY)
		return
	}
	// lock all hash map
	s.mutex.Lock()
	_, ok := s.hashMap[(*commandArgs)[0]]
	s.mutex.Unlock()

	formattedBool := strconv.FormatBool(ok)
	if ok {
		c.write(&formattedBool)
	} else {
		c.write(&formattedBool)
	}
}

func validateArgs(argsNum int, flags []string) error {
	if len(flags) < argsNum || len(flags) > argsNum {
		return INCORRECTARGSQUANTITY
	}
	return nil
}