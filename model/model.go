package model

import (
	"os"
	"time"
)





type FileModel struct {
	Path       string
	Name       string
	Size       int64
	SizeStr    string
	Mode       os.FileMode
	ModTime    time.Time
	ModTimeStr string
	IsDir      bool
	IsImg      bool
	IsVideo    bool
	IsMusic    bool
	IsExe      bool
	Israr      bool
	IsTxt      bool
	IsHtml  bool
}

type FileCount struct {
	Filecount  int
	Imgcount   int
	Doccount   int
	Videocount int
	Musiccount int
}
