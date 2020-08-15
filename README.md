# GoScope

[![Build Status](https://travis-ci.com/averageflow/goscope.svg?branch=master)](https://travis-ci.com/averageflow/goscope)

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
# for MySQL
GOSCOPE_DATABASE_CONNECTION: "root:root@tcp(127.0.0.1:3306)/go_scope"  
GOSCOPE_DATABASE_TYPE: "mysql"
# for SQLite
GOSCOPE_DATABASE_CONNECTION: "file:/Users/joe/workspace/goscope/setup/sqlite.sqlite" 
GOSCOPE_DATABASE_TYPE: "sqlite3"

GOSCOPE_ENTRIES_PER_PAGE: 50
```

GoScope has been extended to work with a repository pattern, thus has the capability of supporting any database driver/engine that will work with Go and uses the `sql` package (returning `*sql.Rows` or `*sql.Row`). 
NoSQL databases are currently not supported, although we think that it would be a great addition, so if you have the know-how please don't hesitate to make a Pull Request.
In the .env file you can specify either the `mysql` driver, `postgres` driver, or `sqlite3` driver, which will use the `github.com/go-sql-driver/mysql`, `github.com/lib/pq` or `github.com/mattn/go-sqlite3` respectively. Ensure you have the correct connection string in your env file.

The application expects a database with a setup that can be recreated by taking a look at the `setup` folder in the root of this repository.

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

GoScope hooks into the logging mechanism of your application by using a custom `io.Writer`. 
This prints to the console for ease of development and saves the logs into the database, for further displaying in the web environment.
Thus you only need to call your usual `log.Println` or `log.Printf` statements or any variants of the log writing package, and that will seamlessly be picked up by GoScope.

![GoScope Dashboard](https://raw.githubusercontent.com/averageflow/goscope/master/showcase/2.png)

### Contributing

Notice that in the Angular app, the environment files can be edited to suit your needs. Of course you will need an instance of GoScope working, with some example logs in the DB in order to see anything in the SPA.
Change that file to match your port, in my case `http://localhost:7004/`.
Any change to assets in the `static` folder, obviously building the SPA too, will require a rebuild of the `bindata.go`. 
For this you will require to have the package installed (via `go get -u github.com/shuLhan/go-bindata/...`).

Then navigate in terminal to the root of the project and run (following commands all assume your pwd is the root of project): 
- Build the Angular app run `cd static/goscope && ng build --prod --output-hashing none --base-href /goscope/`
- To create the bindata file: `cd ../../goscope && go-bindata -nomemcopy  ../static/goscope/dist/...`
- Edit the `bindata.go` file in `/goscope` so that the package is `goscope` instead of `main`


### API Documentation

GoScope is in its essence a middleware, a backend in Go and a front-end in Angular, thus you can hook onto the backend with any tool of your choice, being it other kinds of front-ends or monitoring tools like Zabbix. 
For that purpose please find here the [API spec](SPEC-API.md).
  


