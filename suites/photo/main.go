
package main


import (
	"grids.com/grids"

	"github.com/spf13/viper"

	"log"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	viper.SetConfigName("conf/grids")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("read conf error: %s", err.Error())
	}

	grids.InitTable()
	InitDB()

}

func main() {
	
}