# Example

| file | description |
| ---- | ----------- |
| [**main.go**](main.go) | bootstraps HTTP server under **<http://localhost:8080>** with contextual request and error fields logging |
| [**middleware.go**](middleware.go) | defines contextual request logging HTTP middleware function |
| [**hello.go**](hello.go) | implements naive `/hello` endpoint which emits a contextual log |
| [**try.go**](try.go) | implements `/try` endpoint that most likely fails and logs error with a field |
| [**goroutine.go**](goroutine.go) | implements `/go` endpoint which shows how contextual logging can be used to trace forking multiple goroutines |

## Step by step guide

1. Clone the repository.
1. Enter this directory in terminal.
1. `go run .` to run the application.
1. Open **<http://localhost:8080/hello>** in a browser.
1. Look at the application logs. Notice that all logs emitted from a given request have the same `request_id` field value.
1. Read [**hello.go**](hello.go) to see how to make contextual logs.
1. Open **<http://localhost:8080/try>** in a browser **at least 2 times**.
1. Look at the application logs. Notice that the error log has an additional `point` field.
1. Read [**try.go**](try.go) to see how to wrap error with fields.
1. Read [**main.go**](main.go) and [**middleware.go**](middleware.go) to see how to setup contextual logging and error fields logging.
1. Open **<http://localhost:8080/go>** in a browser.
1. Look at the application logs. Notice the `jobID` and `parentJobID` log fields.
1. Read [**goroutine.go**](goroutine.go) to see how contextual logging can be used to trace goroutines.

**Tip:** Consider debugging the application for deeper understanding.
