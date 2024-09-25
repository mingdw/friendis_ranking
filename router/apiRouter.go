package router

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/config/variable"
	"friends_ranking/http/authorization"
	"friends_ranking/http/validator/factory"
	"friends_ranking/utils/ginRelease"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	if variable.YamlConfig.GetBool("AppDebug") == false {
		//1.gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		//gin.DisableConsoleColor()
		//f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)

		//【生产模式】
		// 根据 gin 官方的说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		// 如果部署到生产环境，请使用以下模式：
		// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
		// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
		// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
		router = ginRelease.ReleaseRouter()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	router.LoadHTMLGlob("template/**/*")
	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/static", "./static") //  定义静态资源路由与实际目录映射关系
	//router.StaticFile("/abcd", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/admin/admin.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "admin",
		})
	})

	//模拟两组路由，一组登陆和注册，不需要校验token，一组受保护的api，需要检验token
	//  创建一个门户类接口路由组
	vApi := router.Group("/admin/")
	{
		// 模拟一个首页路由
		home := vApi.Group("index/")
		{
			home.POST("login", factory.Create(globalConst.ValidatorPrefix+"Login"))
			home.POST("regist", factory.Create(globalConst.ValidatorPrefix+"Regist"))
		}
		user := vApi.Group("user/")
		user.Use(authorization.CheckTokenAuth())
		{
			user.GET("query/:{userId}", factory.Create(globalConst.ValidatorPrefix+"Query"))
			user.POST("delete", factory.Create(globalConst.ValidatorPrefix+"Delete"))
			user.POST("add", factory.Create(globalConst.ValidatorPrefix+"Add"))
		}
	}
	return router

}
