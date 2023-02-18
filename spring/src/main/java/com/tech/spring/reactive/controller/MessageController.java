package com.tech.spring.reactive.controller;

import com.tech.spring.reactive.dto.MessageRequest;
import com.tech.spring.reactive.entity.Message;
import com.tech.spring.reactive.service.MessageService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;


@RestController
@RequestMapping("/message")
public class MessageController {

    @Autowired
    MessageService messageService;

    @PostMapping
    public Mono<Message> postMessage(@RequestBody MessageRequest req) {
        return messageService.saveMessage(req.getText());
    }

}