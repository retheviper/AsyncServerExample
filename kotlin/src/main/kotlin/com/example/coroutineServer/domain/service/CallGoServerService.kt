package com.example.coroutineServer.domain.service

import com.example.coroutineServer.domain.model.CallGoServerDto
import com.example.coroutineServer.infrastructure.client.GoServerClient
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.sync.Semaphore
import kotlinx.coroutines.sync.withPermit
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service

@Service
class CallGoServerService(
    private val goServerClient: GoServerClient
) {
    private val logger = LoggerFactory.getLogger(CallGoServerService::class.java)
    private val tries = (1..10).toList()

    fun callGoServer(): List<CallGoServerDto> {
        logger.info("[CallGoServer] start")
        return tries.map {
            logger.info("[CallGoServer] before request with id: $it")
            runBlocking {
                goServerClient.call(it) ?: CallGoServerDto(it, "failed")
            }
        }
    }

    suspend fun callGoServerAsync(): List<CallGoServerDto> {
        logger.info("[callGoServerAsync] start")
        return coroutineScope {
            tries.map {
                async {
                    logger.info("[callGoServerAsync] before request with id: $it")
                    goServerClient.call(it) ?: CallGoServerDto(it, "failed")
                }
            }.awaitAll()
                .also {
                    it.forEach { result ->
                        logger.info("[callGoServerAsyncDual] after request with id: ${result.id}")
                    }
                    logger.info("[callGoServerAsyncDual] done")
                }
        }
    }

    suspend fun callGoServerAsyncDual(): List<CallGoServerDto> {
        logger.info("[callGoServerAsyncDual] start")
        val semaphore = Semaphore(2)
        return coroutineScope {
            tries.map {
                async {
                    semaphore.withPermit {
                        logger.info("[callGoServerAsyncDual] before request with id: $it")
                        goServerClient.call(it) ?: CallGoServerDto(it, "failed")
                    }
                }
            }
        }.awaitAll()
            .also {
                it.forEach { result ->
                    logger.info("[callGoServerAsyncDual] after request with id: ${result.id}")
                }
                logger.info("[callGoServerAsyncDual] done")
            }
    }
}