package main

import (
	"github.com/agilab/gostone/agiutils"
	"flag"
	//"github.com/book_prereader/resource"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	//"github.com/book_prereader/tables"
	"log"
)

var c = make(chan bool)

func main() {
	var (
		configFile = flag.String("config_file", "./configs/config.yaml", "配置文件地址")
	)
	flag.Parse()
	agiutils.LoadViperFromFiles(*configFile)
//	resource.InitResource()

	//resource.DB.AutoMigrate(&tables.Book{})

	var (
		url string = "http://m.xs.la/81_81260/"
	)
	//使用 goquery 创建 dom 对象
	dom, query_err := goquery.NewDocument(url)
	if query_err != nil {
		fmt.Println(query_err)
	}

	dhead := dom.Find("head")
	dcharset := dhead.Find("meta[http-equiv]")
	charset, _ := dcharset.Attr("content")
	log.Println("charset--"+charset)

//	title := dom.Find(".title").Text()
	image,err:=dhead.Find("meta[property]").Attr("image")
	log.Println(err)
	image1:=dhead.Find(".image").Text()
	hehe,e:=dhead.Find("meta[property]").Attr("content")
	log.Println("hehe--"+hehe)
	log.Println(e)
	lastname:=dhead.Find(".novel:latest_chapter_name").Text()
	log.Println("image--"+image)
	log.Println("image--"+image1)
	log.Println("lastname--"+lastname)
	//if err:=resource.DB.Debug().Table("books").Create(&tables.Book{
	//	Name: title,
	//	Url:  url,
	//}).Error;err!=nil{
	//	fmt.Println(err.Error())
	//}
	select {}
}
