package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
    "os"
    "io/ioutil"
    "router"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
/*
    http.Handle("/", http.FileServer(http.Dir("/web")))

    http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request){
        
        ver, err := ioutil.ReadFile("/config/VERSION")
        check(err)
        fmt.Fprintf(w,string(ver))

    })
*/
    r := router.Router()

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))

}
