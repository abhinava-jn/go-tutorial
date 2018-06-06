package main

import("fmt"
		"net/http"
		"io/ioutil"
		)
type SiteMapIndex struct{
	Location []Location 'xml:"sitemap"'
}
type Location struct{
	Loc string 'xml:"loc"'
}

func main(){
	resp,_ :=http.Get("https://askubuntu.com/questions/179681/cd-and-cd-commands-not-found-how-do-i-use-the-cd-command")
	bytes,_:=ioutil.ReadAll(resp.Body)
	string_body :=string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Location)

}