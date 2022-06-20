package main

import (
    "context"
    "database/sql"
    "log"

    "cmdb-app-mysql/controllers"
    "cmdb-app-mysql/services"

    "github.com/gin-contrib/cors"
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

    roleService    services.RoleService
    roleController controllers.RoleController

    departmentService    services.DepartmentService
    departmentController controllers.DepartmentController

    permissionService    services.PermissionService
    permissionController controllers.PermissionController
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

    // 构造
    userService = services.NewUserService(mysqlClient, redisClient, ctx)
    userController = controllers.NewUserController(userService)

    roleService = services.NewRoleService(mysqlClient, ctx)
    roleController = controllers.NewRoleController(roleService)

    departmentService = services.NewDepartmentService(mysqlClient, ctx)
    departmentController = controllers.NewDepartmentController(departmentService)

    permissionService = services.NewPermissionService(mysqlClient, ctx)
    permissionController = controllers.NewPermissionController(permissionService)

    // gin
    // gin.SetMode(gin.ReleaseMode)
    // server = gin.New()
    server = gin.Default()
    server.SetTrustedProxies(nil)
}

func main() {
    defer mysqlClient.Close()

    // 根路由
    v1 := server.Group("/v1")
    v1.Use(cors.Default())

    // user 路由
    userRoute := v1.Group("/user")
    {
        userRoute.POST("/login", userController.LoginUser)
        userRoute.POST("/logout", userController.LogoutUser)
        userRoute.POST("/refresh", userController.RefreshUser)

        // userRoute.POST("/create", userController.CreateUser)
        // userRoute.GET("/get", userController.GetUser)
        // userRoute.PATCH("/update", userController.UpdateUser)
        // userRoute.DELETE("/delete", userController.DeleteUser)
        // userRoute.GET("/list", userController.GetUserList)
        // userRoute.POST("/change_password", userController.ChangeUserPassword)
    }

    userRoute.Use(userController.AuthMiddleware())
    {
        userRoute.POST("/create", userController.CreateUser)
        userRoute.GET("/get", userController.GetUser)
        userRoute.PATCH("/update", userController.UpdateUser)
        userRoute.DELETE("/delete", userController.DeleteUser)
        userRoute.GET("/list", userController.GetUserList)
        userRoute.POST("/change_password", userController.ChangeUserPassword)
    }

    // role 路由
    roleRoute := v1.Group("/role")
    {
        roleRoute.POST("/create", roleController.CreateRole)
        roleRoute.GET("/get", roleController.GetRole)
        roleRoute.PATCH("/update", roleController.UpdateRole)
        roleRoute.DELETE("/delete", roleController.DeleteRole)
        roleRoute.GET("/list", roleController.GetRoleList)
        roleRoute.GET("/select", roleController.GetRoleOption)
    }

    roleRoute.Use(userController.AuthMiddleware())
    {

    }

    // department 路由
    departmentRoute := v1.Group("/department")
    {
        departmentRoute.POST("/create", departmentController.CreateDepartment)
        departmentRoute.GET("/get", departmentController.GetDepartment)
        departmentRoute.PATCH("/update", departmentController.UpdateDepartment)
        departmentRoute.DELETE("/delete", departmentController.DeleteDepartment)
        departmentRoute.GET("/list", departmentController.GetDepartmentList)
        departmentRoute.GET("/tree", departmentController.GetDepartmentTree)
        departmentRoute.GET("/select", departmentController.GetDepartmentOption)
    }

    departmentRoute.Use(userController.AuthMiddleware())
    {

    }

    // permission 路由
    permissionRoute := v1.Group("/permission")
    {
        permissionRoute.POST("/create", permissionController.CreatePermission)
        permissionRoute.GET("/get", permissionController.GetPermission)
        permissionRoute.PATCH("/update", permissionController.UpdatePermission)
        permissionRoute.DELETE("/delete", permissionController.DeletePermission)
        permissionRoute.GET("/list", permissionController.GetPermissionList)
        permissionRoute.GET("/tree", permissionController.GetPermissionTree)
        permissionRoute.GET("/select", permissionController.GetPermissionOption)
    }

    permissionRoute.Use(userController.AuthMiddleware())
    {

    }

    log.Fatal(server.Run(":9000"))
}
