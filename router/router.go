package router

import (
	"file-manager/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("./static"))
	r.POST("/mkdir", controller.CreateDir)
	r.GET("/downloadtext", controller.DownloadText)
	r.GET("/count", controller.Count)
	r.GET("/id", controller.GetId)
	r.GET("/remove", controller.Delete)
	r.GET("/share", controller.Share)

	r.GET("/getFileDetail", controller.GetFileContent)
	r.POST("/upload", controller.Upload)
	r.POST("/uploadfile", controller.UploadFile)
	r.GET("/cloud", controller.CloudListPage)
	r.GET("/img", controller.CloudListImg)
	r.GET("/pic", controller.Imgpic)
	r.GET("/picdownload", controller.PicDownload)
	r.GET("/piccopy", controller.PicCopy)
	r.GET("/video", controller.CloudListVideo)
	r.GET("/doc", controller.CloudListText)
	r.GET("/music", controller.CloudListMusic)

	r.GET("/download/:filename", controller.CloudDownload)
	r.GET("/delect/:filename", controller.CloudDelete)
	r.POST("/cloudupload", controller.CloudUpload)
	r.GET("/share/:filename", controller.CloudShare)
	r.GET("/search", controller.Search)
	r.GET("/searchimg", controller.SearchImg)
	r.GET("/searchmusic", controller.SearchMusic)
	r.GET("/searchvideo", controller.SearchVideo)
	r.GET("/searchdoc", controller.SearchDoc)
	r.GET("/look/:filename", controller.CloudLook)
	r.GET("/transfer", controller.Transfer)
	r.Use(controller.FileListPage)
	return r
}
