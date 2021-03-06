## TL; DR

Async API call example by [goroutine](https://go.dev/tour/concurrency) and [coroutine](https://kotlinlang.org/docs/coroutines-overview.html).

Tested with Go 1.16.15 & Java 11.

Go server uses [Gin](https://gin-gonic.com) and Kotlin server uses [Spring Boot](https://spring.io/projects/spring-boot).

## Concept

Some API takes 5 second to process, and each server need to call the API for 10 times.

So, call it by async!

## Test

### Go

#### Run Server

```bash
$ go run app.go
```

#### Non-async call to Kotlin Server

```bash
$ curl http://localhost:8800/api/v1/call-kotlin-server
```

#### Async call to Kotlin Server

```bash
# with no cap
$ curl http://localhost:8800/api/v1/call-kotlin-server-async

# with cap(2)
$ curl http://localhost:8800/api/v1/call-kotlin-server-async-dual
```

### Kotlin Server

#### Run Server

```bash
$ ./gradlew bootRun
```

#### Non-async call to Go Server

```bash
$ curl http://localhost:8900/api/v1/call-go-server
```

#### Async call to Go Server

```bash
# with no cap
$ curl http://localhost:8900/api/v1/call-go-server-async

# with cap(2)
$ curl http://localhost:8900/api/v1/call-go-server-async-dual
```
