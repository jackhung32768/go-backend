package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
)

/*
func handler(w http.ResponseWriter, r *http.Request) {
    a := r.URL.Query().Get("a")
    b := r.URL.Query().Get("b")

    url := fmt.Sprintf("http://localhost:5000/calc?a=%s&b=%s", a, b)

    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, "無法連接 Python API", 500)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    fmt.Fprintf(w, "Python 回傳結果：%s", string(body))
}
*/

func handler(w http.ResponseWriter, r *http.Request) {
    a := r.URL.Query().Get("a")
    b := r.URL.Query().Get("b")

    if a == "" || b == "" {
        http.ServeFile(w, r, "index.html")
        return
    }

//    url := fmt.Sprintf("http://localhost:5000/calc?a=%s&b=%s", a, b)
    url := fmt.Sprintf("https://python-api-5rg4.onrender.com/calc?a=%s&b=%s", a, b)

    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, "無法連接 Python API", 500)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    fmt.Fprintf(w, "Python 回傳結果：%s", string(body))
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Go 伺服器啟動：localhost:8080")
//    log.Fatal(http.ListenAndServe(":8080", nil))
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
//    log.Fatal(http.ListenAndServe(":" + port, nil))
    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}