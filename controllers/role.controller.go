package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/services"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

type RoleController struct {
    RoleService services.RoleService
}

func NewRoleController(roleSerivce services.RoleService) RoleController {
    return RoleController{
        RoleService: roleSerivce,
    }
}

/* 创建 */
func (rc *RoleController) CreateRole(ctx *gin.Context) {
    var role models.Role
    if err := ctx.ShouldBindJSON(&role); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := rc.RoleService.CreateRole(&role)
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
        "message": "角色成功创建",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取 */
func (rc *RoleController) GetRole(ctx *gin.Context) {
    id := ctx.Query("id")

    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    role, err := rc.RoleService.GetRole(&id)
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
        "message": "角色信息成功获取",
        "data":    role,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 更新 */
func (rc *RoleController) UpdateRole(ctx *gin.Context) {
    var role models.Role
    if err := ctx.ShouldBindJSON(&role); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := rc.RoleService.UpdateRole(&role)
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
        "message": "角色成功更新",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 删除 */
func (rc *RoleController) DeleteRole(ctx *gin.Context) {
    id := ctx.Query("id")
    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := rc.RoleService.DeleteRole(&id)
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
        "message": "角色成功删除",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取列表 */
func (rc *RoleController) GetRoleList(ctx *gin.Context) {
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

    total, roles, err := rc.RoleService.GetRoleList(&page, &limit, &sort)
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
        "items": roles,
    }

    response := gin.H{
        "code":    20000,
        "message": "角色列表成功获取",
        "data":    data,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取选择项 */
func (rc *RoleController) GetRoleOption(ctx *gin.Context) {
    roles, err := rc.RoleService.GetRoleOptions()
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
        "message": "角色选择项成功获取",
        "data":    roles,
    }
    ctx.JSON(http.StatusOK, response)
}
