package main

import (
	"time"

	"github.com/tjhorner/makerbotd/api"
)

type mbotCtx struct {
	Client *api.Client
}

func (mbotCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (mbotCtx) Done() <-chan struct{} {
	return nil
}

func (mbotCtx) Err() error {
	return nil
}

func (mbotCtx) Value(key interface{}) interface{} {
	return nil
}

func (mbotCtx) String() string {
	return "mbotCtx"
}
