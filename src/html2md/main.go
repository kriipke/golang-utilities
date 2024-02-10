package main 

import (
	"fmt"
	"log"
    	"golang.org/x/net/html"
    	"net/http"
	"reflect"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func convert(html string ) {

	converter := md.NewConverter("", true, nil)


	markdown, err := converter.ConvertString(html)
	if err != nil {
  		log.Fatal(err)
	}
	fmt.Println(markdown)
}


func main() {
    // Send an HTTP GET request to the example.com web page
    resp, err := http.Get("https://www.example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Use the html package to parse the response body from the request
    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    
    fmt.Println(doc)
    fmt.Println(reflect.TypeOf(doc))
    //convert(string(doc))
    //// Find and print all links on the web page
    //var links []string
    //var link func(*html.Node)
    //link = func(n *html.Node) {
    //    if n.Type == html.ElementNode && n.Data == "a" {
    //        for _, a := range n.Attr {
    //            if a.Key == "href" {
    //                // adds a new link entry when the attribute matches
    //                links = append(links, a.Val)
    //            }
    //        }
    //    }

    //    // traverses the HTML of the webpage from the first child node
    //    for c := n.FirstChild; c != nil; c = c.NextSibling {
    //        link(c)
    //    }
    //}
    //link(doc)

    //// loops through the links slice
    //for _, l := range links {
    //    fmt.Println("Link:", l)
    //}
}
