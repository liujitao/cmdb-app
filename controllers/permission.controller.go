package controllers

import (
    "cmdb-app-mysql/services"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

type PermissionController struct {
    PermissionService services.PermissionService
}

func NewPermissionController(permissionSerivce services.PermissionService) PermissionController {
    return PermissionController{
        PermissionService: permissionSerivce,
    }
}

func (pc *PermissionController) CreatePermission(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (pc *PermissionController) GetPermission(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (pc *PermissionController) UpdatePermission(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (pc *PermissionController) DeletePermission(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

/* 获取权限列表 */
func (pc *PermissionController) GetPermissionList(ctx *gin.Context) {
    var page, limit, sort string
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

    total, permissions, err := pc.PermissionService.GetPermissionList(&page, &limit, &sort)
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
        "items": permissions,
    }

    response := gin.H{
        "code":    20000,
        "message": "权限列表成功获取",
        "data":    data,
    }
    ctx.JSON(http.StatusOK, response)
}
