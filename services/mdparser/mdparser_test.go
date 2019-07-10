package mdparser

import (
	"fmt"
	"testing"
)

func TestArticle_ScanMarkdownFiles(t *testing.T) {
	article, e := NewMdParser("../../markdown")
	if e != nil {
		t.Fatal(e)
	}


	strings, e := article.ScanMarkdownFiles()
	if e  != nil {
		t.Fatal(e)
	}

	t.Log(strings)

}

func TestArticle_ParseMarkdown(t *testing.T) {
	article, e := NewMdParser("../../markdown")
	if e != nil {
		t.Fatal(e)
	}

	markdown, e := article.ParseMarkdown(article.Filepath + "/1.md")

	if e != nil {
		t.Fatal(e)
	}


	t.Log(markdown)
}

func TestArticle_InitArticles(t *testing.T) {
	parser, e := NewMdParser("../../markdown")
	if e != nil {
		t.Fatal(e)
	}

	e = parser.InitArticles()
	if e != nil {
		t.Fatal(e)
	}
	for _, article := range Articles {
		fmt.Println(article)
	}


}