package controller

import (
	"file-manager/config"
	"file-manager/model"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {

	var (
		filename []string
		filesize int64
	)
	fileinfo := make([]model.FileModel, 0)
	files := model.FileModel{}
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		HandleError(err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := config.BucketName
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
		os.Exit(-1)
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	search := c.Query("search")

	for {
		lsRes, err := bucket.ListObjects(oss.Prefix(search))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		for _, object := range lsRes.Objects {

			filename = append(filename, object.Key)
			files.Name = object.Key

			if path.Ext(object.Key) == ".txt" {
				files.IsTxt = true

			} else {
				files.IsImg = false
				files.IsExe = false
				files.IsVideo = false
				files.IsMusic = false
				files.Israr = false
			}
			if path.Ext(object.Key) == ".doc" {
				files.IsTxt = true

			}
			if path.Ext(object.Key) == ".png" {

				files.IsImg = true

			}
			if path.Ext(object.Key) == ".jpg" {

				files.IsImg = true

			}
			if path.Ext(object.Key) == ".jpeg" {

				files.IsImg = true

			}
			if path.Ext(object.Key) == ".mp3" {

				files.IsMusic = true
			}

			fileinfo = append(fileinfo, files)
			filesize += object.Size
		}

		break

	}

	c.HTML(http.StatusOK, "cloud.html", gin.H{
		"filename": filename,
		"filesize": filesize / 1024 / 1024,
		"onsize":   1,
		"oncir":    430,
		"code":     200,
		"type":     fileinfo,
	})
	return
}

//搜索图片
func SearchImg(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
	)

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := config.BucketName
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	search := c.Query("search")
	fmt.Println(search)
	lsRes, err := bucket.ListObjects(oss.Prefix(search))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 打印列举结果。默认情况下，一次返回100条记录。
	//fmt.Println("Objects:", lsRes.Objects)
	for _, object := range lsRes.Objects {

		if path.Ext(object.Key) == ".png" {
			filename = append(filename, object.Key)

		}
		if path.Ext(object.Key) == ".jpg" {
			filename = append(filename, object.Key)

		}
		if path.Ext(object.Key) == ".jpeg" {
			filename = append(filename, object.Key)

		}
		filesize += object.Size
	}

	c.HTML(http.StatusOK, "img.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"code":     200,
	})
}

//搜索文档
func SearchDoc(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
	)

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := config.BucketName
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	search := c.Query("search")
	fmt.Println(search)
	lsRes, err := bucket.ListObjects(oss.Prefix(search))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 打印列举结果。默认情况下，一次返回100条记录。
	//fmt.Println("Objects:", lsRes.Objects)
	for _, object := range lsRes.Objects {

		if path.Ext(object.Key) == ".txt" {
			filename = append(filename, object.Key)

		}
		if path.Ext(object.Key) == ".docx" {
			filename = append(filename, object.Key)

		}
		if path.Ext(object.Key) == ".doc" {
			filename = append(filename, object.Key)

		}
		filesize += object.Size
	}

	c.HTML(http.StatusOK, "doc.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"code":     200,
	})
}

//搜索视频
func SearchVideo(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
	)

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := config.BucketName
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	search := c.Query("search")
	fmt.Println(search)
	lsRes, err := bucket.ListObjects(oss.Prefix(search))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 打印列举结果。默认情况下，一次返回100条记录。
	//fmt.Println("Objects:", lsRes.Objects)
	for _, object := range lsRes.Objects {

		if path.Ext(object.Key) == ".mp4" {
			filename = append(filename, object.Key)

		}
		if path.Ext(object.Key) == ".wav" {
			filename = append(filename, object.Key)

		}
		filesize += object.Size
	}

	c.HTML(http.StatusOK, "video.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"code":     200,
	})
}

//搜索音乐
func SearchMusic(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
	)

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := config.BucketName
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	search := c.Query("search")
	fmt.Println(search)
	lsRes, err := bucket.ListObjects(oss.Prefix(search))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 打印列举结果。默认情况下，一次返回100条记录。
	//fmt.Println("Objects:", lsRes.Objects)
	for _, object := range lsRes.Objects {

		if path.Ext(object.Key) == ".mp3" {
			filename = append(filename, object.Key)

		}

		filesize += object.Size
	}

	c.HTML(http.StatusOK, "music.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"code":     200,
	})
}
