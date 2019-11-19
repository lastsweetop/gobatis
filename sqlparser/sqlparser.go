package sqlparser

type SqlSynx struct {
	Action string
	Fields []string `json:"field"`
}

func Parser(sql string) *SqlSynx {
	sh:=&SqlHelper{
		sql:    sql,
		offset: 0,
	}
	action:=sh.GetAction()
	return &SqlSynx{
		Action: action,
		Fields:  sh.GetFileds(),
	}
}
