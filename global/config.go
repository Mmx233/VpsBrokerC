package global

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/config"
	"log"
	"os"
)

// Config 全局配置
var Config models.Config

func init() {
	//读取配置文件
	if e := config.Load(config.Options{
		Config: &Config,
		Default: &models.Config{
			Settings: models.Settings{
				Port: 575,
			},
			Remote: models.Remote{
				Port: 574,
			},
		},
		FillDefault: true,
	}); e != nil {
		if config.IsNew(e) {
			log.Println("已生成配置文件'Config.json'，请编辑后再次启动")
			os.Exit(1)
		}

		log.Fatalln("配置文件初始化失败：\n", e)
	}

	AuthHeader.Add("Authorization", Config.Remote.AccessKey)
}
