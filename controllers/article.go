package controllers

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
	"xcmdblog/services/mdparser"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) List() {
	list := mdparser.Articles
	page, e := c.GetInt("page", 1)
	if e != nil || page < 1 {
		page = 1
	}
	start := (page - 1) * 3

	if start > len(list) {
		start = 0
	}
	end := start + 3
	if len(list) < end {
		end = len(list)
	}
	var ret []interface{}

	for _, article := range list[start:end]{
		ret = append(ret, map[string]string{
			"title" : article.Title,
			"sub_title" : article.SubTitle,
			"image_show" : article.ImageShow,
			"author" :article.Author,
			"type_name" : article.TypeName,
			"update_date" : article.UpdateDate,
			"post_date" : article.PostDate,
			"file_name" : strings.TrimRight(article.FileName, filepath.Ext(article.FileName)),
		})
	}
	c.Data["list"] = ret
	if end < len(list) {
		c.Data["nextpage"] = page + 1
	}else{
		c.Data["nextpage"] = 0
	}

	c.TplName = "index.html"
}


func (c *ArticleController) Detail() {
	name := c.GetString("id", "")
	if len(name) == 0 {
		c.Abort("404")
	}
	name += ".md"
	list := mdparser.Articles
	ret := &mdparser.Article{}
	for _, article := range list{
		if article.FileName == name {
			ret = article
			break
		}
	}
	if ret == nil {
		c.Abort("404")
	}

	c.Data["blog"] = map[string]string {
		"title" : ret.Title,
		"post_date" : ret.PostDate,
		"content" : ret.Content,

	}
	c.TplName = "detail.html"
}
