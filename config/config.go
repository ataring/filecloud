package config



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
