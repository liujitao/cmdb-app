package controllers

import (
    "log"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "cmdb-app-mysql/models"
    "cmdb-app-mysql/services"
    "cmdb-app-mysql/utils"
)

type UserController struct {
    UserService services.UserService
}

func New(userSerivce services.UserService) UserController {
    return UserController{
        UserService: userSerivce,
    }
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
    var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    err := uc.UserService.CreateUser(&user)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
    var id string
    id = ctx.Query("id")

    // id参数为空，获取当前用户信息
    if len(id) == 0 {
        // 从header获取token
        header := ctx.GetHeader("Authorization")
        token := utils.GetTokenFromHeader(header)
        if token == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户token获取失败"})
            ctx.Abort()
            return
        }

        // 从token获取userID
        payload, _ := utils.GetPayloadFromToken(token)
        id = payload["id"].(string)
    }

    log.Println(id)

    user, err := uc.UserService.GetUser(&id)
    if err != nil {
        response := gin.H{
            "code":    10000,
            "message": "服务处理异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    response := gin.H{
        "code":    20000,
        "message": "用户信息成功获取",
        "data":    user,
    }
    ctx.JSON(http.StatusOK, response)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
    var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    err := uc.UserService.UpdateUser(&user)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
    id := ctx.Query("id")

    err := uc.UserService.DeleteUser(&id)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
    users, err := uc.UserService.GetAllUser()
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
    route := rg.Group("/user")
    route.Use(cors.Default())

    authRoute := rg.Group("/user")
    authRoute.Use(cors.Default())
    authRoute.Use(uc.AuthMiddleware())

    route.POST("/create", uc.CreateUser)
    authRoute.GET("/get", uc.GetUser)
    route.PATCH("/update", uc.UpdateUser)
    route.DELETE("/delete", uc.DeleteUser)
    route.GET("/getall", uc.GetAllUser)

    route.POST("/login", uc.LoginUser)
    route.POST("/logout", uc.LogoutUser)
    route.POST("/refresh", uc.RefreshUser)
}
