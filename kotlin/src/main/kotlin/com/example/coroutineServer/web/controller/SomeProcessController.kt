package com.example.coroutineServer.web.controller

import com.example.coroutineServer.domain.service.SomeProcessService
import com.example.coroutineServer.web.model.request.SomeProcessRequest
import com.example.coroutineServer.web.model.response.SomeProcessResponse
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

@RestController
class SomeProcessController(
    private val someProcessService: SomeProcessService
) {

    @PostMapping("/some-process")
    fun someProcess(@RequestBody request: SomeProcessRequest): SomeProcessResponse {
        val dto = someProcessService.process(request.id)
        return SomeProcessResponse(
            id = dto.id,
            result = dto.result
        )
    }
}