package FILEHelper

import (
	"os"
	"path/filepath"
	"strings"
)

/*
文件路径处理
*/
func F_FormatPath(path string) string {
	return filepath.FromSlash(strings.TrimRight(path,"/\\"))
}

/*
获取文件夹的名称
**/
func F_getRealAdbPath(path string)string{
	var  (
		adsPath string
		err error
	)

	basePath,_:=os.Getwd()
	basePath=F_FormatPath(basePath)
	adsPath ,err =filepath.Abs(filepath.Join(basePath,path))
	if err!=nil{
		return  ``
	}
	return  F_FormatPath(adsPath)
}

/*
判断文件是否有权限
*/
func F_FilePermission(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsPermission(err) {
		return true
	}
	return false
}
/*
判断路径文件是否存在，
没有权限，等等原因返回false，其他情况返回ture
*/
func F_FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

/*
判断路径是不是文件夹
不存在或者其他异常返回false
*/

func F_isDir(path string) (bool) {
	f, err := os.Stat(path)

	if err != nil {
		return false
	}
	if	f.IsDir() {
		return true
	}
	return false
}

