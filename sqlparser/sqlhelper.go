package sqlparser

import (
	"log"
	"strings"
)

type SqlHelper struct {
	sql    string
	offset int
}

func (this *SqlHelper) GetAction() string {
	this.offset = strings.Index(this.sql, " ")
	log.Println(this.offset)
	return this.sql[:this.offset]
}

func (this *SqlHelper) GetFileds() []string {
	fields := make([]string, 0)

	temp := strings.Index(this.sql[this.offset:], "from")
	fieldstr := this.sql[this.offset : this.offset+temp]
	fieldstrs := strings.Split(fieldstr, ",")
	for _, s := range fieldstrs {
		fields = append(fields, upperFirstWorld(strings.TrimSpace(s)))
	}
	return fields
}

func upperFirstWorld(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}
