package routers

import (
	"xcmdblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ArticleController{}, "*:List")
    beego.Router("/detail", &controllers.ArticleController{}, "*:Detail")
}
