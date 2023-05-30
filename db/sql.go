package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	InsItem  = "INSERT INTO hadinfo_item(item_id,type,value) VALUES ($1,$2,$3)"
	InsPaste = "INSERT INTO hadinfo_paste(paste_id,title,author,meta,item_id) VALUES($1,$2,$3,$4,$5)"

	QueryItem  = "SELECT * FROM hadinfo_item WHERE item_id=$1"
	QueryPaste = "SELECT * FROM hadinfo_paste WHERE paste_id=$1"
)

func SqlQueryExist(sql string, arg ...any) bool {
	r, err := db.Query(strings.Replace(sql, "SELECT *", "SELECT count(*)", 1), arg...)
	if err != nil {
		logrus.Fatal("PGSQL Statements Error(Exist): ", err)
	}
	var count int
	r.Next()
	err = r.Scan(&count)
	if err != nil {
		count = 0
	}
	_ = r.Close()
	return count != 0
}

func SqlQuery(sql string, arg ...any) *sql.Rows {
	r, err := db.Query(sql, arg...)
	if err != nil {
		logrus.Fatal("PGSQL Statements Error(Query): ", err)
	}
	return r
}

func SqlExec(sql string, arg ...any) sql.Result {
	r, err := db.Exec(sql, arg...)
	if err != nil {
		logrus.Fatal("PGSQL Statements Error(Exec): ", err, ' ', sql, ' ', arg)
	}
	return r
}
