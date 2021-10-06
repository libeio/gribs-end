
package main

import (
	"grids.com/grids"

	"github.com/spf13/viper"

	"html/template"
	"log"
)

type MainPage struct {
	Head 	grids.Head
	Header	grids.Header
    Nav	    grids.Nav
    Footer	grids.Footer
	Body	[]TblPhotoShow
}

var mainPage *template.Template
var data *MainPage


func InitPage() {
	var tblPhotoShows []TblPhotoShow
	db := grids.GetDB()
	err := db.Model(&tblPhotoShows).Column("href","src","description").Order("sorted").Limit(10).Select()
	if err != nil {
		log.Panicln(err)
	}

	data = &MainPage {
		Head: grids.Head {
			Title: "photo",
			Styles: viper.GetStringSlice("frontend.staticStyleFiles"),
			Scripts: viper.GetStringSlice("frontend.staticScriptFiles"),
		},
		Header: grids.Header {
			ImgSrc: viper.GetString("photo.headerImgSrc"),
		},
		Nav: grids.Nav {
            Lists: []grids.NavList {
				{ Href: viper.GetString("suites.photo.href"), Text: viper.GetString("suites.photo.text") },
				{ Href: viper.GetString("suites.cardpkg.href"), Text: viper.GetString("suites.cardpkg.text") },
				{ Href: viper.GetString("suites.help.href"), Text: viper.GetString("suites.help.text") },
				{ Href: viper.GetString("suites.rule.href"), Text: viper.GetString("suites.rule.text") },
				{ Href: viper.GetString("suites.monster.href"), Text: viper.GetString("suites.monster.text") },
				{ Href: viper.GetString("suites.magics.href"), Text: viper.GetString("suites.magics.text") },
				{ Href: viper.GetString("suites.limits.href"), Text: viper.GetString("suites.limits.text") },
            },
		},
		Body: tblPhotoShows,
		Footer: grids.Footer {
			Text: viper.GetString("resource.footerText"),
		},
	}
}
