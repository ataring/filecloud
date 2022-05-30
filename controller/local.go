package controller

import (
	"file-manager/config"
	"file-manager/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var getId string

func GetId(c *gin.Context) {
	getId = c.Query("getId")

}

// 上传多文件【通用接口】
func UploadFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) < 1 {
		c.String(http.StatusBadRequest, "缺少上传文件")
		return
	}
	data := make([]string, 0)

	now := time.Now()
	dateStr := now.Format("2006-01-02")

	uploadDir := path.Join(config.RootPath, config.StaticDirName, config.UploadDirName, dateStr)
	if !util.Exists(uploadDir) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.String(http.StatusBadRequest, "创建文件夹失败")
			return
		}
	}
	for _, file := range files {
		fileName := fmt.Sprintf("%d_", now.Unix()) + file.Filename
		filePath := path.Join(uploadDir, fileName)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传文件%s错误", file.Filename))
			return
		} else {
			data = append(data, path.Join(config.StaticHost, config.UploadDirName, dateStr, fileName))
		}
	}
	c.String(http.StatusOK, "上传成功")
}

//上传文件
func Upload(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    500,
				"message": err,
			})
		}
	}()
	form, _ := c.MultipartForm()
	basePath := c.PostForm("basePath")
	if basePath == "" {
		panic("缺少上传Path")
	}
	files := form.File["files"]
	if len(files) < 1 {
		panic("缺少上传文件")
	}
	for _, file := range files {
		filePath := path.Join(config.RootPath, basePath, file.Filename)

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			panic(fmt.Sprintf("上传文件%s错误", file.Filename))
		}
	}

}

//二维码分享文件
func Share(c *gin.Context) {
	filename := getId
	url := "https://" + config.StaticHost + filename
	//fmt.Println(url)
	f, err := qrcode.Encode(url, qrcode.Highest, 300)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.Writer.WriteString(string(f))

	// _ = qrcode.WriteFile(url, qrcode.Low, 500, "QR"+filename+".png")
}

// 创建文件夹
func CreateDir(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err,
			})
		}
	}()

	basePath := c.PostForm("basePath")
	dirName := c.PostForm("dirName")
	if len(basePath) == 0 || len(dirName) == 0 {
		panic("不能为空")
	}
	dirPath := path.Join(config.RootPath, basePath, dirName)
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		panic("创建文件夹失败: " + err.Error())
	}
	c.Redirect(302, basePath)
}

//转移文件
func Transfer(c *gin.Context) {
	var pathlist []string
	pathlist = strings.Split(getId, "/")
	filePath := getId
	fileName := pathlist[len(pathlist)-1]
	fmt.Println(filePath)
	fullPath := path.Join(config.RootPath, filePath)
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("ataring")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(fileName, fullPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	c.Redirect(http.StatusFound, "/")
}

// 删除文件
func Delete(c *gin.Context) {
	url := c.Param("")
	filePath := getId
	fmt.Println(filePath)

	if len(filePath) == 0 {
		c.String(http.StatusBadRequest, "路径不能为空")
		return
	}
	fullPath := path.Join(config.RootPath, filePath)
	if err := os.Remove(fullPath); err != nil {
		c.String(http.StatusBadRequest, "删除失败: "+err.Error())
		return
	}
	c.Redirect(http.StatusFound, url)
}

// 查看文件内容接口
func GetFileContent(c *gin.Context) {
	relativePath := c.Query("path")
	fullPath := filepath.Join(config.RootPath, relativePath)
	info, err := os.Stat(fullPath)
	if err != nil {
		c.String(http.StatusBadRequest, "file is not exist")
	} else if info.IsDir() {
		c.String(http.StatusBadRequest, "file is dir")
	} else {
		fileType := util.CheckFileType(fullPath)
		if info.Size() > 5*1024*1024 && fileType == util.FileTypeUnknown {
			c.String(http.StatusBadRequest, "file size is larger than 5 mb")
			return
		}
		file, err := os.Open(fullPath)
		if err != nil {
			c.String(http.StatusBadRequest, "open file is error")
			return
		}
		http.ServeContent(c.Writer, c.Request, fullPath, info.ModTime(), file)
	}
}

//读取本地文件
func FileListPage(c *gin.Context) {
	fullpath := filepath.Join(config.RootPath, c.Request.URL.Path)
	info, err := os.Stat(fullpath)
	if err != nil {
		c.String(404, "404 not found")
	} else if !info.IsDir() {
		file, err := os.Open(fullpath)
		if err != nil {
			fmt.Print("open file is error")
			return
		}
		http.ServeContent(c.Writer, c.Request, fullpath, info.ModTime(), file)
	} else {
		res, _ := util.ReadDir(c.Request.URL.Path)
		for _, value := range res {
			val := strings.Trim((strings.Replace(value.Path, "\\", "/", -1)), "/")
			value.Path = val

		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"basePath": c.Request.URL.Path,
			"dirList":  res,
			"code":     200,
		})
	}
}
