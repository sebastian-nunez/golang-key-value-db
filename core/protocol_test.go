package core

import (
	"errors"
	"testing"
)

func TestProtocol(t *testing.T) {
	t.Parallel()

	t.Run("isValidParams", func(t *testing.T) {
		t.Run("should return true given valid params for a command", func(t *testing.T) {
			cmd := CmdGet
			params := []string{"some-key"}

			got := cmd.isValidParams(params)

			if !got {
				t.Errorf("Expected true, got false")
			}
		})

		t.Run("should return false given invalid params for a command", func(t *testing.T) {
			cmd := CmdGet
			params := []string{"some-key", "some-other-key"}

			got := cmd.isValidParams(params)

			if !got {
				t.Errorf("Expected false, got true")
			}
		})
	})

	t.Run("parseCommand", func(t *testing.T) {
		t.Run("should return CmdGet", func(t *testing.T) {
			got, err := parseCommand("GET")

			if got != CmdGet {
				t.Errorf("Expected %s, got %s", CmdGet.CmdStr, got.CmdStr)
			}
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
		})

		t.Run("should return CmdSet", func(t *testing.T) {
			got, err := parseCommand("SET")

			if got != CmdSet {
				t.Errorf("Expected %s, got %s", CmdSet.CmdStr, got.CmdStr)
			}
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
		})

		t.Run("should return CmdDelete", func(t *testing.T) {
			got, err := parseCommand("DELETE")

			if got != CmdDelete {
				t.Errorf("Expected %s, got %s", CmdDelete.CmdStr, got.CmdStr)
			}
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
		})

		t.Run("should return CmdPing", func(t *testing.T) {
			got, err := parseCommand("PING")

			if got != CmdPing {
				t.Errorf("Expected %s, got %s", CmdPing.CmdStr, got.CmdStr)
			}
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
		})

		t.Run("should return empty/invalid command", func(t *testing.T) {
			got, err := parseCommand("INVALID")

			if got != (Command{}) {
				t.Errorf("Expected empty command, got %s", got.CmdStr)
			}
			if !errors.Is(err, ErrInvalidCommand) {
				t.Errorf("Expected %s, got %s", ErrInvalidCommand, err)
			}
		})
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("should return error given invalid command", func(t *testing.T) {
			_, err := ParseProtocol("")

			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !errors.Is(err, ErrInvalidCommand) {
				t.Errorf("Expected %s, got %s", ErrInvalidCommand, err)
			}
		})

		t.Run("parses valid command", func(t *testing.T) {
			req, err := ParseProtocol("GET key")

			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
			if req.Command != CmdGet {
				t.Errorf("Expected %s, got %s", CmdGet.CmdStr, req.Command.CmdStr)
			}
			if len(req.Params) != 1 {
				t.Errorf("Expected 1, got %d", len(req.Params))
			}
			if req.Params[0] != "key" {
				t.Errorf("Expected key, got %s", req.Params[0])
			}
		})

		t.Run("parses valid SET command with an optional TTL", func(t *testing.T) {
			req, err := ParseProtocol("SET key value 100")

			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
			if req.Command != CmdSet {
				t.Errorf("Expected %s, got %s", CmdSet.CmdStr, req.Command.CmdStr)
			}
			if len(req.Params) != 3 {
				t.Errorf("Expected 3, got %d", len(req.Params))
			}
			if req.Params[0] != "key" {
				t.Errorf("Expected key, got %s", req.Params[0])
			}
			if req.Params[1] != "value" {
				t.Errorf("Expected value, got %s", req.Params[1])
			}
			if req.Params[2] != "100" {
				t.Errorf("Expected 100, got %s", req.Params[2])
			}
		})
	})
}
