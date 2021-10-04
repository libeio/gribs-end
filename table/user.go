
package table

import (
	"github.com/go-pg/pg/v10/orm"

	"grids.com/pkgs"

	"time"
	"log"
)

const (
    _seq_tbl_user = "CREATE SEQUENCE IF NOT EXISTS tbl_user_id_seq START WITH 1;"
)

type TblUser struct {
	tableName struct{}	        `pg:"public.tbl_user"`
	ID          int32           `pg:"type:integer,pk,notnull,default:nextval('tbl_user_id_seq')"`
	Name		string			`pg:"type:varchar(32),notnull"`		// 用户名称
	Email		string			`pg:"type:varchar(64),notnull"`		// 用户邮箱
	Phone		string			`pg:"type:varchar(32)"`				// 用户电话				
	Role		string			`pg:"type:varchar(32),notnull"`		// 用户角色
    Groups      []string        `pg:"type:varchar(64)[],array"`		// 加入的组，一般查询使用
    ReginId     string          `pg:"type:varchar(64)"`             // 区域标识
    CreateTime  time.Time       `pg:"type:timestamp,notnull,default:now()::timestamp"`
}

var _user_models = []interface{} {
	(*TblUser)(nil),
}

func InitUser() {
	db := pkgs.GetDB()
	if db == nil {
		log.Panicln("db is nil")
	}

	db.Exec("BEGIN TRANSACTION;")
	defer func() {
		db.Exec("COMMIT;")
	} ()

	// create sequence and table
	for _, model := range _user_models {
		switch model.(type) {
		case *TblUser:
			_, err := db.Exec(_seq_tbl_user)
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

	log.Printf("init user success")
}
