package sqlparser

import (
	"gobatis/utils"
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
		fields = append(fields, utils.UpperFirstWord(strings.TrimSpace(s)))
	}
	return fields
}

func (this *SqlHelper) GetInsertFileds() []string {
	fields := make([]string, 0)
	temp1 := strings.Index(this.sql[this.offset:], "(")
	temp2 := strings.Index(this.sql[this.offset:], ")")
	fieldstr := this.sql[this.offset+temp1+1 : this.offset+temp2]
	fieldstrs := strings.Split(fieldstr, ",")
	for _, s := range fieldstrs {
		fields = append(fields, utils.UpperFirstWord(strings.TrimSpace(s)))
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
		params = append(params, utils.UpperFirstWord(strings.TrimSpace(param[0])))
	}
	return params
}
