package controller

import (
	"bytes"
	"file-manager/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func DownloadText(c *gin.Context) {

	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写Bucket名称，例如examplebucket。
	bucket, err := client.Bucket("ataring")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 下载文件到本地文件，并保存到指定的本地路径中。如果指定的本地文件存在会覆盖，不存在则新建。
	// 如果未指定本地路径，则下载后的文件默认保存到示例程序所属项目对应本地路径中。
	// 依次填写Object完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径(例如D:\\localpath\\examplefile.txt)。Object完整路径中不能包含Bucket名称。
	localpath := "E:\\test\\" + getId
	err = bucket.GetObjectToFile(getId, localpath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	c.HTML(http.StatusOK, "success.html", gin.H{
		"rqs": "下载成功！",
	})

}

//下载
func CloudDownload(c *gin.Context) {
	filename := getId
	url := "https://ataring.oss-cn-shanghai.aliyuncs.com/" + filename
	//fmt.Println(url)
	c.Redirect(http.StatusFound, url)
}

func CloudLook(c *gin.Context) {
	filename := c.Param("filename")
	url := "https://ataring.oss-cn-shanghai.aliyuncs.com/" + filename
	//fmt.Println(url)
	c.Redirect(http.StatusFound, url)
}

//删除
func CloudDelete(c *gin.Context) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	filename := c.Param("filename")
	var url string
	suffix := path.Ext(filename)
	if suffix == ".png" {
		url = "/img"
	} else if suffix == ".jpg" || suffix == ".jpeg" {
		url = "/img"
	} else if suffix == ".txt" || suffix == ".doc" {
		url = "/doc"
	} else if suffix == ".mp3" {
		url = "/music"
	} else if suffix == ".mp4" {
		url = "/video"
	} else {
		url = "/cloud"
	}

	//url :=c.Param("")
	bucketName := config.BucketName
	// objectName表示删除OSS文件时需要指定包含文件后缀，不包含Bucket名称在内的完整路径，例如exampledir/exampleobject.txt。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	objectName := filename

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 删除单个文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	c.Redirect(http.StatusFound, url)
}

//上传文件
func CloudUpload(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    500,
				"message": err,
			})
		}
	}()
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
	form, _ := c.MultipartForm()
	basePath := c.PostForm("basePath")
	fmt.Println(basePath)
	if basePath == "" {
		panic("缺少上传Path")
	}
	files := form.File["files"]
	if len(files) < 1 {
		panic("缺少上传文件")
	}
	for _, file := range files {
		err = bucket.PutObject(file.Filename, bytes.NewReader([]byte("file")))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		if err != nil {
			panic(fmt.Sprintf("上传文件%s错误", file.Filename))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  200,
	})

}

//二维码分享文件
func CloudShare(c *gin.Context) {
	filename := c.Param("filename")
	url := "https://ataring.oss-cn-shanghai.aliyuncs.com/" + filename
	//fmt.Println(url)
	f, err := qrcode.Encode(url, qrcode.Highest, 300)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.Writer.WriteString(string(f))
    
	// _ = qrcode.WriteFile(url, qrcode.Low, 500, "QR"+filename+".png")
}
