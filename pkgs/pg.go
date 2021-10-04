
package pkgs

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"

	"time"
	"log"
)

var _db   *pg.DB

func GetDB() *pg.DB {
	if _db == nil {
		addr := viper.GetString("postgresql.addr")
		user := viper.GetString("postgresql.user")
		password := viper.GetString("postgresql.password")
		dbname := viper.GetString("postgresql.dbname")
	
		_db = pg.Connect(&pg.Options{
			Addr: addr,
			User: user,
			Password: password,
			Database: dbname,
			DialTimeout: 30 * time.Second,
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 5 * time.Second,
			PoolSize: 10,
		})
		if _db == nil {
			log.Panicf("connect db failed(%s,%s,%s)", addr, user, dbname)
		}

		log.Printf("connect db success")
	}

	return _db
}
