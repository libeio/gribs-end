
package grids

import (
	"github.com/go-pg/pg/v10"

	"grids.com/pkgs"
)

func GetDB() *pg.DB {
	return pkgs.GetDB()
}
