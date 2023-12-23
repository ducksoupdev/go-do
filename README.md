# Digital Ocean Golang Helpers

This is a collection of helpers for working with Digital Ocean Apps/Functions using Golang

## Installation

This library is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/ducksoupdev/go-do/functions
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/ducksoupdev/go-do/functions"
```

and run `go get` without parameters.

## Usage

There are structs, functions and methods for working with Digital Ocean Apps/Functions.

### Headers struct

The headers struct is used to set the headers for a JSON response. It has a predefined `ContentType: applicatiom/json` value.

```go
func Main(in Request) (*functions.Response, error) {
  // do something

  return &functions.Response{
    StatusCode: 200,
    Headers:    functions.Headers,
    Body:       string(jsonBody),
  }, nil
}
```

### Response struct

The response struct is used to return a response from a function. It has `StatusCode`, `Headers` and `Body` fields.

```go
func Main(in Request) (*functions.Response, error) {
  // do something

  return &functions.Response{
    StatusCode: 200,
    Headers:    functions.Headers,
    Body:       string(jsonBody),
  }, nil
}
```

### Error response struct

The error response struct is used to return an error response from a function. It has `Title`, `StatusCode` and `Detail` fields.

```go
func Main(in Request) (*functions.Response, error) {
  // do something

  errResponse, _ := json.Marshal(functions.ErrorResponse{
    Title:      title,
    StatusCode: statusCode,
    Detail:     message,
  })
  return &functions.Response{
    StatusCode: statusCode,
    Headers:    functions.Headers,
    Body:       string(errResponse),
  }, nil
}
```

### Error handler function

The `ErrorHandler` function is used to return an error response from a function. It accepts `title`, `statusCode` and `message` arguments.

```go
func Main(in Request) (*functions.Response, error) {
  // do something

  return functions.ErrorHandler(functions.StatusCodeTitle(statusCode), statusCode, message)
}
```

### Status code title function

The `StatusCodeTitle` function is used to return a title for a status code. It accepts a single `statusCode` argument.

```go
func Main(in Request) (*functions.Response, error) {
  // do something

  return functions.ErrorHandler(functions.StatusCodeTitle(statusCode), statusCode, message)
}
```

The title is returned from the following map:

```go
var statusCodes = map[int]string{
  100: "Continue",
  101: "Switching Protocols",
  102: "Processing",
  103: "Early Hints",
  200: "OK",
  201: "Created",
  202: "Accepted",
  203: "Non-Authoritative Information",
  204: "No Content",
  205: "Reset Content",
  206: "Partial Content",
  207: "Multi-Status",
  208: "Already Reported",
  226: "IM Used",
  300: "Multiple Choices",
  301: "Moved Permanently",
  302: "Found",
  303: "See Other",
  304: "Not Modified",
  305: "Use Proxy",
  306: "Switch Proxy",
  307: "Temporary Redirect",
  308: "Permanent Redirect",
  400: "Bad Request",
  401: "Unauthorized",
  402: "Payment Required",
  403: "Forbidden",
  404: "Not Found",
  405: "Method Not Allowed",
  406: "Not Acceptable",
  407: "Proxy Authentication Required",
  408: "Request Timeout",
  409: "Conflict",
  410: "Gone",
  411: "Length Required",
  412: "Precondition Failed",
  413: "Payload Too Large",
  414: "URI Too Long",
  415: "Unsupported Media Type",
  416: "Range Not Satisfiable",
  417: "Expectation Failed",
  418: "I'm a teapot",
  421: "Misdirected Request",
  422: "Unprocessable Entity",
  423: "Locked",
  424: "Failed Dependency",
  425: "Too Early",
  426: "Upgrade Required",
  428: "Precondition Required",
  429: "Too Many Requests",
  431: "Request Header Fields Too Large",
  451: "Unavailable For Legal Reasons",
  500: "Internal Server Error",
  501: "Not Implemented",
  502: "Bad Gateway",
  503: "Service Unavailable",
  504: "Gateway Timeout",
  505: "HTTP Version Not Supported",
  506: "Variant Also Negotiates",
  507: "Insufficient Storage",
  508: "Loop Detected",
  510: "Not Extended",
  511: "Network Authentication Required",
}
```

## License

This project is licensed under the MIT License - see the [`LICENSE`](LICENSE) file for details.

## Contributing

Any kind of positive contribution is welcome! Please help us to grow by contributing to the project.

If you wish to contribute, you can work on any features you think would enhance the library. After adding your code, please send us a Pull Request.

> Please read [CONTRIBUTING](CONTRIBUTING.md) for details on our [CODE OF CONDUCT](CODE_OF_CONDUCT.md), and the process for submitting pull requests to us.

## Support

We all need support and motivation. Please give this project a ⭐️ to encourage and show that you liked it. Don't forget to leave a star ⭐️ before you move away.

