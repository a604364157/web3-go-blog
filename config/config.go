package config

import (
	"path/filepath"
	"runtime"
	"time"
)

/*
这里是一些项目基本配置
可以采用配置文件的方式来管理这些配置
这里懒得写，就直接写死
*/
var (
	JWTSecret = []byte("secret")
	TokenTTl  = time.Hour * 2
	Port      = "8080"
	DBPath    = filepath.Join(getRootDir(), "blog.db")
)

func getRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}
