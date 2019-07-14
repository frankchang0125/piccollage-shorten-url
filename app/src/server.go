package main

import (
    "fmt"
    "os"
    "net/http"
    "strings"

    "pic-collage.com/shorten_url/controllers"
    "pic-collage.com/shorten_url/services"
    "pic-collage.com/shorten_url/clients/mysql"
    "pic-collage.com/shorten_url/clients/redis"

	"github.com/gorilla/mux"
    "github.com/spf13/viper"
    log "github.com/sirupsen/logrus"
)

func main() {
    viper.SetConfigName("env")

    env := os.Getenv("ENV")
    switch strings.ToUpper(env) {
        case "PROD":
            viper.AddConfigPath("./configs/prod/")
        case "DEV":
        fallthrough
        default:
            viper.AddConfigPath("./configs/dev/")
    }

    err := viper.ReadInConfig()
    if err != nil {
        log.Error("Fail to load config")
        return
    }

    logLevel := viper.GetString("server.loglevel")
    switch strings.ToLower(logLevel) {
    case "trace":
        log.SetLevel(log.TraceLevel)
    case "warn":
        log.SetLevel(log.WarnLevel)
    case "debug":
        log.SetLevel(log.DebugLevel)
    case "info":
        fallthrough
    default:
        log.SetLevel(log.InfoLevel)
    }

    r := mux.NewRouter().StrictSlash(true)

    for _, route := range controllers.Routes {
        path := fmt.Sprintf("/%s", route.Endpoint)
        r.Methods(route.Method).Path(path).HandlerFunc(route.Handler)
    }

    // Init MySQL client.
    db, err := mysql.InitClient()
    if err != nil {
        return
    }

    // Init Redis client.
    redisClient := redis.InitClient()

    // Init services.
    err = services.Init(db, redisClient)
    if err != nil {
        return
    }

    port := viper.GetInt("server.port")
    serverURL := fmt.Sprintf("0.0.0.0:%d", port)

    log.WithField("url", serverURL).Info("Starting server")

    err = http.ListenAndServe(serverURL, r)
    if err != nil {
        log.WithError(err).Error("Start server failed")
    }
}
