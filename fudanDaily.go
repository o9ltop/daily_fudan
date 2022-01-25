/**
 * @Author Oliver
 * @Date 1/25/22
 **/

package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	client        = &http.Client{}
	fudanDailyUrl = "https://zlapp.fudan.edu.cn/site/ncov/fudanDaily"
	loginUrl      = "https://uis.fudan.edu.cn/authserver/login?service=https%3A%2F%2Fzlapp.fudan.edu.cn%2Fa_fudanzlapp%2Fapi" +
		"%2Fsso%2Findex%3Fredirect%3Dhttps%253A%252F%252Fzlapp.fudan.edu.cn%252Fsite%252Fncov%252FfudanDaily" +
		"%26from%3Dwap "
	getInfoUrl = "https://zlapp.fudan.edu.cn/ncov/wap/fudan/get-info"
	saveLogUrl = "https://zlapp.fudan.edu.cn/wap/log/save-log"
	saveUrl    = "https://zlapp.fudan.edu.cn/ncov/wap/fudan/save"
	userAgent  = "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.18(0x17001229) NetType/WIFI Language/zh_CN miniProgram"
	origin     = "https://zlapp.fudan.edu.cn"
	Referer    = fudanDailyUrl
)

type userInfo struct {
	Username string
	Password string
	Email    string
}

func setHeader(r *http.Request) {
	r.Header.Add("User-Agent", userAgent)
	r.Header.Add("Origin", origin)
	r.Header.Add("Referer", Referer)
}

func getUrlValue(form map[string]string) url.Values {
	res := url.Values{}
	for k, v := range form {
		res.Add(k, v)
	}
	return res
}

func login(info userInfo) {
	request, _ := http.NewRequest("GET", loginUrl, nil)
	resp, _ := client.Do(request)
	form := map[string]string{}
	body, _ := ioutil.ReadAll(resp.Body)
	h, _ := htmlquery.Parse(strings.NewReader(string(body)))
	a := htmlquery.Find(h, "//input")
	for i := range a {
		name := htmlquery.SelectAttr(a[i], "name")
		value := htmlquery.SelectAttr(a[i], "value")
		form[name] = value
	}
	form["username"] = info.Username
	form["password"] = info.Password
	request, _ = http.NewRequest("POST", loginUrl, ioutil.NopCloser(strings.NewReader(getUrlValue(form).Encode())))
	setHeader(request)
	resp, _ = client.Do(request)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	login(userInfo{
		Username: "20210240194",
		Password: "Liu159632",
	})
}
