package core

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sebastian-nunez/golang-key-value-db/store"
)

type RequestProcessor interface {
	Process(context.Context, Request) (Response, error)
}

type CommandProcessor struct {
	store store.DataStore[string, []byte]
}

func NewCommandProcessor(s store.DataStore[string, []byte]) *CommandProcessor {
	return &CommandProcessor{
		store: s,
	}
}

func (cp *CommandProcessor) Process(ctx context.Context, req Request) (Response, error) {
	switch req.Command {
	case CmdGet:
		return cp.processGet(req)
	case CmdSet:
		return cp.processSet(ctx, req)
	case CmdDelete:
		return cp.processDelete(req)
	case CmdPing:
		return cp.processPing()
	default:
		return Response{}, ErrInvalidCommand
	}
}

func (cp *CommandProcessor) processGet(req Request) (Response, error) {
	key := req.Params[0]
	res, err := cp.store.Get(key)
	if err != nil {
		return Response{}, err
	}

	return Response{Success: true, Value: res}, nil
}

func (cp *CommandProcessor) processSet(ctx context.Context, req Request) (Response, error) {
	key := req.Params[0]
	val := req.Params[1]
	cp.store.Set(key, []byte(val))

	if len(req.Params) > 2 {
		ttl, err := toSeconds(req.Params[2])
		if err != nil {
			// Delete the key whenever the TTL can not be set.
			cp.store.Delete(key)
			return Response{}, err
		}

		go cp.processExpiry(ctx, key, ttl)
	}

	return Response{Success: true}, nil
}

func (cp *CommandProcessor) processDelete(req Request) (Response, error) {
	key := req.Params[0]
	cp.store.Delete(key)
	return Response{Success: true}, nil
}

func (cp *CommandProcessor) processPing() (Response, error) {
	return Response{Success: true, Value: []byte("PONG")}, nil
}

func (cp *CommandProcessor) processExpiry(ctx context.Context, key string, ttl time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Printf("Canceling TTL for key '%s' due to server shutdown.\n", key)
		return
	default:
		// Uses a blocking timer to pause execution until the TTL expires,
		// after which the key is removed from the store.
		<-time.After(ttl)
		cp.store.Delete(key)
	}
}

func toSeconds(ttl string) (time.Duration, error) {
	ttlInt, err := strconv.Atoi(ttl)
	if err != nil {
		return 0, err
	}

	return time.Duration(ttlInt) * time.Second, nil
}
