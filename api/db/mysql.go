package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/toolkits/pkg/file"
	"os"
	"xorm.io/xorm"
)
const ConfPath string = "/Users/tglmm/github_project/Vshared/api/etc/config.yml"

type MysqlConf struct{
	Address string `yaml:"address"`
	Port int `yaml:"post"`
	User string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
}

var Mcf map[string]MysqlConf
var MysqlEng *xorm.Engine

func init(){
	SetMysqlEng()
}

// 设置数据库全局引擎
func SetMysqlEng()error{
	err := ParseFile(ConfPath)
	if err != nil{
		fmt.Println("prase config file error:",err)
		os.Exit(1)
	}
	db := Mcf["mysql"]
	MysqlEng , err = xorm.NewEngine("mysql" , fmt.Sprintf("%s:%s@/%s?charset=utf8" , db.User,db.Password , db.Database))
	if err != nil{
		fmt.Println("xorm.NewEngine error:",err)
	}
	return nil
}

// 解析配置文件
func ParseFile(path string)error{
	if !file.IsExist(path){
		//
		fmt.Println(path," not found")
		os.Exit(1)
	}
	err := file.ReadYaml(path ,&Mcf)
	if err != nil{
		return err
	}
	return nil
}