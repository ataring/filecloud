package controller

import (
	"file-manager/config"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"

	"github.com/go-echarts/go-echarts/charts"
)

var nameItems = []string{"图片", "文档", "音乐", "视频", "应用程序", "未知类型"}
var seed = rand.NewSource(time.Now().UnixNano())

//读取文件
func count() []int {
	var list []int = []int{}
	var (
		img     int
		doc     int
		music   int
		video   int
		exe     int
		unknown int
	)
	img = 1
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

			if path.Ext(object.Key) == ".txt" || path.Ext(object.Key) == ".doc" {
				doc += 1
			} else if path.Ext(object.Key) == ".png" || path.Ext(object.Key) == ".jpg" || path.Ext(object.Key) == ".jpeg" {

				img += 1

			} else if path.Ext(object.Key) == ".mp3" {
				music += 1

			} else if path.Ext(object.Key) == ".mp4" {
				video += 1

			} else if path.Ext(object.Key) == ".exe" {
				exe += 1
			} else {
				unknown += 1
			}

		}
		if lsRes.IsTruncated {
			startAfter = lsRes.StartAfter
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}
	list = append(list, img, doc, music, video, exe, unknown)
	return list
}

func Count(c *gin.Context) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "文件数量"}, charts.ToolboxOpts{Show: true}, charts.InitOpts{Width: "800px", Height: "400px"})

	bar.AddXAxis(nameItems).
		AddYAxis("个数", count(), charts.ColorOpts{"lightblue"})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(c.Writer, f) // Render 可接收多个 io.Writer 接口
}
