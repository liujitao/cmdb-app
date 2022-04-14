package main

import (
    "context"
    "database/sql"
    "log"

    "cmdb-app/controllers"
    "cmdb-app/services"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

var (
    mysqlClient *sql.DB

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

    log.Printf("mysql connection established")

    // 业务
    ctx := context.Background()
    userService = services.NewUserService(mysqlClient, ctx)
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
