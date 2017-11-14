package main

import (
	"fmt"
	"log"
  
	"github.com/PuerkitoBio/goquery"
  )

var DOWNLOAD_URL string = "https://book.douban.com/top250"

func download_page(url string) *goquery.Document {
	doc,err := goquery.NewDocument(DOWNLOAD_URL)
	fmt.Println(doc.Html())
	if err != nil {
		log.Fatal(err)
	} 
	return doc
}

func parse_doc(doc *goquery.Document)(book_list []string ,new_url string){
	booklist := make([]string,0)

	booknamelist :=  doc.Find("#content").Find(".indent")
	booknamelist.Find("table").Each(func(i int, s *goquery.Selection){
		s.Find("div .pl2").Each(func(i int, s *goquery.Selection){
			bookname := s.Find("a").Text()
			booklist = append(booklist,bookname)
			fmt.Print(bookname)
		})
	})

	next_page := doc.Find("span .next").Find("a")
	fmt.Println(next_page.Text())
	if next_page != nil {
		return book_list,new_url
	} 

	return book_list,"nil"
}


func main(){
	url := DOWNLOAD_URL
	doc:= download_page(url)
	parse_doc(doc)
	
	
	// for url != "nil" {
	// 	books := make([]string,0)
	// 	doc := download_page(url)
	// 	books,url = parse_doc(doc)
	// 	fmt.Println(books)
	// }
}