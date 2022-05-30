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

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

//读取文件
func CloudListPage(c *gin.Context) {

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

			filename = append(filename, object.Key)
			files.Name = object.Key

			files.IsTxt = false
			files.IsImg = false
			files.IsExe = false
			files.IsVideo = false
			files.IsMusic = false
			files.Israr = false
			if path.Ext(object.Key) == ".txt" || path.Ext(object.Key) == ".doc" {
				files.IsTxt = true
			} else if path.Ext(object.Key) == ".png" || path.Ext(object.Key) == ".jpg" || path.Ext(object.Key) == ".jpeg" {

				files.IsImg = true

			} else if path.Ext(object.Key) == ".mp4" {

				files.IsVideo = true

			} else if path.Ext(object.Key) == ".exe" {

				files.IsExe = true

			} else if path.Ext(object.Key) == ".mp3" {

				files.IsMusic = true
			}

			fileinfo = append(fileinfo, files)
			filesize += object.Size
		}
		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
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

//读取照片

func CloudListImg(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
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
			filesize += object.Size
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
		sizestr = filesize / 1024 / 1024

		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	c.HTML(http.StatusOK, "img.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"onsize":   sizestr / 51200,
		"code":     200,
	})
	return
}

//读取视频

func CloudListVideo(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
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
			filesize += object.Size
			if path.Ext(object.Key) == ".mp4" {
				filename = append(filename, object.Key)

			}

		}
		sizestr = filesize / 1024 / 1024

		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	c.HTML(http.StatusOK, "video.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"onsize":   sizestr / 51200,
		"code":     200,
	})
	return
}

//读取文档

func CloudListText(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
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
			filesize += object.Size
			if path.Ext(object.Key) == ".txt" {
				filename = append(filename, object.Key)

			}
			if path.Ext(object.Key) == ".doc" {
				filename = append(filename, object.Key)

			}

		}
		sizestr = filesize / 1024 / 1024

		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	c.HTML(http.StatusOK, "doc.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"onsize":   sizestr / 51200,
		"code":     200,
	})
	return
}

//读取音乐

func CloudListMusic(c *gin.Context) {
	var (
		filename []string
		filesize int64
		sizestr  int64
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
			filesize += object.Size
			if path.Ext(object.Key) == ".mp3" {
				filename = append(filename, object.Key)

			}

		}
		sizestr = filesize / 1024 / 1024

		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	c.HTML(http.StatusOK, "music.html", gin.H{
		"filename": filename,
		"filesize": sizestr,
		"onsize":   sizestr / 51200,
		"code":     200,
	})
	return
}
