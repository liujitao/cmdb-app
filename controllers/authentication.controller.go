package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

/*
认证中间件
*/
func (uc *UserController) AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 获取token
        header := ctx.GetHeader("Authorization")
        if header == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"message": "token error"})
            ctx.Abort()
            return
        }

        token := strings.Split(header, " ")
        if len(token) != 2 || token[0] != "Bearer" {
            ctx.JSON(http.StatusBadRequest, gin.H{"message": "token error"})
            ctx.Abort()
            return
        }

        // 验证token
        // 1. 判断token合法
        payload, err := utils.ParseToken(token[1])
        if err != nil {
            if err.Error() == "Token is expired" {
                ctx.JSON(http.StatusBadGateway, gin.H{"error": "token过期"})
            } else {
                ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
            }
            ctx.Abort()
            return
        }

        // 2. 判断token注销
        t, err := uc.UserService.ReadFromRedis(payload["id"].(string))
        if err != nil || t != token[1] {
            ctx.JSON(http.StatusBadRequest, gin.H{"message": "invaild token"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}

/*
用户登录
*/
func (uc *UserController) LoginUser(ctx *gin.Context) {
    var user models.UserLogin
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    login, err := uc.UserService.LoginUser(&user)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, login)
}

/*
用户注销 (管理员可注销其他用户，非管理员只能注销自己)
*/
func (uc *UserController) LogoutUser(ctx *gin.Context) {
    id := ctx.DefaultQuery("id", "")
    if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "request is bad"})
        return
    }

    if err := uc.UserService.LogoutUser(); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

/*
用户刷新
*/
func (uc *UserController) RefreshUser(ctx *gin.Context) {
    refresh, err := uc.UserService.RefreshUser()
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, refresh)
}
