package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/services"
    "log"
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

/* 创建 */
func (pc *PermissionController) CreatePermission(ctx *gin.Context) {
    var permission models.Permission
    if err := ctx.ShouldBindJSON(&permission); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := pc.PermissionService.CreatePermission(&permission)
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
        "message": "权限信息成功创建",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取 */
func (pc *PermissionController) GetPermission(ctx *gin.Context) {
    id := ctx.Query("id")

    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    permission, err := pc.PermissionService.GetPermission(&id)
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
        "message": "权限信息成功获取",
        "data":    permission,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 更新 */
func (pc *PermissionController) UpdatePermission(ctx *gin.Context) {
    var permission models.Permission
    if err := ctx.ShouldBindJSON(&permission); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    log.Println("json sort_id -> ", permission.SortID)

    err := pc.PermissionService.UpdatePermission(&permission)
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
        "message": "权限信息成功更新",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 删除 */
func (pc *PermissionController) DeletePermission(ctx *gin.Context) {
    id := ctx.Query("id")
    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := pc.PermissionService.DeletePermission(&id)
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
        "message": "权限信息成功删除",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取列表 */
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

/* 获取树 */
func (pc *PermissionController) GetPermissionTree(ctx *gin.Context) {
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
        "message": "权限树成功获取",
        "data":    permissions,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取选择项 */
func (pc *PermissionController) GetPermissionOption(ctx *gin.Context) {
    permissions, err := pc.PermissionService.GetPermissionOption()
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
        "message": "权限选择项成功获取",
        "data":    permissions,
    }
    ctx.JSON(http.StatusOK, response)
}
