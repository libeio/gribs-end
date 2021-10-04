
package table

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"grids.com/pkgs"

	"log"
)

const (
    _seq_tbl_permission = "CREATE SEQUENCE IF NOT EXISTS tbl_permission_id_seq START WITH 1;"
    _seq_tbl_role = "CREATE SEQUENCE IF NOT EXISTS tbl_role_id_seq START WITH 1;"
)

type TblPermission struct {
	tableName struct{}			`pg:"public.tbl_permission"`
    ID    int32                 `pg:"type:integer,pk,notnull,default:nextval('tbl_permission_id_seq')"`
    Name  string                `pg:"type:varchar(100),unique,notnull"`
}

type TblRole struct {
    tableName   struct{}        `pg:"public.tbl_role"`
    ID    int32                 `pg:"type:integer,pk,notnull,default:nextval('tbl_role_id_seq')"`
    Name  string                `pg:"type:varchar(32),unique,notnull"`
}

var _permission_models = []interface{} {
	(*TblPermission)(nil),
	(*TblRole)(nil),
}

var _permissions  = []TblPermission {
	{ Name: "can_read" },
	{ Name: "can_edit" },
	{ Name: "can_delete" },
	{ Name: "can_create" },
}

var _roles = []TblRole {
	{ Name: "Admin" },
	{ Name: "Public" },
	{ Name: "Viewer" },
	{ Name: "User" },
	{ Name: "Op" },
}

func InitPermission() {
	db := pkgs.GetDB()
	if db == nil {
		log.Panicln("db is nil")
	}

	db.Exec("BEGIN TRANSACTION;")
	defer func() {
		db.Exec("COMMIT;")
	} ()

	// create sequence and table
	for _, model := range _permission_models {
		switch model.(type) {
		case *TblPermission:
			_, err := db.Exec(_seq_tbl_permission)
			if err != nil {
				log.Panicln(err)
			}
		case *TblRole:
			_, err := db.Exec(_seq_tbl_role)
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

	// insert
	insertPermissions(db)
	insertRoles(db)

	log.Printf("init permission success")
}

func insertPermissions(db *pg.DB) {
	var ps []interface{}
	for _, p := range _permissions {
		x := p
		ps = append(ps, &x)
	}
	for _, e := range ps {
		if _, err := db.Model(e).OnConflict("DO NOTHING").Insert(); err != nil {
			log.Println(err)
		}
	}
}

func insertRoles(db *pg.DB) {
	var rs []interface{}
	for _, r := range _roles {
		x := r
		rs = append(rs, &x)
	}
	for _, e := range rs {
		if _, err := db.Model(e).OnConflict("DO NOTHING").Insert(); err != nil {
			log.Println(err)
		}
	}
}
