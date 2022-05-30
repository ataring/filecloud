# filecloud
更适合个人使用的云存储云盘

配置config.go
```golang
var (
	RootPath string = "../"
	//RootPath      string = "/root/crazyball" // 根路径
	ServerPort    int    = 80        //端口号
	TempDirName   string = "test"    // 临时文件夹
	StaticDirName string = "./"      // 静态资源地址
	UploadDirName string = "uploads" // 上传文件夹
	StaticHost    string = "127.0.0.1/"

	Endpoint        string = "<>"
	AccessKeyId     string = "<>"
	AccessKeySecret string = "<>"
	BucketName      string = "<>"
)

```

```golang
使用go modules管理项目依赖，配置configgo完成后go run .运行

```
![This is an image](../doc/示例.png)
