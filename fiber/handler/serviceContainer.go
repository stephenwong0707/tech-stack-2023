package handler

import "store/service"

var messageService service.MessageService

func Init() {
	messageService.Init()
}
