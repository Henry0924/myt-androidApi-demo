package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	command = flag.String("command", "snap", "操作指令, snap--获取屏幕截图, devRandom--随机设备信息, devCustom--自定义设备信息, hideApp--隐藏应用")
	port    = flag.String("p", "", "安卓容器接口的端口，例9082，10005")
	host    = flag.String("host", "", "安卓host")
	level   = flag.Int("l", 3, "获取屏幕截图时的可选参数，可选值有1,2,3")
	app     = flag.String("app", "", "隐藏app时的可选参数，可填多个用逗号分割，例com.ss.android.ugc.aweme,com.mth_player.oaid")
)

func GetSnap(addr string, level int) {
	resp, err := http.Get(fmt.Sprintf("http://%s/task=snap&level=%d", addr, level))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fmt.Sprintf("%d.png", time.Now().Unix()), os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		return
	}
}

func ModifyDevRandom(addr string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/modifydev?cmd=2", addr))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}

func ModifyDevCustomized(addr string) {
	content, err := os.ReadFile("dev.json")
	if err != nil {
		log.Fatal(err)
	}

	dev := base64.StdEncoding.EncodeToString(content)
	resp, err := http.PostForm(fmt.Sprintf("http://%s/modifydev", addr),
		url.Values{"cmd": {"1"}, "data": {dev}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}

func HideApp(addr string, app string) {
	appSplit := strings.Split(app, ",")
	marshal, err := json.Marshal(appSplit)
	if err != nil {
		log.Fatal(err)
	}

	dev := base64.StdEncoding.EncodeToString(marshal)
	resp, err := http.PostForm(fmt.Sprintf("http://%s/modifydev", addr),
		url.Values{"cmd": {"3"}, "data": {dev}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatal("安卓host不能为空")
	}

	if *port == "" {
		log.Fatal("端口不能为空")
	}

	addr := fmt.Sprintf("%s:%s", *host, *port)

	switch *command {
	case "snap":
		GetSnap(addr, *level)
	case "devRandom":
		ModifyDevRandom(addr)
	case "devCustom":
		ModifyDevCustomized(addr)
	case "hideApp":
		HideApp(addr, strings.TrimSpace(*app))
	default:
		log.Fatal("command错误")
	}
}
