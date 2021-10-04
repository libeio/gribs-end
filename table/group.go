
package table

import (
	"github.com/go-pg/pg/v10/orm"

	"grids.com/pkgs"

	"time"
	"log"
)

const (
    _seq_tbl_group = "CREATE SEQUENCE IF NOT EXISTS tbl_group_id_seq START WITH 1;"
)

type TblGroup struct {
	tableName struct{}	        `pg:"public.tbl_group"`
    ID          int32           `pg:"type:integer,pk,notnull,default:nextval('tbl_group_id_seq')"`
    Name        string          `pg:"type:varchar(64),unique,notnull"`  // 组名
    Members     []string        `pg:"type:varchar(64)[],array"`         // 组中成员，关联查询使用
	ReginId     string          `pg:"type:varchar(64)"`                 // 区域标识
	PassWord	string			`pg:"type:varchar(64)"`					// 区域进入码，后续设置为 notnull
    CreateTime  time.Time       `pg:"type:timestamp,notnull,default:now()::timestamp"`
}

var _group_models = []interface{} {
	(*TblGroup)(nil),
}

func InitGroup() {
	db := pkgs.GetDB()
	if db == nil {
		log.Panicln("db is nil")
	}

	db.Exec("BEGIN TRANSACTION;")
	defer func() {
		db.Exec("COMMIT;")
	} ()

	// create sequence and table
	for _, model := range _group_models {
		switch model.(type) {
		case *TblGroup:
			_, err := db.Exec(_seq_tbl_group)
			if err != nil {
				log.Panicln(err)
			}
		default:
			log.Panicln("unsupport table")
		}

		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,
		})
		if err != nil {
			log.Panicln(err)
		}
	}

	log.Printf("init group success")
}
