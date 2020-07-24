# GoScope

[![Build Status](https://travis-ci.com/averageflow/goscope.svg?branch=master)](https://travis-ci.com/averageflow/goscope)
![golangci-lint](https://github.com/averageflow/goscope/workflows/golangci-lint/badge.svg?branch=master)

Watch incoming requests and outgoing responses from your Go Gin application. All is logged into a database for persistence and paginated for performance.

The aim of this application is to be a plug and play addition to your application, not a hurdle, thus to setup, you only require a one-liner in your main function.

Once all is set up you can access the web interface by visiting `http://your-app.com/goscope`. 

It is recommended that you will protect this route from external/public access so that you do not leak important application data.

### Setup


The application requires that your `.env` file contains the following variables:

```yaml
APPLICATION_ID: "application-id"
APPLICATION_NAME: "Your Application Name"
APPLICATION_TIMEZONE: "Europe/Amsterdam"
GOSCOPE_DATABASE_CONNECTION: "root:root@tcp(127.0.0.1:3306)/go_scope"
GOSCOPE_DATABASE_TYPE: "mysql"
GOSCOPE_ENTRIES_PER_PAGE: 50
```

In the .env file you can specify either the `mysql` driver or `postgres` driver, which will use the `github.com/go-sql-driver/mysql` or `github.com/lib/pq`
The application expects a database with a setup that can be recreated by taking a look at the `mysql-setup.sql` file in the root of this repository.

### Example
Example implementation code, please note that you should use plain gin without middlewares, since GoScope will use Gin Gonic's logger and recovery middlewares, but with a customized twist, thus the requirement is that initially you have a clean `gin.Engine` instance.

```go
package main
import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/averageflow/goscope/goscope"
)

func main(){
    router := gin.New()
    _ = godotenv.Load()
    goscope.Setup(router)
    
    router.GET("/ping")
}
```

### Request & Response

In order to understand possible unexpected situations or simply to reassure you that everything is working correctly, GoScope provides more detailed information about request and response, including status codes and request/response bodies, as well as any useful information. Any help with expanding this would be greatly appreciated.

![GoScope Dashboard](https://raw.githubusercontent.com/averageflow/goscope/master/showcase/0.png)

![GoScope Dashboard](https://raw.githubusercontent.com/averageflow/goscope/master/showcase/1.png)

### Logging

GoScope hooks into the logging mechanism of your application by using a custom `io.Writer` that prints to the console for ease of development and saves the logs into the database, for further displaying in the web environment.

This means that you only need to call your usual `log.Println` or `log.Printf` statements or any variants of the log writing package, and that will seamlessly be picked up by GoScope.

![GoScope Dashboard](https://raw.githubusercontent.com/averageflow/goscope/master/showcase/2.png)

### Contributing

Any change to assets in the `static` folder will require a rebuild of the `bindata.go`. For this you will require to have the package installed (via `go get -u github.com/shuLhan/go-bindata/...
`) and then navigate in terminal to the root of the project and run `cd goscope && go-bindata -nomemcopy  ../static/...`. For live edits use the debug mode `cd goscope && go-bindata -debug ../static/...`, then edit the `bindata.go` file in `/goscope` so that the package is `goscope` instead of `main`.