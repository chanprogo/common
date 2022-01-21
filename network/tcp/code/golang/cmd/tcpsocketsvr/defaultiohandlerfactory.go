package main

import "temporary/pkg/tcpserver"

type DefaultIoHandlerFactory struct {
}

func (defaultFactory *DefaultIoHandlerFactory) CreateIoHandler() tcpserver.IoHandler {
	return new(DefaultIoHandler)
}
