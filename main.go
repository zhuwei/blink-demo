package main

import (
	"errors"
	"log"

	"github.com/zhuwei/blink-demo/ui"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/raintean/blink"
)

func main() {
	//退出信号
	exit := make(chan bool)

	//设置调试模式
	blink.SetDebugMode(true)

	//初始化blink模块
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}

	//注册虚拟网络文件系统到域名app
	blink.RegisterFileSystem("app", &assetfs.AssetFS{
		Asset:     ui.Asset,
		AssetDir:  ui.AssetDir,
		AssetInfo: ui.AssetInfo,
	})

	//新建view,加载URL
	view := blink.NewWebView(false, 1366, 920)
	//直接加载虚拟文件系统中的网页
	view.LoadURL("http://app/html/index.html")
	view.SetWindowTitle("Golang GUI Application")
	view.MoveToCenter()
	view.MaximizeWindow()
	view.ShowWindow()
	view.ShowDevTools()
	view.On("destroy", func(_ *blink.WebView) {
		close(exit)
	})

	//golang注入值
	view.Inject("title", "Tool")

	//golang注入方法
	view.Inject("GetData", func(num int) (int, error) {
		if num > 10 {
			return 0, errors.New("num不能大于10")
		} else {
			return num + 1, nil
		}
	})

	//等待退出
	<-exit
}
