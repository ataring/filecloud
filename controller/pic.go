package controller

import (
	"file-manager/config"
	"net/http"
	"os"
	"path"

	"github.com/atotto/clipboard"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

func PicDownload(c *gin.Context) {
	filename := c.Query("id")
	url := "https://ataring.oss-cn-shanghai.aliyuncs.com/" + filename
	//fmt.Println(url)
	c.Redirect(http.StatusFound, url)
}

func PicCopy(c *gin.Context) {
	filename := c.Query("id")
	url := "https://ataring.oss-cn-shanghai.aliyuncs.com/" + filename
	clipboard.WriteAll(url)
}

//读取照片

func Imgpic(c *gin.Context) {
	var (
		filename []string
	)

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
	// 列举指定存储空间下所有文件。
	startAfter := ""
	continueToken := ""

	for {
		lsRes, err := bucket.ListObjectsV2(oss.StartAfter(startAfter), oss.ContinuationToken(continueToken))
		if err != nil {
			HandleError(err)
			os.Exit(-1)
		}
		for _, object := range lsRes.Objects {
			//if object.Type

			if path.Ext(object.Key) == ".png" {
				filename = append(filename, object.Key)

			}
			if path.Ext(object.Key) == ".jpg" {
				filename = append(filename, object.Key)

			}
			if path.Ext(object.Key) == ".jpeg" {
				filename = append(filename, object.Key)

			}
		}
		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	c.HTML(http.StatusOK, "pic.html", gin.H{
		"filename": filename,
		"code":     200,
	})
	return
}
