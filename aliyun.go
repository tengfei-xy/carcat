package main

import (
	"env"
	"io/ioutil"
	"net/http"
	pnt "print"
	"strings"
)

func aliyunMain(picdata string) (string, bool) {

	aliyun_jdata := aliyunPost("https://ocrcp.market.alicloudapi.com/rest/160601/ocr/ocr_vehicle_plate.json", env.AppCode, aliyunSendData(picdata))
	pnt.Infof("aliyun over! appcode:%s,data:%s", env.AppCode, aliyun_jdata)
	carid, err := ifCar(aliyun_jdata)
	pnt.Infof("车牌为:%s", carid)
	return carid, err
}

func aliyunSendData(pic string) string {
	return `{
		"image": "` + pic + `",
		"configure": {"multi_crop":false}
	}`
}
func aliyunPost(link string, AppCode string, picdata string) string {
	client := &http.Client{}
	r, err := http.NewRequest("POST", link, strings.NewReader(picdata))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Authorization", "AppCode "+AppCode)
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	byteBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	return string(byteBody)
}
func ifCar(aliyun_jdata string) (string, bool) {
	var ai aliyun_info
	d := []byte(aliyun_jdata)
	err := parseJSON(&d, &ai)
	if err != nil {
		panic(err)
	}

	if ai.Success {
		c := len(ai.Plates)
		pnt.Infof("存在%d个返回结果", c)

		if c == 1 {
			return ai.Plates[0].Txt, ai.Success
		} else {
			return ai.Plates[0].Txt, ai.Success
		}
	}
	return "", ai.Success
}
