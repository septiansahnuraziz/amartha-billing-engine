package cmd

import (
	"amartha-billing-engine/cacher"
	"amartha-billing-engine/config"
	"amartha-billing-engine/docs"
	"amartha-billing-engine/internal/controller/http"
	"amartha-billing-engine/internal/database"
	"amartha-billing-engine/internal/database/transactioner"
	//"amartha-billing-engine/internal/job"
	"amartha-billing-engine/utils"
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	httpSrv "net/http"
	"os"
	"os/signal"
	"strings"
	//"syscall"
	"time"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   server,
}

func init() {
	RootCmd.AddCommand(runServer)
}

func server(cmd *cobra.Command, args []string) {
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	db, err := database.InitializePostgresConnection()
	if err != nil {
		logrus.Fatal("err initialize db")
	}

	postgresDB, err := database.PostgreSQL.DB()
	defer utils.WrapCloser(postgresDB.Close)

	cacheManager := cacher.ConstructCacheManager()

	if config.EnableCaching() {
		redisDB, err := database.InitializeRedigoRedisConnectionPool(config.RedisCacheHost(), redisOptions)
		continueOrFatal(err)
		defer utils.WrapCloser(redisDB.Close)

		cacheManager.SetConnectionPool(redisDB)
	}

	cacheManager.SetDisableCaching(!config.EnableCaching())

	//redisOpt, err := asynq.ParseRedisURI(config.RedisWorkerHost())
	//continueOrFatal(err)

	app := gin.Default()

	// get default url request
	app.UseRawPath = true
	app.UnescapePathValues = true
	app.RemoveExtraSlash = true

	// cors configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization", "Device-Id", "Device", "Source", "Device-Platform", "campaign-type", "campaign-name", "sms-client", "signature")
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "POST", "GET", "DELETE")

	app.Use(cors.New(corsConfig))

	//httpClient := httpclient.NewHTTPConnection(config.DefaultHTTPOptions())
	//taskQueue := job.ConstructQueue(redisOpt, config.WorkerNamespace())
	gormTransactioner := transactioner.NewGormTransactioner(db)

	loanService := InitLoan(db, gormTransactioner)

	http.NewRoute(
		&app.RouterGroup,
		loanService,
	)

	initSwaggerDocs(&app.RouterGroup)

	httpServer := &httpSrv.Server{
		Addr:    ":" + config.HTTPPort(),
		Handler: app,
	}

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		for {
			select {
			case <-sigCh:
				gracefulShutdown(httpServer)
				quitCh <- true
			case e := <-errCh:
				log.Error(e)
				gracefulShutdown(httpServer)
				quitCh <- true
			}
		}
	}()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, httpSrv.ErrServerClosed) {
			errCh <- err
		}
	}()
	<-quitCh

	//taskQueue.Stop()

	log.Info("Server exiting 🔴")
}

func gracefulShutdown(httpServer *httpSrv.Server) {
	database.StopTickerCh <- true

	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown 🔴: ", err)
		}
	}
}

func initSwaggerDocs(app *gin.RouterGroup) {
	swaggerEndpoint := config.SwaggerEndpoint()
	swaggerSchemes := []string{"https"}

	if config.EnvironmentMode() == "local" {
		swaggerSchemes = []string{"http"}
	}

	// swagger configuration
	docs.SwaggerInfo.Title = config.AppName()
	docs.SwaggerInfo.Description = "POS-B2B API"
	docs.SwaggerInfo.Version = config.AppVersion()
	docs.SwaggerInfo.Host = swaggerEndpoint
	docs.SwaggerInfo.Schemes = swaggerSchemes

	swagConfig := &ginSwagger.Config{
		URL: swaggerEndpoint + "/docs/swagger/doc.json",
	}

	// swagger endpoint with authentication
	swaggerDocs := app.Group("/docs", gin.BasicAuth(gin.Accounts{config.SwaggerUsername(): config.SwaggerPassword()}))
	{
		swaggerDocs.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swagConfig, swaggerFiles.Handler))
	}
}

func continueOrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
