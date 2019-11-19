package sqlparser

import (
	"strings"
)

type SqlHelper struct {
	sql    string
	offset int
}

func (this *SqlHelper) GetAction() string {
	this.offset = strings.Index(this.sql, " ")
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

func (this *SqlHelper) GetParams() []string {
	params := make([]string, 0)
	temp := strings.Index(this.sql, "where ") + 6
	if temp == 5 {
		return params
	}
	str := this.sql[temp:]
	strs := strings.Split(str, "and")
	for _, s := range strs {
		param := strings.Split(strings.TrimSpace(s), "=")
		params = append(params, upperFirstWorld(strings.TrimSpace(param[0])))
	}
	return params
}

func upperFirstWorld(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}
