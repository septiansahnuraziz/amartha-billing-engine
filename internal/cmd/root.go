package cmd

import (
	"amartha-billing-engine/config"
	"amartha-billing-engine/internal/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var redisOptions *database.RedisConnectionPoolOptions

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-pos-b2b-service",
	Short: "go pos b2b service console",
	Long:  `This is go pos b2b service console`,
}

func init() {
	config.LoadConfig()

	redisOptions = &database.RedisConnectionPoolOptions{
		DialTimeout:     config.RedisDialTimeout(),
		ReadTimeout:     config.RedisReadTimeout(),
		WriteTimeout:    config.RedisWriteTimeout(),
		IdleCount:       config.RedisMaxIdleConn(),
		PoolSize:        config.RedisMaxActiveConn(),
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 1 * time.Minute,
	}

	log.Info("Environment: ", config.EnvironmentMode())
}

// Execute :nodoc:
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
