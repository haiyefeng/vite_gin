package feishu

import (
	"fmt"
	"net/http"
	"strings"
)

func sendMsg(apiUrl, msg string) {
	// json
	contentType := "application/json"
	// data
	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + msg + `"}
	}`
	// request
	result, err := http.Post(apiUrl, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func SendData(msg string) {
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/01c7c8a8-0e00-4894-974d-c0d1faa3ff26"
	sendMsg(url, msg)
}
