package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "cmdb-app-mysql/models"
    "cmdb-app-mysql/services"
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
    id := ctx.Query("id")

    user, err := uc.UserService.GetUser(&id)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, user)
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
    authRoute := rg.Group("/user")
    authRoute.Use(uc.AuthMiddleware())

    route.POST("/create", uc.CreateUser)
    authRoute.GET("/get", uc.GetUser)
    route.PATCH("/update", uc.UpdateUser)
    route.DELETE("/delete", uc.DeleteUser)
    route.GET("/getall", uc.GetAllUser)

    route.POST("/login", uc.LoginUser)
    route.POST("/logout", uc.LogoutUser)
    route.POST("/refresh", uc.CreateUser)
}
