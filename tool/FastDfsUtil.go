package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tedcy/fdfs_client"
)

func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileID, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fileID
}

/**
* 从配置文件读取文件服务器的IP和端口配置
 */

func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //读取一行
		if err != nil {
			return ""
		}
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
	}
}
