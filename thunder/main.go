package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"thunder/FILEHelper"

	"gopkg.in/alecthomas/kingpin.v2"
)

//func

func Cmder_init(path string, force bool) {
	path=FILEHelper.F_getRealAdbPath(path)
	if path ==``{
		fmt.Println(`没有正确的文件路径 :(`)
		return
	}

	fmt.Println("path",path,"force", force,FILEHelper.F_FileExist(path))
	// 判断文件是否存在
	if FILEHelper.F_FileExist(path){
		if !FILEHelper.F_FilePermission(path){
			fmt.Println(`你权限不足哦哦`)
		}
		if force{
			fmt.Println("###", FILEHelper.F_FormatPath(path))
			if FILEHelper.F_isDir(path)&&
				FILEHelper.F_FileExist( filepath.Join(path+`thunder.config.json`)){
				os.RemoveAll(path)
			}else{
				fmt.Println(`只能强制创建thunder，并且我是通过thunder.config.json文件来判断的`)
			}
		}else{
			fmt.Println(`目录存在该文件或者文件夹`)
		}
	}
	// 创建文件夹
	os.MkdirAll(path,0766)


	//	下载基本配置文件

	


	//	创建配置文件
	if f,e:=os.Create( filepath.Join(path  ,`thunder.config.json`)); e==nil{
		f.WriteString(`{
			siteName:"一个 thunder 站点",
			siteSubName:"One thunder",
			url:""
		}`)
	}
	about := filepath.Join(path,`about`)
	os.Mkdir(about,0766)
	if f,e:=os.Create( filepath.Join(about  ,`about.md`)); e==nil{
		f.WriteString(`thunder 静态文件生成器`)
	}



}
func Cmder_build(path string)               {}
func Cmder_add(name string, isSpecial bool) {}
func Cmder_up(url string)                   {}

func main() {
	app := kingpin.New("thunder", "静态博客生产工具")
	app.Version("0.0.1")
	var (
		cmd_init  = app.Command("init", "初始化项目,如果目录已经存在者不会有任何操作。但是可以使用--force=true强制初始化")
		cmd_build = app.Command("build", "编译项目")
		cmd_add   = app.Command("add", "添加文件")
		cmd_up    = app.Command("up", "初始化项目")
	)
	var (

		cmd_init_path  = cmd_init.Flag("path", "路径").Default(".").String()
		cmd_init_force = cmd_init.Flag("force", "强制初始").Default("false").Bool()

		cmd_build_path = cmd_build.Arg("path", "路径,会自动编译文件夹或者文件").Default(".").String()

		cmd_add_isSpecial = cmd_init.Flag("isSpecial", "是否是专题 如果专题的文件格式是<.../文章标题/index.md>").Default("false").Bool()
		cmd_add__name     = cmd_add.Arg("name", "添加文件<分类名称/子分类/子子分类/文章标题/(文章标题.md|index.md> ，第一个是模块的名称, 文章标题不可以是index,robots,favicon").String()

		cmd_up_url = cmd_up.Arg("url", `协议地址 协议类型https`).String()
	)

	c, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parse cli args error: %s\n", err)
	}
	switch c {
	case "init":
		Cmder_init(*cmd_init_path, *cmd_init_force)
		break
	case "build":
		Cmder_build(*cmd_build_path)
		break

	case "add":
		Cmder_add(*cmd_add__name, *cmd_add_isSpecial)
		break

	case "up":
		Cmder_up(*cmd_up_url)
		break
	default:
		fmt.Println("未知参数")
		break
	}
}
