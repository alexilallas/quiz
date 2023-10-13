## Quiz App

This is a simple quiz app using [gRPC](https://pkg.go.dev/google.golang.org/grpc) to server and [cobra](https://github.com/spf13/cobra) as front-end cli.

### Usage
- start gRPC server for quiz
  ```
  go run main.go server
  ```
- start quiz client and answer questions!
  ```
  go run main.go quiz
  ```

  