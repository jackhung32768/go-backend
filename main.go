package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
    "encoding/json"
    "html/template"	
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

/*
func adder_handler(w http.ResponseWriter, r *http.Request) {
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

//    fmt.Fprintf(w, "Python 回傳結果：%s", string(body))
    // 定義結果結構
    type CalcResult struct {
        Result *float64 `json:"result,omitempty"`
        Error  string   `json:"error,omitempty"`
    }

    var result CalcResult
    json.Unmarshal(body, &result)

    // 自訂格式化函式
    funcMap := template.FuncMap{
        "formatFloat": func(f *float64) string {
            if f == nil {
                return ""
            }
            return fmt.Sprintf("%.2f", *f)
        },
    }

    const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>計算結果</title>
</head>
<body>
    <h1>兩數相加結果</h1>
    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else}}
        <p>加總結果是：{{formatFloat .Result}}</p>
    {{end}}
    <br>
    <a href="/adder">🔙 回到加法頁</a>
</body>
</html>
`

    t := template.Must(template.New("calctmpl").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func temperature_converter(w http.ResponseWriter, r *http.Request) {
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

    // 建立模板並加上函式（解除指標並格式化）
    funcMap := template.FuncMap{
        "formatFloat": func(f *float64) string {
            if f == nil {
                return ""
            }
            return fmt.Sprintf("%.2f", *f)
        },
    }

    const tmpl = `
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
        {{if .Celsius}}<p>攝氏：{{formatFloat .Celsius}} °C</p>{{end}}
        {{if .Fahrenheit}}<p>華氏：{{formatFloat .Fahrenheit}} °F</p>{{end}}
    {{end}}
    <br>
    <a href="/temperature_converter">🔙 回到轉換頁</a>
</body>
</html>
`

    t := template.Must(template.New("result").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}
*/

func adder_handler(w http.ResponseWriter, r *http.Request) {
    a := r.URL.Query().Get("a")
    b := r.URL.Query().Get("b")

    // 定義結果結構
    type CalcResult struct {
        A      string
        B      string
        Result *float64 `json:"result,omitempty"`
        Error  string   `json:"error,omitempty"`
    }

    result := CalcResult{A: a, B: b}

    if a != "" && b != "" {
        url := fmt.Sprintf("https://python-api-5rg4.onrender.com/calc?a=%s&b=%s", a, b)
        resp, err := http.Get(url)
        if err != nil {
            result.Error = "無法連接 Python API"
        } else {
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(body, &result)
        }
    }

    funcMap := template.FuncMap{
        "formatFloat": func(f *float64) string {
            if f == nil {
                return ""
            }
            return fmt.Sprintf("%.2f", *f)
        },
    }

/*
    t := template.Must(template.New("index").Funcs(funcMap).ParseFiles("index.html"))
//    t.Execute(w, result)
    if err := t.Execute(w, result); err != nil {
        log.Println("執行模板錯誤:", err)
        http.Error(w, "內部錯誤", http.StatusInternalServerError)
    }
	*/
    t, err := template.New("index").Funcs(funcMap).ParseFiles("index.html")
    if err != nil {
        log.Println("模板載入失敗:", err)
        http.Error(w, "模板載入錯誤", http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, result)
    if err != nil {
        log.Println("模板執行錯誤:", err)
        http.Error(w, "執行模板失敗", http.StatusInternalServerError)
    }
}

func temperature_converter(w http.ResponseWriter, r *http.Request) {
    celsius := r.URL.Query().Get("celsius")
    fahrenheit := r.URL.Query().Get("fahrenheit")

    type TempResult struct {
        Celsius    *float64 `json:"celsius,omitempty"`
        Fahrenheit *float64 `json:"fahrenheit,omitempty"`
        Error      string   `json:"error,omitempty"`
        RawCelsius string
        RawFahrenheit string
    }

    result := TempResult{RawCelsius: celsius, RawFahrenheit: fahrenheit}

    if celsius != "" || fahrenheit != "" {
        url := fmt.Sprintf(
            "https://python-api-5rg4.onrender.com/convert_between_celsius_and_fahrenheit?celsius=%s&fahrenheit=%s",
            celsius, fahrenheit)

        resp, err := http.Get(url)
        if err != nil {
            result.Error = "無法連接 Python API"
        } else {
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(body, &result)
        }
    }

    funcMap := template.FuncMap{
        "formatFloat": func(f *float64) string {
            if f == nil {
                return ""
            }
            return fmt.Sprintf("%.2f", *f)
        },
    }

    t := template.Must(template.New("index_temp").Funcs(funcMap).ParseFiles("index_temp.html"))
//    t.Execute(w, result)
    if err := t.Execute(w, result); err != nil {
        log.Println("執行模板錯誤:", err)
        http.Error(w, "內部錯誤", http.StatusInternalServerError)
    }
}


func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

    http.HandleFunc("/adder", adder_handler)
    http.HandleFunc("/temperature_converter", temperature_converter)
    fmt.Println("Go 伺服器啟動：localhost:8080")
//    log.Fatal(http.ListenAndServe(":8080", nil))
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
//    log.Fatal(http.ListenAndServe(":" + port, nil))
    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}