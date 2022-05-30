package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"file-manager/config"
	"file-manager/model"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type FileType uint32

const (
	FileTypeUnknown FileType = 0
	FileTypeImage   FileType = 2
	FileTypeVideo   FileType = 3
	FileTypeTxt     FileType = 1
	FileTypeRar     FileType = 5
	FileTypeMusic   FileType = 4
	FileTypeExe     FileType = 6
	FileTypeHtml    FileType = 7
)

// 读取文件夹
func ReadDir(dir string) ([]*model.FileModel, error) {
	fullPath := filepath.Join(config.RootPath, dir)
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		panic("文件夹错误")
	}
	var fileSlice []*model.FileModel
	for _, fileInfo := range fileInfos {
		isImg := false
		isVideo := false
		isMusic := false
		isExe := false
		israr := false
		isTxt := false
		isHtml :=false
		if !fileInfo.IsDir() {
			fileType := CheckFileType(filepath.Join(fullPath, fileInfo.Name()))
			if fileType == 1 {
				isTxt = true
			} else if fileType == 2 {
				isImg = true
			} else if fileType == 3 {
				isVideo = true
			} else if fileType == 4 {
				isMusic = true
			} else if fileType == 5 {
				israr = true
			} else if fileType == 6 {
				isExe = true
			}else if fileType == 7 {
			    isHtml = true
			}
		}
		file := &model.FileModel{
			filepath.Join(dir, fileInfo.Name()),
			fileInfo.Name(),
			fileInfo.Size(),
			FormatFileSize(fileInfo.Size()),
			fileInfo.Mode(),
			fileInfo.ModTime(),
			FormatTime(fileInfo.ModTime()),
			fileInfo.IsDir(),
			isImg,
			isVideo,
			isMusic,
			isExe,
			israr,
			isTxt,
			isHtml,
		}
		fileSlice = append(fileSlice, file)
	}
	return fileSlice, nil
}

func CheckFileType(filePath string) FileType {

	if GetFileTypeInt(filePath) == 0 {
		return FileTypeUnknown
	}
	if GetFileTypeInt(filePath) == 1 {
		return FileTypeTxt
	}
	if GetFileTypeInt(filePath) == 2 {
		return FileTypeImage
	}
	if GetFileTypeInt(filePath) == 3 {
		return FileTypeVideo
	}
	if GetFileTypeInt(filePath) == 4 {
		return FileTypeMusic
	}
	if GetFileTypeInt(filePath) == 5 {
		return FileTypeRar
	}
	if GetFileTypeInt(filePath) == 6 {
		return FileTypeExe
	}
		if GetFileTypeInt(filePath) == 7 {
		return FileTypeHtml
	}
	return FileTypeUnknown
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FormatFileSize(b int64) string {
	fb := float64(b)
	if fb < 1024 {
		return fmt.Sprintf("%0.2f B", fb)
	}
	kb := fb / 1024
	if kb < 1024 {
		return fmt.Sprintf("%0.2f KB", kb)
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%0.2f MB", mb)
	}
	gb := mb / 1024
	return fmt.Sprintf("%0.2f GB", gb)
}

func FormatTime(t time.Time) string {
	// year, month, day := t.Date()
	return t.Format("2006/01/01 15:04:05")
}

//SHA256生成哈希值
func GetSHA256HashCode(file *os.File) string {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	_, _ = io.Copy(hash, file)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

}

// md5加密
func EncodeMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//判断文件后缀获取类型id
func GetFileTypeInt(filePath string) FileType {
	filePath = path.Ext(filePath)
	if filePath == ".txt" {
		return 1
	}
	if filePath == ".jpg" || filePath == ".png" || filePath == ".gif" || filePath == ".jpeg" {
		return 2
	}
	if filePath == ".mp4" || filePath == ".avi" || filePath == ".mov" {
		return 3
	}
	if filePath == ".mp3" || filePath == ".wav" || filePath ==".m4a"{
		return 4
	}
	if filePath == ".rar" || filePath == ".tar" || filePath == ".zip" {
		return 5
	}
	if filePath == ".exe" {
		return 6
	}
	if filePath == ".html"{
	    return 7
	}
	return 0
}

//将body中=号格式的字符串转为map
func ConvertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}
