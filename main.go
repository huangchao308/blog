package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/huangchao308/blog/global"
	"github.com/huangchao308/blog/internal/model"
	"github.com/huangchao308/blog/internal/routers"
	"github.com/huangchao308/blog/pkg/logger"
	"github.com/huangchao308/blog/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setting err: %+v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init db engine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init logger err: %v", err)
	}
	global.Logger.Infof("init success")
}

func setupSetting() error {
	// 初始化配置文件
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	dbEngine, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DBEngine = dbEngine
	return nil
}

func setupLogger() error {
	fmt.Printf("%+v\n", global.LogSetting)
	zapLogger, err := logger.NewLogger(global.LogSetting)
	if err != nil {
		return err
	}
	global.Logger = zapLogger
	return nil
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	router := routers.NewRouter()

	// router.Run() 其实也是调用了 net/http.ListenAndServe, 但是它只帮我们设置了 Addr 和 Handler，如果想要定义其他属性，就需要自己定义 http.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
