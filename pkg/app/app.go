package app

import (
	v1 "rabbitmqtest/pkg/app/controllers/v1"
	"rabbitmqtest/pkg/domain"
	"rabbitmqtest/pkg/infra/config"
	"rabbitmqtest/pkg/infra/insights"
	"rabbitmqtest/pkg/infra/logger"
	serviceprovider "rabbitmqtest/pkg/infra/serviceProvider"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
	// Registering singletons
	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewInsightsConfig),
		fx.Provide(insights.NewInsightsCore),
		fx.Provide(serviceprovider.NewServiceProviderFactory),
		fx.Provide(domain.NewMqttService),
		fx.Provide(v1.NewMqttController),
		fx.Invoke(StartService),
	)
	app.Run()

	// Invoke cleanups
}

func StartService(
	provFactory *serviceprovider.ServiceProviderFactory,
	lgr *zap.Logger,
	v1mcast *v1.MqttController,
) {

	v1mcast.Csub("Notification")

	// // - Setting up logger

	// router := gin.New()
	// // gin.SetMode(gin.ReleaseMode)
	// router.SetTrustedProxies(nil)

	// // - Swagger

	// // - Setting up middlewares
	// router.Use(gingorr.RootRecoveryMiddleware(lgr))
	// router.Use(trex.TxContextMiddleware(provFactory))
	// router.Use(trex.RequestTracerMiddleware(func(
	// 	context interface{},
	// 	method,
	// 	path,
	// 	query,
	// 	agent,
	// 	ip string,
	// 	status,
	// 	bytes int,
	// 	start,
	// 	end time.Time) {
	// 	sp := context.(*serviceprovider.ServiceProvider)
	// 	latency := end.Sub(start)

	// 	sp.GetTracer().TraceRequest(
	// 		// this true is being ignored :)
	// 		true,
	// 		method,
	// 		path,
	// 		query,
	// 		status,
	// 		bytes,
	// 		ip,
	// 		agent,
	// 		start,
	// 		end,
	// 		map[string]string{},
	// 	)
	// 	sp.GetLogger().Info(
	// 		"Request",
	// 		zap.Int("status", status),
	// 		zap.String("method", method),
	// 		zap.String("path", path),
	// 		zap.String("query", query),
	// 		zap.String("ip", ip),
	// 		zap.String("userAgent", agent),
	// 		zap.Time("mvts", end),
	// 		zap.String("pmvts", end.Format("2006-01-02T15:04:05-0700")),
	// 		zap.Duration("latency", latency),
	// 		zap.String("pLatency", latency.String()),
	// 	)
	// },
	// ))
	// router.Use(gingorr.RecoveryMiddleware("tx-context", lgr))
	// router.GET(
	// 	"/swagger/*any",
	// 	ginSwagger.WrapHandler(swaggerFiles.Handler),
	// )
	// router.Use(gingorr.ErrorHandlerMiddleware("tx-context"))

	// // - Setting up routes
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"status": "alive",
	// 	})
	// })

	// v1mcast

	// router.Run(":8080")
}
