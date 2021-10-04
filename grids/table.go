
package grids

import (
	"grids.com/table"
)

func InitTable() {
	table.InitPermission()
	table.InitGroup()
	table.InitUser()
}
