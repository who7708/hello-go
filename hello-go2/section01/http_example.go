package section01

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpGetTest() {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		// handle error
	}
	// 程序在使用完回复后必须关闭回复的主体。
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error")
		}
	}(response.Body)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func HttpPostTest() {
	body := "{\"action\":20}"
	res, err := http.Post("http://xxx.com", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func HttpNewClientTest() {
	v := url.Values{}
	v.Set("username", "xxxx")
	v.Set("password", "xxxx")
	// 利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，
	// 并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	body := io.NopCloser(strings.NewReader(v.Encode())) // 把form数据编下码
	client := &http.Client{}                            // 客户端,被Get,Head以及Post使用
	reqest, err := http.NewRequest("POST", "http://xxx.com/logindo", body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	// 给一个key设定为响应的value.
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") // 必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest) // 发送请求
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Close Error ", err.Error())
		}
	}(resp.Body) // 一定要关闭resp.Body
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		fmt.Println("Write error", err.Error())
	}
}

func HttpWebTest() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println("http create error", err.Error())
	}

}

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在执行应用处理器前编译中间件的逻辑
		next.ServeHTTP(w, r)
		// 在执行应用处理器后将会执行的编译中间件的逻辑
	})
}
