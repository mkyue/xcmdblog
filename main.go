package main

import (
	"log"
	_ "xcmdblog/routers"
	"github.com/astaxie/beego"
	"xcmdblog/services/mdparser"
)

func main() {
	parser, e := mdparser.NewMdParser("./markdown")
	if e != nil {
		log.Fatalln(e)
	}

	e = parser.InitArticles()
	if e != nil {
		log.Fatalln(e)
	}
	log.Println("ok")
	beego.Run()
}

