# <img src="images/gopher-knight.svg" alt="Gopher knight" width="24"> Roman Numeral REST API

Simple REST API to convert/`GET` ordinary number to roman numeral. This REST API doesn't use any frameworks, only use
builtin package `net/http`. And no database implementation.

## <img src="images/pushing-cart.png" alt="Gopher pushing cart" width="48"> Features

- Coverage unit testing within utils directory
- Convert ordinal numbers to roman up to 3999

## <img src="images/gopher-gamer.svg" alt="Gopher gamer" width="24"> Installation

Run it temporarily using `go run main.go`, or build the binary and run.

```
go build
./roman-numeral-api
```

Or install to local machine `go install github.com/study-hary-id/roman-numeral-api`.

## <img src="images/messenger-running.png" alt="Messenger running" width="24"> Endpoints

| HTTP Verb | Path                | Action | Resource      |
| --------- | ------------------- | ------ | ------------- |
| GET       | `/roman-numbers/17` | Show   | roman-numbers |

### Get Roman Numeral

```
GET /roman-numbers/{num:int}
```

### Successful Response

```json
{
  "data": {
    "roman": "XVII",
    "ordinal": 17
  }
}
```

### Error Response

```json
{
  "errors": {
    "status": 404,
    "title": "Not Found",
    "detail": "'/' not found."
  }
}
```

## <img src="images/crash-dummy.svg" alt="Gopher robo crash" width="24"> Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## <img src="images/heart-balloon.svg" alt="Gopher heart balloon" width="24"> Credits

The idea is from this
book [Building RESTful Web services with Go_](https://www.packtpub.com/product/building-restful-web-services-with-go/9781788294287)
, with some modifications.

The images, svg icons and gif animations are from [Gophers](https://github.com/egonelbre/gophers)

Copyright :copyright: 2022. This project is under [CC0](LICENSE) licensed.
