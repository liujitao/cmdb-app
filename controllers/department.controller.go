package controllers

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/services"
    "net/http"

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

/* 创建 */
func (dc *DepartmentController) CreateDepartment(ctx *gin.Context) {
    var department models.Department

    if err := ctx.ShouldBindJSON(&department); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := dc.DepartmentService.CreateDepartment(&department)
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
        "message": "部门成功创建",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取 */
func (dc *DepartmentController) GetDepartment(ctx *gin.Context) {
    id := ctx.Query("id")

    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    department, err := dc.DepartmentService.GetDepartment(&id)
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
        "message": "部门信息成功获取",
        "data":    department,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 更新 */
func (dc *DepartmentController) UpdateDepartment(ctx *gin.Context) {
    var department models.Department
    if err := ctx.ShouldBindJSON(&department); err != nil {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
            "error":   err.Error(),
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := dc.DepartmentService.UpdateDepartment(&department)
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
        "message": "部门信息成功更新",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 删除 */
func (dc *DepartmentController) DeleteDepartment(ctx *gin.Context) {
    id := ctx.Query("id")
    if len(id) == 0 {
        response := gin.H{
            "code":    10000,
            "message": "请求参数异常",
        }
        ctx.JSON(http.StatusBadGateway, response)
        return
    }

    err := dc.DepartmentService.DeleteDepartment(&id)
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
        "message": "部门信息成功删除",
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取列表 */
func (dc *DepartmentController) GetDepartmentList(ctx *gin.Context) {
    departments, err := dc.DepartmentService.GetDepartmentList()
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
        "message": "部门列表成功获取",
        "data":    departments,
    }
    ctx.JSON(http.StatusOK, response)
}

/* 获取树 */
func (dc *DepartmentController) GetDepartmentTree(ctx *gin.Context) {
    departments, err := dc.DepartmentService.GetDepartmentTree()
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
        "message": "部门树成功获取",
        "data":    departments,
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
