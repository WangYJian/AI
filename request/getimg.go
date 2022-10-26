package request

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetImg(url string, filename string, path string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)
	file, err := os.Create(path + filename + ".png")
	if err != nil {
		panic(err)
	}
	
	// 获得文件的writer对象
	writers := bufio.NewWriter(file)
	written, _ := io.Copy(writers, reader)
	fmt.Printf("Total length: %d", written)
}