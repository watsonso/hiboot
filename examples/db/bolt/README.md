# Web Application with Bolt Database

This is the web application with Bolt database.

As you can see below, the usage is extremely simple, hiboot support dependency injection, all you have to do is config data source in application.yml, add tags in struct, then you have database repository injected in your service.

## Application entry point 

main.go

```go
package main

import (
	"github.com/hidevopsio/hiboot/pkg/starter/web"
	_ "github.com/hidevopsio/hiboot/examples/web/jwt/controllers"
)

func main()  {
	// create new web application and run it
	web.NewApplication().Run()
}
```

## Config data source

```yaml
dataSources:
- type: bolt
  database: hi.db
  mode: 0600
  timeout: 1

```

## Inject Service into Controller

In order to inject Repository into Service, you need to 

* add tag `component:"service"` to field UserService of UserController

```go
type UserController struct {
	web.Controller

	UserService *services.UserService `component:"service"`
}

```

## Inject Repository into Service

In order to inject Repository into Service, you need to 

* import github.com/hidevopsio/hiboot/pkg/starter/db
* add tag `component:"repository" dataSourceType:"bolt"` to the field Repository of UserService


```go

import (
	"github.com/hidevopsio/hiboot/pkg/starter/db"
)

type UserService struct {
	Repository db.KVRepository `component:"repository" dataSourceType:"bolt"`
}

```

## Run unit test
```bash
go test ./...
```

## Run the example code
```bash
go run main.go
```

## Run test

### Post API

Post User in JSON

```bash

curl -H -X POST -d '{"id": "1", "name": "John Doe", "age": 25}' http://localhost:8080/user

```

The output will be 

```json
{
    "code": 200, 
    "data": {
        "Age": 25, 
        "Id": "1", 
        "Name": "John Doe"
    }, 
    "message": "Success"
}
```

### Get API

Get User with Id

```bash
curl http://localhost:8080/user?id=1
```

the output will be

```json
{
    "code": 200, 
    "data": {
        "Age": 25, 
        "Id": "1", 
        "Name": "John Doe"
    }, 
    "message": "Success"
}
```

### Delete API

Delete User

```bash
curl -X DELETE http://localhost:8080/user?id=1  
```

The output will be

```json
{
    "code": 200, 
    "data": null, 
    "message": "Success"
}

```