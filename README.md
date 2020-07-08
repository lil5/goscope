# GoScope

Watch incoming requests and outgoing responses from your Go Gin application. All is logged into a database for persistence and paginated for performance.

The aim of this application is to be a plug and play addition to your application, not a hurdle.
Thus to setup, you only require a one-liner in your main function.
Once all is set up you can access the web interface by visiting `http://your-app.com/goscope`. It is recommended that you will protect this route from external/public access so that you do not leak important application data.

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

```
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `requests`;
CREATE TABLE `requests` (
  `uid` varchar(32) NOT NULL,
  `application` varchar(255) NOT NULL,
  `client_ip` varchar(255) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  `headers` text,
  `body` longtext,
  `referrer` varchar(255) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `responses`;
CREATE TABLE `responses` (
  `uid` varchar(32) NOT NULL,
  `request_uid` varchar(32) NOT NULL,
  `application` varchar(255) NOT NULL,
  `client_ip` varchar(255) DEFAULT NULL,
  `status` int(255) DEFAULT NULL,
  `time` bigint(20) DEFAULT NULL,
  `body` longtext,
  `path` varchar(255) DEFAULT NULL,
  `headers` text,
  `size` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE,
  KEY `request_uid_foreign` (`request_uid`),
  CONSTRAINT `request_uid_foreign` FOREIGN KEY (`request_uid`) REFERENCES `requests` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
```