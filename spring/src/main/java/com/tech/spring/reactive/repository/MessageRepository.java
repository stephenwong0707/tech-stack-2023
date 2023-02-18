package com.tech.spring.reactive.repository;

import com.tech.spring.reactive.entity.Message;
import org.springframework.data.cassandra.repository.ReactiveCassandraRepository;

import java.util.Date;

public interface MessageRepository extends ReactiveCassandraRepository<Message, Date> { 
    
}
