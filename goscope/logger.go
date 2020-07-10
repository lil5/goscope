package goscope

import (
	"database/sql"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"html"
	"log"
	"os"
	"time"
)

type LoggerGoScope struct {
}

func (logger LoggerGoScope) Write(p []byte) (n int, err error) {
	go Log(string(p))
	return len(p), nil
}

func Log(message string) {
	fmt.Println(message)
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	uid, _ := uuid.NewV4()
	query := "INSERT INTO `logs` (`uid`, `application`, `error`, `time`) VALUES " +
		"('%s', '%s', '%s', %v)"
	resultingQuery := fmt.Sprintf(query, uid, os.Getenv("APPLICATION_ID"), html.EscapeString(message), time.Now().Unix())
	_, err = db.Exec(resultingQuery)
	if err != nil {
		log.Println(err.Error())
	}
}
