package controllers

import (
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

func (rc *RoleController) CreateRole(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (rc *RoleController) GetRole(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (rc *RoleController) UpdateRole(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (rc *RoleController) DeleteRole(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

/* 获取角色列表 */
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
