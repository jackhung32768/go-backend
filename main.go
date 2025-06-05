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


func adder_handler(w http.ResponseWriter, r *http.Request) {
    a := r.URL.Query().Get("a")
    b := r.URL.Query().Get("b")

    if a == "" || b == "" {
        http.ServeFile(w, r, "index.html")
        return
    }

    url:=""
    port := os.Getenv("PORT")
    if port == "" {
//        port = "8080"
      url = fmt.Sprintf("http://localhost:5000/calc?a=%s&b=%s", a, b)
    }else{
      url = fmt.Sprintf("https://python-api-5rg4.onrender.com/calc?a=%s&b=%s", a, b)
	}
//      url = fmt.Sprintf("https://python-api-5rg4.onrender.com/calc?a=%s&b=%s", a, b)
    fmt.Println(url)
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, "無法連接 Python API", 500)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    // 定義結果結構
type CalcResult struct {
    A      string
    B      string
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
    <title>兩數相加</title>
</head>
<body>
    <h1>兩數相加</h1>

    <form method="get" action="/adder">
        <input type="text" name="a" placeholder="輸入 a" value="{{.A}}">
        +
        <input type="text" name="b" placeholder="輸入 b" value="{{.B}}">
        <button type="submit">計算</button>
    </form>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Result}}
        <p>加總結果是：{{formatFloat .Result}}</p>
    {{end}}
	<br>
    <p><a href="/">⬅ 回主頁</a></p>
</body>
</html>
`

    t := template.Must(template.New("calctmpl").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func temperature_converter(w http.ResponseWriter, r *http.Request) {
    celsius := r.URL.Query().Get("celsius")
    fahrenheit := r.URL.Query().Get("fahrenheit")

   // ✅ 沒有輸入參數，顯示靜態 HTML 頁面
    if celsius == "" && fahrenheit == "" {
        http.ServeFile(w, r, "index_temp.html")
        return
    }


    type TempResult struct {
        Celsius        *float64 `json:"celsius,omitempty"`
        Fahrenheit     *float64 `json:"fahrenheit,omitempty"`
        Error          string   `json:"error,omitempty"`
        RawCelsius     string
        RawFahrenheit  string
    }

    result := TempResult{
        RawCelsius:    celsius,
        RawFahrenheit: fahrenheit,
    }

    // 如果有輸入，呼叫 Python API
    if celsius != "" || fahrenheit != "" {


    url:=""
    port := os.Getenv("PORT")
    if port == "" {

        url = fmt.Sprintf(
            "http://localhost:5000/convert_between_celsius_and_fahrenheit?celsius=%s&fahrenheit=%s",
            celsius, fahrenheit)



    }else{
        url = fmt.Sprintf(
            "https://python-api-5rg4.onrender.com/convert_between_celsius_and_fahrenheit?celsius=%s&fahrenheit=%s",
            celsius, fahrenheit)

	}


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

    const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>攝氏 / 華氏轉換</title>
</head>
<body>
    <h1>攝氏 &lt;=&gt; 華氏 轉換器</h1>

    <form method="get" action="/temperature_converter">
        攝氏：
        <input type="text" name="celsius" placeholder="輸入攝氏" value="{{.RawCelsius}}">
        或
        華氏：
        <input type="text" name="fahrenheit" placeholder="輸入華氏" value="{{.RawFahrenheit}}">
        <button type="submit">轉換</button>
    </form>

    <br>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Celsius}}
        <p>攝氏：{{formatFloat .Celsius}} °C</p>
    {{else if .Fahrenheit}}
        <p>華氏：{{formatFloat .Fahrenheit}} °F</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>
</body>
</html>
`

    t := template.Must(template.New("temp").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func inch_cm_converter(w http.ResponseWriter, r *http.Request) {
    inch := r.URL.Query().Get("inch")
    cm := r.URL.Query().Get("cm")

    // 無輸入時顯示靜態頁面
    if inch == "" && cm == "" {
        http.ServeFile(w, r, "index_inch_cm.html")
        return
    }

    type Result struct {
        Inch       *float64 `json:"inch,omitempty"`
        CM         *float64 `json:"cm,omitempty"`
        Error      string   `json:"error,omitempty"`
        RawInch    string
        RawCM      string
    }

    result := Result{
        RawInch: inch,
        RawCM:   cm,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_inch_and_cm?inch=%s&cm=%s", inch, cm)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_inch_and_cm?inch=%s&cm=%s", inch, cm)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
    <title>英吋 / 公分 轉換</title>
</head>
<body>
    <h1>英吋 &lt;=&gt; 公分 轉換器</h1>

    <form method="get" action="/inch_cm_converter">
        英吋：
        <input type="text" name="inch" placeholder="輸入英吋" value="{{.RawInch}}">
        或
        公分：
        <input type="text" name="cm" placeholder="輸入公分" value="{{.RawCM}}">
        <button type="submit">轉換</button>
    </form>

    <br>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Inch}}
        <p>英吋：{{formatFloat .Inch}} in</p>
    {{else if .CM}}
        <p>公分：{{formatFloat .CM}} cm</p>
    {{end}}

    <br>

    <p><a href="/">⬅ 回主頁</a></p>

</body>
</html>
`

    t := template.Must(template.New("inchcm").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func mile_km_converter(w http.ResponseWriter, r *http.Request) {
    mile := r.URL.Query().Get("mile")
    km := r.URL.Query().Get("km")

    if mile == "" && km == "" {
        http.ServeFile(w, r, "index_mile_km.html")
        return
    }

    type Result struct {
        Mile     *float64 `json:"mile,omitempty"`
        KM       *float64 `json:"km,omitempty"`
        Error    string   `json:"error,omitempty"`
        RawMile  string
        RawKM    string
    }

    result := Result{
        RawMile: mile,
        RawKM:   km,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_mile_and_km?mile=%s&km=%s", mile, km)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_mile_and_km?mile=%s&km=%s", mile, km)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
    <title>英哩 / 公里 轉換</title>
</head>
<body>
    <h1>英哩 &lt;=&gt; 公里 轉換器</h1>

    <form method="get" action="/mile_km_converter">
        英哩：
        <input type="text" name="mile" placeholder="輸入英哩" value="{{.RawMile}}">
        或
        公里：
        <input type="text" name="km" placeholder="輸入公里" value="{{.RawKM}}">
        <button type="submit">轉換</button>
    </form>

    <br>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Mile}}
        <p>英哩：{{formatFloat .Mile}} mi</p>
    {{else if .KM}}
        <p>公里：{{formatFloat .KM}} km</p>
    {{end}}

    <br>

    <p><a href="/">⬅ 回主頁</a></p>

</body>
</html>
`

    t := template.Must(template.New("milekm").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func meter_foot_converter(w http.ResponseWriter, r *http.Request) {
    meter := r.URL.Query().Get("meter")
    foot := r.URL.Query().Get("foot")

    if meter == "" && foot == "" {
        http.ServeFile(w, r, "index_meter_foot.html")
        return
    }

    type Result struct {
        Meter    *float64 `json:"meter,omitempty"`
        Foot     *float64 `json:"foot,omitempty"`
        Error    string   `json:"error,omitempty"`
        RawMeter string
        RawFoot  string
    }

    result := Result{
        RawMeter: meter,
        RawFoot:  foot,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_meter_and_foot?meter=%s&foot=%s", meter, foot)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_meter_and_foot?meter=%s&foot=%s", meter, foot)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
<head><meta charset="UTF-8"><title>公尺 / 英尺 轉換</title></head>
<body>
    <h1>公尺 &lt;=&gt; 英尺 轉換器</h1>

    <form method="get" action="/meter_foot_converter">
        公尺：
        <input type="text" name="meter" placeholder="輸入公尺" value="{{.RawMeter}}">
        或
        英尺：
        <input type="text" name="foot" placeholder="輸入英尺" value="{{.RawFoot}}">
        <button type="submit">轉換</button>
    </form>
    <br>
    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Meter}}
        <p>公尺：{{formatFloat .Meter}} m</p>
    {{else if .Foot}}
        <p>英尺：{{formatFloat .Foot}} ft</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>

</body>
</html>
`
    t := template.Must(template.New("meterfoot").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func yard_meter_converter(w http.ResponseWriter, r *http.Request) {
    yard := r.URL.Query().Get("yard")
    meter := r.URL.Query().Get("meter")

    if yard == "" && meter == "" {
        http.ServeFile(w, r, "index_yard_meter.html")
        return
    }

    type Result struct {
        Yard     *float64 `json:"yard,omitempty"`
        Meter    *float64 `json:"meter,omitempty"`
        Error    string   `json:"error,omitempty"`
        RawYard  string
        RawMeter string
    }

    result := Result{
        RawYard:  yard,
        RawMeter: meter,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_yard_and_meter?yard=%s&meter=%s", yard, meter)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_yard_and_meter?yard=%s&meter=%s", yard, meter)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
<head><meta charset="UTF-8"><title>碼 / 公尺 轉換</title></head>
<body>
    <h1>碼 &lt;=&gt; 公尺 轉換器</h1>

    <form method="get" action="/yard_meter_converter">
        碼：
        <input type="text" name="yard" placeholder="輸入碼" value="{{.RawYard}}">
        或
        公尺：
        <input type="text" name="meter" placeholder="輸入公尺" value="{{.RawMeter}}">
        <button type="submit">轉換</button>
    </form>
    <br>
    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Yard}}
        <p>碼：{{formatFloat .Yard}} yd</p>
    {{else if .Meter}}
        <p>公尺：{{formatFloat .Meter}} m</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>

</body>
</html>
`
    t := template.Must(template.New("yardmeter").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func kg_lb_converter(w http.ResponseWriter, r *http.Request) {
    kg := r.URL.Query().Get("kg")
    lb := r.URL.Query().Get("lb")

    if kg == "" && lb == "" {
        http.ServeFile(w, r, "index_kg_lb.html")
        return
    }

    type Result struct {
        KG     *float64 `json:"kg,omitempty"`
        LB     *float64 `json:"lb,omitempty"`
        Error  string   `json:"error,omitempty"`
        RawKG  string
        RawLB  string
    }

    result := Result{
        RawKG: kg,
        RawLB: lb,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_kg_and_lb?kg=%s&lb=%s", kg, lb)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_kg_and_lb?kg=%s&lb=%s", kg, lb)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
<head><meta charset="UTF-8"><title>公斤 / 磅 轉換</title></head>
<body>
    <h1>公斤 &lt;=&gt; 磅 轉換器</h1>

    <form method="get" action="/kg_lb_converter">
        公斤：
        <input type="text" name="kg" placeholder="輸入公斤" value="{{.RawKG}}">
        或
        磅：
        <input type="text" name="lb" placeholder="輸入磅" value="{{.RawLB}}">
        <button type="submit">轉換</button>
    </form>
    <br>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .KG}}
        <p>公斤：{{formatFloat .KG}} kg</p>
    {{else if .LB}}
        <p>磅：{{formatFloat .LB}} lb</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>
</body>
</html>
`

    t := template.Must(template.New("kglb").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func g_oz_converter(w http.ResponseWriter, r *http.Request) {
    g := r.URL.Query().Get("g")
    oz := r.URL.Query().Get("oz")

    if g == "" && oz == "" {
        http.ServeFile(w, r, "index_g_oz.html")
        return
    }

    type Result struct {
        G     *float64 `json:"g,omitempty"`
        OZ    *float64 `json:"oz,omitempty"`
        Error string   `json:"error,omitempty"`
        RawG  string
        RawOZ string
    }

    result := Result{
        RawG:  g,
        RawOZ: oz,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_g_and_oz?g=%s&oz=%s", g, oz)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_g_and_oz?g=%s&oz=%s", g, oz)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
<head><meta charset="UTF-8"><title>公克 / 盎司 轉換</title></head>
<body>
    <h1>公克 &lt;=&gt; 盎司 轉換器</h1>

    <form method="get" action="/g_oz_converter">
        公克：
        <input type="text" name="g" placeholder="輸入公克" value="{{.RawG}}">
        或
        盎司：
        <input type="text" name="oz" placeholder="輸入盎司" value="{{.RawOZ}}">
        <button type="submit">轉換</button>
    </form>
    <br>
    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .G}}
        <p>公克：{{formatFloat .G}} g</p>
    {{else if .OZ}}
        <p>盎司：{{formatFloat .OZ}} oz</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>
</body>
</html>
`

    t := template.Must(template.New("goz").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func tonne_ton_converter(w http.ResponseWriter, r *http.Request) {
    tonne := r.URL.Query().Get("tonne")
    ton := r.URL.Query().Get("ton")

    if tonne == "" && ton == "" {
        http.ServeFile(w, r, "index_tonne_ton.html")
        return
    }

    type Result struct {
        Tonne    *float64 `json:"tonne,omitempty"`
        Ton      *float64 `json:"ton,omitempty"`
        Error    string   `json:"error,omitempty"`
        RawTonne string
        RawTon   string
    }

    result := Result{
        RawTonne: tonne,
        RawTon:   ton,
    }

    url := ""
    port := os.Getenv("PORT")
    if port == "" {
        url = fmt.Sprintf("http://localhost:5000/convert_between_tonne_and_ton?tonne=%s&ton=%s", tonne, ton)
    } else {
        url = fmt.Sprintf("https://python-api-5rg4.onrender.com/convert_between_tonne_and_ton?tonne=%s&ton=%s", tonne, ton)
    }

    resp, err := http.Get(url)
    if err != nil {
        result.Error = "無法連接 Python API"
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    }

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
<head><meta charset="UTF-8"><title>公噸 / 英噸 轉換</title></head>
<body>
    <h1>公噸 &lt;=&gt; 英噸 轉換器</h1>

    <form method="get" action="/tonne_ton_converter">
        公噸：
        <input type="text" name="tonne" placeholder="輸入公噸" value="{{.RawTonne}}">
        或
        英噸：
        <input type="text" name="ton" placeholder="輸入英噸" value="{{.RawTon}}">
        <button type="submit">轉換</button>
    </form>
    <br>

    {{if .Error}}
        <p style="color:red;">錯誤：{{.Error}}</p>
    {{else if .Tonne}}
        <p>公噸：{{formatFloat .Tonne}} 公噸</p>
    {{else if .Ton}}
        <p>英噸：{{formatFloat .Ton}} 英噸</p>
    {{end}}

    <br>
    <p><a href="/">⬅ 回主頁</a></p>
</body>
</html>
`

    t := template.Must(template.New("tonneton").Funcs(funcMap).Parse(tmpl))
    t.Execute(w, result)
}

func home_page(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index_all.html")
}


func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
//inch_cm_converter
    http.HandleFunc("/", home_page)
    http.HandleFunc("/adder", adder_handler)
    http.HandleFunc("/temperature_converter", temperature_converter)
    http.HandleFunc("/inch_cm_converter", inch_cm_converter)
    http.HandleFunc("/mile_km_converter", mile_km_converter)
    http.HandleFunc("/meter_foot_converter", meter_foot_converter)
    http.HandleFunc("/yard_meter_converter", yard_meter_converter)
    http.HandleFunc("/kg_lb_converter", kg_lb_converter)
    http.HandleFunc("/g_oz_converter", g_oz_converter)
    http.HandleFunc("/tonne_ton_converter", tonne_ton_converter)
	
    fmt.Println("Go 伺服器啟動：localhost:8080")
//    log.Fatal(http.ListenAndServe(":8080", nil))
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
//    log.Fatal(http.ListenAndServe(":" + port, nil))
    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}