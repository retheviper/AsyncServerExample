package com.example.coroutineServer

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class CoroutineServerApplication

fun main(args: Array<String>) {
	runApplication<CoroutineServerApplication>(*args)
}
