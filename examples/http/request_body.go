package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/nil", getBodyIsNil)
	http.HandleFunc("/body/param", queryParam)
	http.HandleFunc("/body/url", wholeUrl)
	http.HandleFunc("/body/form", form)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v \n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _:=json.Marshal(r.URL)
	fmt.Fprintf(w, "url: %s \n",data)
}

func queryParam(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "param is %+v \n", values)
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	// http request unused
	if r.GetBody == nil {
		fmt.Fprintf(w, "Get Body is nil \n")
	} else {
		fmt.Fprintf(w, "Get Body is not nil\n")
	}
}


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is home page\n")
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed %v\n", err)
		return
	}
	fmt.Fprintf(w, "read the data %s \n", string(body))
	// 再次尝试读取，无法读取到任何值
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed %v\n", err)
	}
	fmt.Fprintf(w, "read data one more time; [%s] and read data length %d \n", string(body), len(body))
}
