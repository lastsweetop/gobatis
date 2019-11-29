package sqlparser

import "log"

type SqlSynx struct {
	Action string
	Fields []string `json:"field"`
	Params []string `json:"params"`
}

func Parser(sql string) *SqlSynx {
	sh := &SqlHelper{
		sql:    sql,
		offset: 0,
	}
	action := sh.GetAction()
	switch action {
	case "select":
		fields := sh.GetFileds()
		params := sh.GetParams()
		log.Println(params)
		return &SqlSynx{
			Action: action,
			Fields: fields,
			Params: params,
		}
		break
	case "delete":
		params := sh.GetParams()
		log.Println(params)
		return &SqlSynx{
			Action: action,
			Fields: nil,
			Params: params,
		}
		break
	case "insert":
		fields := sh.GetInsertFileds()
		return &SqlSynx{
			Action: action,
			Fields: fields,
			Params: nil,
		}
		break
	}
	return nil
}
