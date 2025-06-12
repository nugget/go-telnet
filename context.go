package telnet

import (
	"net"
)

type Context interface {
	Logger() Logger
	InjectLogger(Logger) Context

	Conn() net.Conn
	InjectConn(net.Conn) Context
}

type internalContext struct {
	conn   net.Conn
	logger Logger
}

func NewContext() Context {
	ctx := internalContext{}

	return &ctx
}

func (ctx *internalContext) Logger() Logger {
	return ctx.logger
}

func (ctx *internalContext) InjectLogger(logger Logger) Context {
	ctx.logger = logger

	return ctx
}

func (ctx *internalContext) Conn() net.Conn {
	return ctx.conn
}

func (ctx *internalContext) InjectConn(conn net.Conn) Context {
	ctx.conn = conn

	return ctx
}
