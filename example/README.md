# Example

| file | description |
| ---- | ----------- |
| [**main.go**](main.go) | bootstraps HTTP server under **<http://localhost:8080>** with contextual request and error fields logging |
| [**middleware.go**](middleware.go) | defines contextual request logging HTTP middleware function |
| [**hello.go**](hello.go) | implements naive `/hello` endpoint which emits a contextual log |
| [**try.go**](try.go) | implements `/try` endpoint that most likely fails and logs error with a field |
| [**goroutine.go**](goroutine.go) | implements `/go` endpoint which shows how contextual logging can be used to trace forking multiple goroutines |

## Step by step guide

1. Open terminal.
1. Clone this repository.
   ```
   git clone https://github.com/go-logrusutil/logrusutil.git
   ```
1. Enter this directory in terminal.
   ```
   cd logrusutil/example
   ```
1. Run the application.
   ```
   go run .
   ```
   Output:
   ```
   time="2019-11-11T22:33:59.1811608+01:00" level=info msg="server starting" app=example
   ```
1. Open **<http://localhost:8080/hello>** in a browser.
   Output:
   ```
   time="2019-11-11T22:34:05.1865488+01:00" level=info msg="request started" app=example reqID=5577006791947779410
   time="2019-11-11T22:34:05.187543+01:00" level=info msg="hello world" app=example foo=bar reqID=5577006791947779410
   time="2019-11-11T22:34:05.187543+01:00" level=info msg="request finished" app=example reqID=5577006791947779410
   ```
   Notice that all logs emitted from a given request have the same `request_id` field value.
1. Read [**hello.go**](hello.go) to see how to make contextual logs.
1. Open **<http://localhost:8080/try>** in a browser **2 times**.
   Output:
   ```
   time="2019-11-11T22:34:08.6542691+01:00" level=info msg="request started" app=example reqID=8674665223082153551
   time="2019-11-11T22:34:08.655269+01:00" level=info msg="try succeded" app=example reqID=8674665223082153551
   time="2019-11-11T22:34:08.6562676+01:00" level=info msg="request finished" app=example reqID=8674665223082153551
   time="2019-11-11T22:34:10.4275334+01:00" level=info msg="request started" app=example reqID=3916589616287113937
   time="2019-11-11T22:34:10.4285334+01:00" level=error msg="try failed" app=example error="failed to generate an excelent point" point="{2 1}" reqID=3916589616287113937
   time="2019-11-11T22:34:10.42955+01:00" level=info msg="request finished" app=example reqID=3916589616287113937
   ```
   Notice that the error log has an additional `point` field.
1. Read [**try.go**](try.go) to see how to wrap error with fields.
1. Read [**main.go**](main.go) and [**middleware.go**](middleware.go) to see how to setup contextual logging and error fields logging.
1. Open **<http://localhost:8080/go>** in a browser.
   Output:
   ```
   time="2019-11-11T22:34:12.6953435+01:00" level=info msg="request started" app=example reqID=1443635317331776148
   time="2019-11-11T22:34:12.6963432+01:00" level=info msg="new jobID" app=example jobID=1 reqID=1443635317331776148
   time="2019-11-11T22:34:12.6973451+01:00" level=info msg="request finished" app=example reqID=1443635317331776148
   time="2019-11-11T22:34:12.6973451+01:00" level=info msg="new jobID" app=example jobID=2 parentJobID=1 reqID=1443635317331776148
   time="2019-11-11T22:34:12.6973451+01:00" level=info msg="new jobID" app=example jobID=3 parentJobID=1 reqID=1443635317331776148
   time="2019-11-11T22:34:13.7329305+01:00" level=info msg="easy job done" app=example jobID=2 reqID=1443635317331776148
   time="2019-11-11T22:34:22.7370428+01:00" level=info msg="hard job done" app=example jobID=3 reqID=1443635317331776148
   ```
   Notice the `jobID` and `parentJobID` log fields.
1. Read [**goroutine.go**](goroutine.go) to see how contextual logging can be used to trace goroutines.

**Tip:** Consider debugging the application for deeper understanding.
