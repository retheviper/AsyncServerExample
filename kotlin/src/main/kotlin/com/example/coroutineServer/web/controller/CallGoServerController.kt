package com.example.coroutineServer.web.controller

import com.example.coroutineServer.domain.service.CallGoServerService
import com.example.coroutineServer.web.model.response.CallGoServerResponse
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class CallGoServerController(
    private val callGoServerService: CallGoServerService
) {

    @GetMapping("/call-go-server")
    fun callGoServer(): List<CallGoServerResponse> {
        return callGoServerService.callGoServer().map {
            CallGoServerResponse(it.id, it.result)
        }
    }

    @GetMapping("/call-go-server-async")
    suspend fun callGoServerAsync(): List<CallGoServerResponse> {
        return callGoServerService.callGoServerAsync().map {
            CallGoServerResponse(it.id, it.result)
        }
    }

    @GetMapping("/call-go-server-async-dual")
    suspend fun callGoServerAsyncDual(): List<CallGoServerResponse> {
        return callGoServerService.callGoServerAsyncDual().map {
            CallGoServerResponse(it.id, it.result)
        }
    }

}