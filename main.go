package main

import (
	"cat-slave/router/middleware"
	"errors"
	"net/http"
	"time"

	"cat-slave/config"
	"cat-slave/model"
	"cat-slave/router"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// 设置初始化参数 可以通过-h 查看
	// -c config.yaml
	// 方法返回第三个参数, 如果没有传入, 则为""
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	//files, err := file.GetAllFiles("./html")
	//if err != nil {
	//	panic(err)
	//}
	//for file := range files {
	//	fmt.Println(file)
	//}



	// 必须
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine without default config
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middleware.RequestId(),
		middleware.Logging(),
	)

	model.DB.Init()
	defer model.DB.Close()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("启动失败", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// 测试是否启动
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// 1s后重试
		time.Sleep(time.Second)
	}
	return errors.New("无法成功调用/sd/health接口")
}
