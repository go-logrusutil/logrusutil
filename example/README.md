# Example

| file | description |
| ---- | ----------- |
| [**main.go**](main.go) | bootstraps HTTP server under **<http://localhost:8080>** with contextual request logging and error fields |
| [**middleware.go**](middleware.go) | defines contextual request logging HTTP middleware function |
| [**hello.go**](hello.go) | implements naive `/hello` endpoint which emits a contextual log |
| [**try.go**](try.go) | implements `/try` endpoint that most likely fails and logs error with a field |
| [**goroutine.go**](goroutine.go) | implements `/go` endpoint which shows how contextual logging can be used to trace forking multiple goroutines |

## Step by step guide

1. Clone the repository
1. Enter this directory in terminal
1. `go run .` to run the application
1. Open **<http://localhost:8080/hello>** in a browser
1. Look at the application logs. Notice that all logs emitted from a given request have the same `request_id` field value.
1. Read [**hello.go**](hello.go) to see basic usage of contextual logging
1. Open **<http://localhost:8080/try>** in a browser
1. Look at the application logs. Notice that error log has an additional `point` field.
1. Read [**hello.go**](hello.go) to see basic usage of error wrapping and logging with fields
1. Read [**main.go**](main.go) and [**middleware.go**](middleware.go) to see how to setup contextual logging and error fields logging
1. Open **<http://localhost:8080/go>** in a browser
1. Look at the application logs. Notice the `ctxID` and `prevCtxID` log fields.
1. Read [**goroutine.go**](goroutine.go) to see  contextual logging can be used to trace goroutines

Consider debugging the application to have deep understanding.
