package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

/*
认证中间件
*/
func (uc *UserController) AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 从header获取token
        header := ctx.GetHeader("Authorization")
        token := utils.GetTokenFromHeader(header)
        if token == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"message": "用户token获取失败"})
            ctx.Abort()
            return
        }

        // 验证token
        // 1. 判断token合法
        payload, err := utils.GetPayloadFromToken(token)
        // "error": "signature is invalid"
        if err != nil {
            if err.Error() == "Token is expired" {
                ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token过期"})
            } else {
                ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token非法"})
            }
            ctx.Abort()
            return
        }

        // 2. 判断token注销
        userID := payload["id"].(string)
        t, err := uc.UserService.ReadFromRedis(userID)
        if err != nil || t != token {
            ctx.JSON(http.StatusBadRequest, gin.H{"message": "用户token非法"})
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
用户注销
*/
func (uc *UserController) LogoutUser(ctx *gin.Context) {
    // 从header获取token
    header := ctx.GetHeader("Authorization")
    token := utils.GetTokenFromHeader(header)
    if token == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户token获取失败"})
        ctx.Abort()
        return
    }

    // 从token获取userID
    payload, err := utils.GetPayloadFromToken(token)
    // "error": "signature is invalid"
    if err != nil {
        if err.Error() == "Token is expired" {
            ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token过期"})
        } else {
            ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token非法"})
        }
        ctx.Abort()
        return
    }

    if err := uc.UserService.LogoutUser(payload["id"].(string)); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户token注销失败"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "用户成功注销"})
}

/*
用户刷新
*/
func (uc *UserController) RefreshUser(ctx *gin.Context) {
    // 从header获取token
    header := ctx.GetHeader("Authorization")
    token := utils.GetTokenFromHeader(header)
    if token == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户token获取失败"})
        ctx.Abort()
        return
    }

    // 从token获取userID
    payload, err := utils.GetPayloadFromToken(token)
    // "error": "signature is invalid"
    if err != nil {
        if err.Error() == "Token is expired" {
            ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token过期"})
        } else {
            ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token非法"})
        }
        ctx.Abort()
        return
    }

    refresh, err := uc.UserService.RefreshUser(payload["id"].(string))
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"error": "用户token刷新失败"})
        return
    }

    ctx.JSON(http.StatusOK, refresh)
}
