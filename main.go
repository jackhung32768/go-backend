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

func temp_handler(w http.ResponseWriter, r *http.Request) {
    celsius := r.URL.Query().Get("celsius")
    fahrenheit := r.URL.Query().Get("fahrenheit")

    if celsius == "" && fahrenheit == "" {
        http.ServeFile(w, r, "index_temp.html")
        return
    }

    url := fmt.Sprintf(
        "https://python-api-5rg4.onrender.com/convert_between_celsius_and_fahrenheit?celsius=%s&fahrenheit=%s",
        celsius, fahrenheit)

    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, "無法連接 Python API", 500)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

//    fmt.Fprintf(w, "Python 回傳結果：%s", string(body))
    type TempResult struct {
        Celsius    *float64 `json:"celsius,omitempty"`
        Fahrenheit *float64 `json:"fahrenheit,omitempty"`
        Error      string   `json:"error,omitempty"`
    }

    var result TempResult
    json.Unmarshal(body, &result)

    // 定義 HTML 模板
    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>溫度轉換結果</title>
</head>
<body>
    <h1>轉換結果</h1>
    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else}}
        {{if .Celsius}}<p>攝氏：{{printf "%.2f" .Celsius}} °C</p>{{end}}
        {{if .Fahrenheit}}<p>華氏：{{printf "%.2f" .Fahrenheit}} °F</p>{{end}}
    {{end}}
    <br>
    <a href="/temperature_convert">🔙 回到轉換頁</a>
</body>
</html>
    `

    t := template.Must(template.New("result").Parse(tmpl))
    t.Execute(w, result)

}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/temperature_convert", temp_handler)
    fmt.Println("Go 伺服器啟動：localhost:8080")
//    log.Fatal(http.ListenAndServe(":8080", nil))
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
//    log.Fatal(http.ListenAndServe(":" + port, nil))
    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}