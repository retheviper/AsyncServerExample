package com.example.coroutineServer.domain.service

import com.example.coroutineServer.domain.model.ProcessDto
import org.springframework.stereotype.Service

@Service
class SomeProcessService {

    fun process(id: Int): ProcessDto {
        Thread.sleep(5000)
        return ProcessDto(
            id = id,
            result = if (id % 2 != 0) "success" else "something went wrong"
        )
    }
}