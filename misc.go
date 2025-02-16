package common

import (
	"net/http"
	"os"
)

func DetectMIME(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 读取文件前512字节用于判断文件类型
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// 检测文件类型
	fileType := http.DetectContentType(buffer)
	return fileType, nil
}
