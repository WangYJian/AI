package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func POST(url string, data map[string]string, Header string) (map[string]interface{}, error){
	var c map[string]interface{}
	//发送请求获取token
	method := "POST"
	client := &http.Client{}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for key, value := range data {
		_ = writer.WriteField(key, value)
	}
	err := writer.Close()
	if Header == "" {
		Header = writer.FormDataContentType()
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", Header)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(body, &c)
    return c, err
}