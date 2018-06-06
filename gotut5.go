package main

import("fmt"
	  "net/http"
	  "io/ioutil")

func main(){
	resp,_ :=http.Get("https://askubuntu.com/questions/179681/cd-and-cd-commands-not-found-how-do-i-use-the-cd-command")
	bytes,_:=ioutil.ReadAll(resp.Body)
	String_body := string(bytes)
	fmt.Println(String_body)
	resp.Body.Close()
	  
}

