package core

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidCommand = errors.New("invalid command")

// Request serves as the parsed TCP request payload that the user attempts to execute.
type Request struct {
	Command Command
	Params  []string
}

// Response is returned by the server after each request execution.
type Response struct {
	Success bool
	Value   []byte
}

type Command struct {
	CmdStr            string
	MinRequiredParams int
}

func (c Command) isValidParams(params []string) bool {
	return len(params) >= c.MinRequiredParams
}

// List of available/support commands.
var (
	CmdGet    Command = Command{CmdStr: "GET", MinRequiredParams: 1}
	CmdSet    Command = Command{CmdStr: "SET", MinRequiredParams: 2}
	CmdDelete Command = Command{CmdStr: "DELETE", MinRequiredParams: 1}
	CmdPing   Command = Command{CmdStr: "PING", MinRequiredParams: 0}
)

// ParseProtocol parses an input string and attempts to extract the command and params to build a `Request` object.
// General command definition: `<command_type> <param1> <param2>...`
func ParseProtocol(input string) (Request, error) {
	tokens := strings.Split(input, " ")
	if len(tokens) == 0 {
		return Request{}, fmt.Errorf("empty request")
	}

	cmd, err := parseCommand(tokens[0])
	if err != nil {
		return Request{}, err
	}

	params := tokens[1:]
	if !cmd.isValidParams(params) {
		return Request{}, fmt.Errorf("invalid params for command '%s'. Expected '%d' and got '%d'", cmd.CmdStr, cmd.MinRequiredParams, len(params))
	}

	return Request{Command: cmd, Params: params}, nil
}

func parseCommand(cmd string) (Command, error) {
	switch cmd {
	case CmdGet.CmdStr:
		return CmdGet, nil
	case CmdSet.CmdStr:
		return CmdSet, nil
	case CmdDelete.CmdStr:
		return CmdDelete, nil
	case CmdPing.CmdStr:
		return CmdPing, nil
	default:
		return Command{}, ErrInvalidCommand
	}
}
