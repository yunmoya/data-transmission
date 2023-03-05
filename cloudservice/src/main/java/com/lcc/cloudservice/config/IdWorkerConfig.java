package com.lcc.cloudservice.config;

import com.lcc.cloudservice.uid.IdWorker;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class IdWorkerConfig {
    @Value("${csp.id}")
    private Long workerId;

    @Bean
    public IdWorker idWorker() {
        return new IdWorker(workerId, 1L, 1L);
    }
}
