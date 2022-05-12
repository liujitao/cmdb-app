package controllers

import (
    "cmdb-app-mysql/services"
    "net/http"

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
    permissions, err := pc.PermissionService.GetPermissionTree()
    if err != nil {
        response := gin.H{
            "code":    10000,
            "message": "服务处理异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    response := gin.H{
        "code":    20000,
        "message": "权限列表成功获取",
        "data":    permissions,
    }
    ctx.JSON(http.StatusOK, response)
}
