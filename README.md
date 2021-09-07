# Roman Numeral REST API

Simple REST API to convert/get ordinary number to roman numeral.
This REST API doesn't use any frameworks,
only use builtin package `net/http`.
And no database implementation.

## Features

- Coverage unit testing within utils
- Coverage endpoints testing using Postman

## Installation

Run it temporarily using `go run main.go`, or build the binary and run.
```
go build
./roman-numeral-api
```
Or install to local machine `go install github.com/study-hary-id/roman-numeral-api`.

## Endpoints

### Get Roman Numeral

```
GET /roman-numbers/{num:int}
```
### Successful Response

```json
{
  "data": {
    "roman": string,
    "ordinal": int
  }
}
```
## Credits

The idea is from this book [Building RESTful Web services with Go_ Learn how to build RESTful APIs with Golang that scale gracefully](https://www.packtpub.com/product/building-restful-web-services-with-go/9781788294287), with some modifications.

#### Copyright © 2021 [@study-hary-id](https://github.com/study-hary-id). This project is under [CC0](LICENSE) licensed
