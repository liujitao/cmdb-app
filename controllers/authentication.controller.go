package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

/*
跨域中间件
*/

/*
认证中间件
*/
func (uc *UserController) AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 从header获取token
        header := ctx.GetHeader("Authorization")
        token := utils.GetTokenFromHeader(header)
        if token == "" {
            response := gin.H{
                "code":    10000,
                "message": "请求参数异常",
            }
            ctx.JSON(http.StatusBadRequest, response)
            ctx.Abort()
            return
        }

        // 验证token
        // 1. 判断token合法
        payload, err := utils.GetPayloadFromToken(token)
        // "error": "signature is invalid"
        if err != nil {
            if err.Error() == "Token is expired" {
                response := gin.H{
                    "code":    50014,
                    "message": "用户token过期",
                    "error":   err.Error(),
                }
                ctx.JSON(http.StatusBadRequest, response)
            } else {
                response := gin.H{
                    "code":    50008,
                    "message": "用户token非法",
                    "error":   err.Error(),
                }
                ctx.JSON(http.StatusBadRequest, response)
            }
            ctx.Abort()
            return
        }

        // 2. 判断token注销
        userID := payload["id"].(string)
        t, err := uc.UserService.ReadFromRedis(userID)
        if err != nil || t != token {
            response := gin.H{
                "code":    50008,
                "message": "用户token非法",
                "error":   err.Error(),
            }
            ctx.JSON(http.StatusBadRequest, response)
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
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    login, err := uc.UserService.LoginUser(&user)
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
        "message": "用户成功登录",
        "data":    login,
    }
    ctx.JSON(http.StatusOK, response)
}

/*
用户注销
*/
func (uc *UserController) LogoutUser(ctx *gin.Context) {
    // 从header获取token
    header := ctx.GetHeader("Authorization")
    token := utils.GetTokenFromHeader(header)
    if token == "" {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadRequest, response)
        ctx.Abort()
        return
    }

    // 从token获取userID
    payload, err := utils.GetPayloadFromToken(token)
    // "error": "signature is invalid"
    if err != nil {
        if err.Error() == "Token is expired" {
            response := gin.H{
                "code":    50014,
                "message": "用户token过期",
                "error":   err.Error(),
            }
            ctx.JSON(http.StatusBadRequest, response)
        } else {
            response := gin.H{
                "code":    50008,
                "message": "用户token非法",
                "error":   err.Error(),
            }
            ctx.JSON(http.StatusBadRequest, response)
        }
        ctx.Abort()
        return
    }

    if err := uc.UserService.LogoutUser(payload["id"].(string)); err != nil {
        response := gin.H{
            "code":    50001,
            "message": "用户注销失败",
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    response := gin.H{
        "code":    20000,
        "message": "用户注销成功",
    }
    ctx.JSON(http.StatusOK, response)
}

/*
用户刷新
*/
func (uc *UserController) RefreshUser(ctx *gin.Context) {
    // 从header获取token
    header := ctx.GetHeader("Authorization")
    token := utils.GetTokenFromHeader(header)
    if token == "" {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadRequest, response)
        ctx.Abort()
        return
    }

    // 从token获取userID
    payload, err := utils.GetPayloadFromToken(token)
    // "error": "signature is invalid"
    if err != nil {
        if err.Error() == "Token is expired" {
            response := gin.H{
                "code":    50015,
                "message": "用户refresh_token过期",
                "error":   err.Error(),
            }
            ctx.JSON(http.StatusBadRequest, response)
        } else {
            response := gin.H{
                "code":    50008,
                "message": "用户refresh_token非法",
                "error":   err.Error(),
            }
            ctx.JSON(http.StatusBadRequest, response)
        }
        ctx.Abort()
        return
    }

    refresh, err := uc.UserService.RefreshUser(payload["id"].(string))
    if err != nil {
        response := gin.H{
            "code":    50002,
            "message": "用户刷新失败",
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    response := gin.H{
        "code":    20000,
        "message": "用户刷新成功",
        "data":    refresh,
    }
    ctx.JSON(http.StatusOK, response)
}
