package com.example.coroutineServer.infrastructure.client

import com.example.coroutineServer.domain.model.CallGoServerDto
import com.example.coroutineServer.infrastructure.model.CallGoServerRequest
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import org.springframework.http.HttpEntity
import org.springframework.http.HttpHeaders
import org.springframework.http.MediaType
import org.springframework.stereotype.Component
import org.springframework.web.client.RestTemplate

@Component
class GoServerClient {

    private val client = RestTemplate()

    private val header = HttpHeaders().apply {
        set(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
    }

    suspend fun call(id: Int): CallGoServerDto? {
        val request = HttpEntity(CallGoServerRequest(id), header)
        return withContext(Dispatchers.IO) {
            client.postForObject("http://localhost:8800/api/v1/some-process", request, CallGoServerDto::class.java)
        }
    }
}