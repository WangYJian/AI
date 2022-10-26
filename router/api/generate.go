package api

import (
	"AI/Key"
	"AI/request"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Generate(f *gin.Context) {
	text := f.Param("text")
	style := f.Param("style")

	var c map[string]interface{}
	//发送请求获取token
	url := "https://wenxin.baidu.com/moduleApi/portal/api/oauth/token?grant_type=client_credentials&client_id=" + Key.Key.ApiKey + "&client_secret=" + Key.Key.SecreatKey
	c, err := request.POST(url, nil, "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println(err)
	}
	token := c["data"].(string)
	fmt.Println(c)

	//发送请求获取ID
	url = "https://wenxin.baidu.com/moduleApi/portal/api/rest/1.0/ernievilg/v1/txt2img?access_token=$" + token
	payload := map[string]string{
		"text" : text,
		"style" : style,
	}
	if err != nil {
		fmt.Println(err)
	}
	c, err = request.POST(url, payload, "")
	if err != nil {
		fmt.Println(err)
	}
	taskId := strconv.Itoa(int((c["data"].(map[string]interface{}))["taskId"].(float64)))
	fmt.Println(c)

	//获取图像
	for {
		url = "https://wenxin.baidu.com/moduleApi/portal/api/rest/1.0/ernievilg/v1/getImg?access_token=" + token
		payload = map[string]string{
			"taskId" : taskId,
		}
		if err != nil {
			fmt.Println(err)
		}
		c, err = request.POST(url, payload, "")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(c)
		if c["data"].(map[string]interface{})["waiting"].(string) == "0" {
			break
		}

		//如果需要等待则等待40s
		time.Sleep(time.Duration(40)*time.Second)
	}

	//下载图像到本地
    imgUrl := c["data"].(map[string]interface{})["img"].(string)
	println(imgUrl)
	request.GetImg(imgUrl, "img", "./")
	f.File("./img.png")
}