package test

import (
	"gobatis/sqlparser"
	"log"
	"testing"
)

func Test(t *testing.T) {
	sql:=sqlparser.Parser("select id,account from user")
	log.Println(sql)
}
