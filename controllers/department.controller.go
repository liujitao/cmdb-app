package controllers

import (
    "cmdb-app-mysql/services"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

type DepartmentController struct {
    DepartmentService services.DepartmentService
}

func NewDepartmentController(departmentSerivce services.DepartmentService) DepartmentController {
    return DepartmentController{
        DepartmentService: departmentSerivce,
    }
}

func (dc *DepartmentController) CreateDepartment(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (dc *DepartmentController) GetDepartment(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (dc *DepartmentController) UpdateDepartment(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (dc *DepartmentController) DeleteDepartment(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

/* 获取部门列表 */
func (dc *DepartmentController) GetDepartmentList(ctx *gin.Context) {
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

    total, departments, err := dc.DepartmentService.GetDepartmentList(&page, &limit, &sort)
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
        "items": departments,
    }

    response := gin.H{
        "code":    20000,
        "message": "部门列表成功获取",
        "data":    data,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取选择项 */
func (dc *DepartmentController) GetDepartmentOption(ctx *gin.Context) {
    departments, err := dc.DepartmentService.GetDepartmentOption()
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
        "message": "部门选择项成功获取",
        "data":    departments,
    }
    ctx.JSON(http.StatusOK, response)
}
