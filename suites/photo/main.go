
package main

import (
	"grids.com/grids"

	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
    "log"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	viper.SetConfigName("grids")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("read conf error: %s", err.Error())
	}

	grids.InitTable()
    InitDB()
    InitPage()
}

func tHandler(w http.ResponseWriter, r *http.Request) {
    mainPage.ExecuteTemplate(w, "photo", data)
}

func main() {
	router := gin.Default()

	photoListImgPath := viper.GetString("photo.listImgPath")
	router.Static(photoListImgPath[strings.LastIndex(photoListImgPath, "/"):], photoListImgPath)

	staticStylePath := viper.GetString("frontend.staticStylePath")
	staticScriptPath := viper.GetString("frontend.staticScriptPath")
	router.Static(staticStylePath[strings.LastIndex(staticStylePath, "/"):], staticStylePath)
	router.Static(staticScriptPath[strings.LastIndex(staticScriptPath, "/"):], staticScriptPath)

	router.LoadHTMLGlob(viper.GetString("backend.tmplPath"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "photo", data)
	})

	router.Run(":8080")
}