package upload

import (
	"douyin/pkg/utils"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// GetIntactFileName 获取文件原始名称，并对其进行MD5加密，并拼接后缀，得到完整文件名
func GetIntactFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileName 获取文件名
func GetFileName(name string) string {
	return strings.TrimSuffix(name, GetFileExt(name))
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存地址
func GetSavePath() string {
	return viper.GetString("upload.save_path")
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckFileExt(name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToLower(ext)
	if ext == ".mp4" {
		return true
	}
	return false
}

func CheckMaxSize(f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	if size >= viper.GetInt("upload.image_max_size")*1024*1024 {
		return true
	}
	return false
}
