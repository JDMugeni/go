package main

import ("fmt"
"net/http"
"log"
"io/ioutil"
"strings"
)

func (vh *viewHandler) ServerHTTP (w http.ResponseWriter, r *http.Request){

path:=r.URL.Path [1:]
log.Println (path);
data , err :=ioutil.ReadFile (string (path))

if err == nil {

var contentType string 

if strings.HasSuffix (path,".html"){
contentType ="text/html"
} else if strings.HasSuffix (path,"css"){
contentType ="text/css"
}else if strings.HasSuffix (path,"js"){
contentType ="application/javascript"
}else if strings.HasSuffix (path,"png"){
contentType = "image/png"
} else if strings.HasSuffix (path,"jpg"){
contentType ="image/jpg"
}
w.Header ().Add ("Content-type", contentType)
w.write (data)
}else {
log.Println ("ERROR")
w.WriteHeader (404)
w.write ([]byte ("404 - " + http.StatusText (404)))
}

}

func main (){
fmt.Printf ("Hello Jordy Mugeni , Welcome to Golang\n")
http.Handle ("/",new (viewHandler))
http.ListenAndServe (":8282",nil)
}
