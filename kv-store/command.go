package main

type commandId int

//Possible commands id
const (
	CMD_GET commandId = iota
	CMD_SET
	CMD_GETSET
	CMD_EXISTS
	CMD_EXISTS_ARGS_NUM = 1
	CMD_GET_ARGS_NUM
	CMD_SET_ARGS_NUM = 2
	CMD_GETSET_ARGS_NUM
)

//Struct represent command structure
type Command struct {
	id     commandId
	flags  []string
	client *client
}
