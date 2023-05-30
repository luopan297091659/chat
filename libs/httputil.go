package libs

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var host string = getvalue("conf/app.conf", "bmc_host")
var api_key string = getvalue("conf/app.conf", "api_key")
var timestamp string = strconv.FormatInt(time.Now().Unix(), 10)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func HttpRequest(method string, url_path string, msg string) string {
	url := host + url_path
	s := method + "\n" + url_path + "\n" + timestamp + "\n" + api_key
	md5Str := MD5(s)
	md5_value := api_key + ":" + md5Str
	// 超时时间：5秒
	//log.Print(host,"\n",url,"\n",s,"\n",md5_value)
	body := bytes.NewBuffer([]byte(msg))
	client := &http.Client{Timeout: 60 * time.Second}
	request, err := http.NewRequest(method, url, body)
	request.Header.Set("Date", timestamp)
	request.Header.Set("Authorization", md5_value)
	response, err := client.Do(request)
	if err != nil {
		log.Print("request err")
	}
	defer response.Body.Close()
	r, err := ioutil.ReadAll(response.Body)
	//log.Print(string(r))
	if response.StatusCode == 200 {
		//fmt.Println(err)
		log.Print(response.StatusCode, " request success")
	}
	return string(r)
}

func Get(url_path string, msg string) string {
	return HttpRequest("GET", url_path, msg)
}

func Post(url_path string, msg string) string {
	return HttpRequest("POST", url_path, msg)
}
