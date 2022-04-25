package main

import (
    "context"
    "database/sql"
    "log"

    "cmdb-app-mysql/controllers"
    "cmdb-app-mysql/services"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    _ "github.com/go-sql-driver/mysql"
)

var (
    mysqlClient *sql.DB
    redisClient *redis.Client
    ctx         context.Context

    err    error
    server *gin.Engine

    userService    services.UserService
    userController controllers.UserController
)

func init() {
    // 数据库
    mysqlClient, err = sql.Open("mysql", "cmdb:cmdb@tcp(127.0.0.1:3306)/cmdb?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
    if err != nil {
        log.Fatal(err.Error())
    }

    err = mysqlClient.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("mysql connection established")

    // 缓存
    opt, err := redis.ParseURL("redis://:@localhost:6379/0")
    if err != nil {
        log.Fatal(err.Error())
    }

    ctx = context.Background()
    redisClient := redis.NewClient(opt)
    if _, err := redisClient.Ping(ctx).Result(); err != nil {
        log.Fatal(err.Error())
    }
    log.Println("redis connection established")

    // 业务
    userService = services.NewUserService(mysqlClient, redisClient, ctx)
    userController = controllers.New(userService)

    // gin
    // gin.SetMode(gin.ReleaseMode)
    // server = gin.New()
    server = gin.Default()
    server.SetTrustedProxies(nil)
}

func main() {
    defer mysqlClient.Close()

    baseUrl := server.Group("/v1")
    userController.RegisterUserRoutes(baseUrl)

    log.Fatal(server.Run(":9000"))
}
