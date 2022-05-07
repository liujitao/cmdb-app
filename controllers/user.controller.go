package controllers

import (
    "net/http"
    "strconv"
    "strings"

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

func (uc *UserController) GetUserList(ctx *gin.Context) {
    var page, limit, sort, status, keyword string
    page = ctx.DefaultQuery("page", "0")
    if page != "0" {
        p, _ := strconv.Atoi(page)
        page = strconv.Itoa(p - 1)
    }

    limit = ctx.DefaultQuery("limit", "20")

    sort = ctx.Query("sort")
    if sort == "" {
        sort = "id ASC"
    } else {
        s := strings.Split(sort, "")
        if s[0] == "-" {
            s = append(s[1:], " DESC")
        } else {
            s = append(s[1:], " ASC")
        }
        sort = strings.Join(s, "")
    }

    status = ctx.DefaultQuery("status", "")
    keyword = ctx.DefaultQuery("keyword", "")

    total, users, err := uc.UserService.GetUserList(&page, &limit, &sort, &status, &keyword)
    if err != nil {
        response := gin.H{
            "code":    10000,
            "message": "服务处理异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    data := gin.H{
        "total": total,
        "items": users,
    }

    response := gin.H{
        "code":    20000,
        "message": "用户列表成功获取",
        "data":    data,
    }
    ctx.JSON(http.StatusOK, response)
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
    authRoute.GET("/list", uc.GetUserList)

    route.POST("/login", uc.LoginUser)
    route.POST("/logout", uc.LogoutUser)
    route.POST("/refresh", uc.RefreshUser)
}
