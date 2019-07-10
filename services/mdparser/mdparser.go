package mdparser

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"xcmdblog/services/mdnotify"
)

type Article struct {
	Filepath string
	FileName string
	//类型名称
	TypeName string
	//标题
	Title string
	//副标题
	SubTitle string
	//作者
	Author string
	//提交时间
	PostDate string
	//更新时间
	UpdateDate string
	//展示图片
	ImageShow string
	//文档内容
	Content string
}

type MdParser struct {
	Filepath string
}

var Articles []*Article
var Obj *MdParser

func (a *Article) List() []*Article {
	return Articles
}

func NewMdParser(mdpath string) (*MdParser, error) {
	if Articles != nil {
		return Obj, nil
	}

	path, e := filepath.Abs(mdpath)
	if e != nil {
		log.Println(e.Error())
		return nil, e
	}

	Obj := &MdParser{
		Filepath: path,
	}

	return Obj, nil
}

func (a *MdParser) ParseMarkdown(filename string) (*Article, error) {
	file, e := os.OpenFile(filename, os.O_RDONLY, 0777)
	if e != nil {
		return nil, e
	}

	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		return nil, e
	}

	article := Article{}

	split := strings.Split(string(bytes), "\n")

	for i := 0; i < len(split); i++ {
		clearItem := strings.TrimFunc(split[i], func(r rune) bool {
			return string(r) == "<" || string(r) == ">"
		})

		//demo := "Title INSERT... ON DUPLICATE KEY UPDATE引起的问题"
		firstBlank  := strings.Index(clearItem, " ")
		if firstBlank >0 {
			name, value := clearItem[0:firstBlank], clearItem[firstBlank + 1 :]

			if article.saveProperty(name, value) {

				if i == 0 {
					split = split[1:]
					i--
				} else if i == len(split)-1 {
					split = split[:i]
					i--
				} else {
					split = append(split[0:i], split[i+1:]...)
					i--
				}
			}
		}

	}

	article.Content = strings.Join(split, "\n")
	return &article, nil

}

func (a *MdParser) InitArticles() error {
	if Articles != nil {
		 return nil
	}

	files, e := a.ScanMarkdownFiles()
	if e != nil {
		return e
	}

	for _,name := range files{
		article, e := a.ParseMarkdown(a.Filepath + "/" + name)
		if e != nil {
			return e
		}

		if article.Title == "" || article.SubTitle =="" {
			continue
		}
		article.FileName = name
		Articles = append(Articles, article)
	}

	//监听文件夹
	notify, e := mdnotify.New(a.Filepath)
	if e != nil {
		return e
	}

	go notify.Watch(a)

	return nil
}

func (a *MdParser) removeArticle(filename string) {
	log.Println("remove", filename)
	for i:=0;i<len(Articles); i++ {
		if Articles[i].FileName == filename {
			Articles = append(Articles[0:i], Articles[i+1:] ...)
			break
		}
	}
}

func (a *MdParser) appendArticle(filename string) error {
	article, e := a.ParseMarkdown(a.Filepath + "/" + filename)
	if e != nil {
		return e
	}
	if article.Title == "" || article.SubTitle =="" {
		return errors.New("property missing")
	}


	article.FileName = filename

	for i,a := range Articles {
		if a.FileName == article.FileName {
			Articles[i] = article
			return nil
		}
	}
	Articles = append(Articles, article)

	return nil
}

func (a *MdParser) Create(name string) {
	e := a.appendArticle(filepath.Base(name))
	log.Println(e)
}

func (a *MdParser) Write(name string) {

	e := a.appendArticle(filepath.Base(name))
	log.Println(e)
}

func (a *MdParser) Remove(name string) {

	a.removeArticle(filepath.Base(name))
}

func (a *MdParser) Rename(name string) {

	a.removeArticle(filepath.Base(name))
}

func (a *Article) saveProperty(key string, value string) bool {
	isProperty := true
	switch key {
	case "Title":
		a.Title = value
	case "SubTitle":
		a.SubTitle = value
	case "Author":
		a.Author = value
	case "PostDate":
		a.PostDate = value
	case "UpdateDate":
		a.UpdateDate = value
	case "ImageShow":
		a.ImageShow = value
	default:
		isProperty = false
	}
	return isProperty
}

func (a *MdParser) ScanMarkdownFiles() ([]string, error) {

	file, e := os.Open(a.Filepath)
	if e != nil {
		log.Println(e.Error())
		return nil, e
	}

	infos, e := file.Readdir(1000)
	if e != nil {
		log.Println(e.Error())
		return nil, e
	}

	ret := []string{}

	for _, v := range infos {
		ret = append(ret, v.Name())
		log.Println(v.Name())
	}

	return ret, nil
}
