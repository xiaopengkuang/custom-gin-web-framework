package main

import (
	"fmt"
	"gin-web/common/web"
	"gin-web/config"
	"gin-web/datadocker"
	"gin-web/engine"
	"gin-web/module/app"
	"gin-web/module/base"
	"gin-web/module/oss"
	"gin-web/module/pc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// 默认处理函数
func defaultFunction(gc *gin.Context) {
	// 浏览器会尝试发送Option 请求
	if strings.EqualFold(gc.Request.Method, "OPTION") {
		gc.JSON(http.StatusOK, nil)
		return
	}

	//request, err := util.BuildRequest(gc)
	//if err != nil {
	//	printJsonResponse(gc, web.FailResp(err.Error(), nil))
	//	return
	//}

	// 执行方法
	response := engine.ExecuteFunction(gc)
	printJsonResponse(gc, response)
}

func printJsonResponse(gc *gin.Context, response web.Response) {
	// 打印执行结果
	gc.JSON(http.StatusOK, response)
}

func init() {
	// 开启数据链接
	err := datadocker.InitDataDocker()
	if err != nil && len(err) > 0 {
		log.Println(err)
		os.Exit(-1)
	}

	// 注册module
	errs := initModules()
	if errs != nil && len(errs) > 0 {
		log.Println(errs)
		os.Exit(-1)
	}
}

func initModules() []error {
	// 注册modudle
	engine.RegisterModule(app.APPModuleName, app.GetModule())
	engine.RegisterModule(base.BaseModuleName, base.GetModule())
	engine.RegisterModule(pc.PCModuleName, pc.GetModule())
	engine.RegisterModule(oss.OSSModuleName, oss.GetModule())

	return engine.GetRegisterError()
}

func main() {

	// gin.DisableConsoleColor()
	//os.Setenv("GIN_MODE", "release")
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// 解决跨域问题
	router.Use(web.Cors())

	// 设置路由
	router.Any("/:module/:service/:operation", defaultFunction)

	s := &http.Server{
		Addr:           config.ServerAddress,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	fmt.Println("server started...")
}
