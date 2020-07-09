# GoScope

Watch incoming requests and outgoing responses from your Go Gin application. All is logged into a database for persistence and paginated for performance.

The aim of this application is to be a plug and play addition to your application, not a hurdle.

Thus to setup, you only require a one-liner in your main function.

Once all is set up you can access the web interface by visiting `http://your-app.com/goscope`. 

It is recommended that you will protect this route from external/public access so that you do not leak important application data.

![GoScope Dashboard](https://pro-warehouse-res.cloudinary.com/image/upload/v1594213052/git-repositories/goscope/zgl0npjbhwcs89snwm8u.png)

Example implementation code: 
```go
package main
import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "bitbucket.org/prowarehouse-nl/goscope/goscope"
)

func main(){
    router := gin.Default()
    _ = godotenv.Load()
    goscope.Setup(router)
    
    router.GET("/ping")
}
```

The application requires that your `.env` file contains the following variables:

```yaml
APPLICATION_NAME: "Your Application Name"
APPLICATION_ID: "application-id"
WATCHER_DATABASE_CONNECTION: "root:root@tcp(127.0.0.1:3306)/go_scope"
```

The application expects a MySQL database (preferably MariaDB) with a setup that can be recreated by taking a look at the `setup.sql` file in the root of this repository.

More detailed information about request and response is provided:

![GoScope Dashboard](https://pro-warehouse-res.cloudinary.com/image/upload/v1594213049/git-repositories/goscope/ar5rsnlku4xm8nvf5wj0.png)

![GoScope Dashboard](https://pro-warehouse-res.cloudinary.com/image/upload/v1594213049/git-repositories/goscope/hbeyhmcfii6zmgixyn90.png)