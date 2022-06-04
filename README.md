## TL; DR

Async API call example by [goroutine](https://go.dev/tour/concurrency) and [coroutine](https://kotlinlang.org/docs/coroutines-overview.html).

Tested with Go 16.15 & Java 11.

Go server uses [Gin](https://gin-gonic.com) and Kotlin server uses [Spring Boot](https://spring.io/projects/spring-boot).

## Run

### Go

#### Run server

```shell
$ go build app.go
```

#### Async call to Kotlin Server

```
# with no cap
$ curl http://localhost:8800/api/v1/call-kotlin-server-async

# with cap(2)
$ curl http://localhost:8800/api/v1/call-kotlin-server-async-dual

```

### Kotlin Server

```shell
$ ./gradlew bootRun
```

#### Async call to Kotlin Server

```
# with no cap
$ curl http://localhost:8900/api/v1/call-go-server-async

# with cap(2)
$ curl http://localhost:8900/api/v1/call-go-server-async-dual

```
