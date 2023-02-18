package com.tech.spring.reactive.service;

import com.tech.spring.reactive.entity.Message;
import com.tech.spring.reactive.repository.MessageRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

import java.util.Date;

@Service
public class MessageService {

    @Autowired
    MessageRepository messageRepository;
    
    public Mono<Message> saveMessage(String text) {
        Message message = new Message();
        message.setCreateTime(new Date());
        message.setContent(text);
        return messageRepository.save(message);
    }
}
