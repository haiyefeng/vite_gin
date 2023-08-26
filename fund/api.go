package fund

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type FundEstimatedValue struct {
	FundCode       string `json:"fundcode"`
	Name           string `json:"name"`
	ValDate        string `json:"jzrq"`
	Value          string `json:"dwjz"`
	EstimatedValue string `json:"gsz"`
	EstimatedRange string `json:"gszzl"`
	EstimatedTime  string `json:"gztime"`
}

func getCodeVal(code, timestamp string) (string, error) {
	apiUrl := "http://fundgz.1234567.com.cn/js/" + code + ".js?rt=" + timestamp
	result, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return "", err
	}
	defer result.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	return string(body), nil
}

func GetCodeVal() (*[]FundEstimatedValue, error) {
	currentTime := time.Now()
	milliseconds := currentTime.UnixNano() / 1e6
	timestamp := fmt.Sprintf("%d", milliseconds)

	codes := []string{
		"002834",
		"001614",
		"006234",
		"012769",
		"012043",
		"015736",
		"001678",
		"161725",
	}
	var fundDatas []FundEstimatedValue
	for _, code := range codes {

		var fundData FundEstimatedValue
		res, err := getCodeVal(code, timestamp)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}
		re := regexp.MustCompile(`(\{.*\})`)
		matches := re.FindAllStringSubmatch(res, -1)

		// 遍历匹配项并输出捕获的内容
		for _, match := range matches {
			if len(match) > 1 {
				err = json.Unmarshal([]byte(match[1]), &fundData)
				if err != nil {
					fmt.Println("Error unmarshaling JSON:", err)
					return nil, err
				}
				fundDatas = append(fundDatas, fundData)
				break
			}
		}
	}
	return &fundDatas, nil
}
