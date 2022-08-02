package repository

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/denisenkom/go-mssqldb" //Driver SQL Server
)

const LayoutTimeStampMSSQL = "2006-01-02 15:04:05"

func InitDBMSSQL(host string, port int, user string, pass string, dbName string) (*sql.DB, error) {
	if port > 0 {
		host = fmt.Sprintf("%s:%d", host, port)
	}
	var instance string
	if strings.Contains(host, "\\") {
		hostBuf := strings.Split(host, "\\")
		host = hostBuf[0]
		instance = hostBuf[1]
	}

	query := url.Values{}
	query.Add("app name", "Locadora")
	query.Add("database", dbName)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(user, pass),
		Host:     host,
		Path:     instance,
		RawQuery: query.Encode(),
	}

	db, errdb := sql.Open("sqlserver", u.String())
	if errdb != nil {
		return nil, errdb
	}
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
