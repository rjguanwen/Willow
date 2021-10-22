// @Title  config.go
// @Description  加载配置文件信息，供其他功能使用
// @Author  zhengbin  2021/9/13 16:00
package public

import (
	log "github.com/cihub/seelog"
	"github.com/rjguanwen/willow/utils"
	"github.com/spf13/viper"
	"path"
)

// @Desc 服务器基础配置结构体
type BaseConfig struct {
	Version       	string     `yaml:"version"`
	Port          	string        `yaml:"server.port"`
	MemberTagsFile 	string     `yaml:"dataFilePath.memberTagsFile"`
}

// @title    InitBaseConfig
// @description   根据传入的环境参数，读取配置文件 base.yml，生成 BaseConfig 结构
// @auth      zhengbin  2021/9/13 16:00
// @param     env        string         "环境类型参数"
// @return    baseConfig        BaseConfig         "基础配置结构体"
// @return    err        error         "错误信息"
func InitBaseConfig(env string) (baseConfig BaseConfig, err error) {
	// 配置文件路径
	configFilePath := path.Join(utils.AbsPath("conf"), env)
	configFileName := "base"

	v := viper.New()
	v.SetConfigName(configFileName) //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	v.SetConfigType("yml")
	v.AddConfigPath(configFilePath) //设置配置文件的搜索目录

	err = v.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Error("读取配置文件错误：", err) // 读取配置文件失败致命错误
	}

	// 服务器基础配置
	baseConfig = BaseConfig{
		Version:       v.GetString("version"),
		Port:          v.GetString("server.port"),
		MemberTagsFile:	   v.GetString("dataFilePath.memberTagsFile"),
	}
	return
}
