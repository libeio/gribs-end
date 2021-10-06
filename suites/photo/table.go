
package main

import (
	"github.com/go-pg/pg/v10/orm"

	"grids.com/grids"

	"time"
	"log"
)

const (
	_seq_tbl_photo_show = "CREATE SEQUENCE IF NOT EXISTS tbl_photo_show_id_seq START WITH 1;"
)

type TblPhotoShow struct {
    tableName   struct{}        `pg:"public.tbl_photo_show"`
	ID          int32           `pg:"type:integer,pk,notnull,default:nextval('tbl_photo_show_id_seq')"`
	ResourceId	string			`pg:"type:varchar(16),notnull,default:'photo'"`		// 标识资源
	Role        string          `pg:"type:varchar(32),notnull,default:'Public'"`	// 该资源适用的角色用户
	Groups		[]string		`pg:"type:varchar(32)[],array"`						// 组，除适用角色用户外，也适用于指定的组
    Additional	[]string		`pg:"type:varchar(32)[],array"`                     // 额外，除适用于角色和组用户外，也适用于某些个体
    Excluded    []string        `pg:"type:varchar(32)[],array"`                     // 指定用户不适用于该资源
    
    Sorted      int32           `pg:"type:integer,notnull"`							// 页面排序用
	Tag         []string        `pg:"type:varchar(16)[],array"`						// 页面过滤搜索用
	// 后续将这里加一张切片作为表，应该也不需要，加个表名就行了
    Href        string          `pg:"type:text,notnull"`
    Src         string          `pg:"type:text,notnull"`
	Description string          `pg:"type:text"`
	// 添加一张真实资源表

    CreateTime  time.Time       `pg:"type:timestamp,notnull,default:now()::timestamp"`
    UpdateTime  time.Time       `pg:"type:timestamp,notnull,default:now()::timestamp"`
}

var _models = []interface{} {
	(*TblPhotoShow)(nil),
}

func InitDB() {
	db := grids.GetDB()
	if db == nil {
		log.Panicln("db is nil")
	}

	db.Exec("BEGIN TRANSACTION;")
	defer func() {
		db.Exec("COMMIT;")
	} ()

	// create sequence and table
	for _, model := range _models {
		switch model.(type) {
		case *TblPhotoShow:
			_, err := db.Exec(_seq_tbl_photo_show)
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

	log.Printf("init db success")
}
