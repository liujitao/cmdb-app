package services

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "context"
    "database/sql"
)

type DepartmentService interface {
    // CreateDepartment(*models.Department) error
    // GetDepartment(*string) (*models.Department, error)
    // UpdatDepartment(*models.Department) error
    // DeleteDepartment(*string) error
    GetDepartmentList(*string, *string, *string) (*int64, []*models.Department, error)
    GetDepartmentTree() ([]*models.DepartmentTree, error)
}

type DepartmentServiceImpl struct {
    mysqlClient *sql.DB
    ctx         context.Context
}

func NewDepartmentService(mysqlClient *sql.DB, ctx context.Context) DepartmentService {
    return &DepartmentServiceImpl{
        mysqlClient: mysqlClient,
        ctx:         ctx,
    }
}

func (ds *DepartmentServiceImpl) CreateDepartment(role *models.Department) error {
    return nil
}

func (ds *DepartmentServiceImpl) GetDepartment(id *string) (*models.Department, error) {
    return nil, nil
}

func (ds *DepartmentServiceImpl) UpdateDepartment(role *models.Department) error {
    return nil
}

func (ds *DepartmentServiceImpl) DeleteDepartment(id *string) error {
    return nil
}

/* 获取部门列表 */
func (ds *DepartmentServiceImpl) GetDepartmentList(page *string, limit *string, sort *string) (*int64, []*models.Department, error) {
    var departments []*models.Department
    var sql string
    var total *int64

    sql = `
    select count(*) from sys_department
    `

    row := ds.mysqlClient.QueryRowContext(ds.ctx, sql)
    row.Scan(&total)
    if *total == 0 {
        return total, nil, nil
    }

    sql = `
        select
            a.id, a.department_name, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_department a
            where id >= (select id from sys_department limit ?, 1)
        order by ` + *sort +
        ` limit ?`

    rows, err := ds.mysqlClient.QueryContext(ds.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var department models.Department
        rows.Scan(&department.ID, &department.Name, &department.CreateAt, &department.CreateUser, &department.UpdateAt, &department.UpdateUser)

        //获取用户关联角色
        users, err := ds.GetUserByDepartmentID(&department.ID)
        if err != nil {
            return nil, nil, err
        }

        if users == nil {
            department.User = []models.SimpleUser{}
        } else {
            department.User = users
        }

        departments = append(departments, &department)
    }

    return total, departments, nil
}

/* 获取部门树 */
func (ds *DepartmentServiceImpl) GetDepartmentTree() ([]*models.DepartmentTree, error) {
    var departments []*models.DepartmentTree

    sql := `select id, parent_id, department_name from sys_department order by department_name`

    rows, err := ds.mysqlClient.QueryContext(ds.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        department := &models.DepartmentTree{}
        if err := rows.Scan(&department.ID, &department.ParentID, &department.Name); err != nil {
            return nil, err
        }
        department.Children = nil
        departments = append(departments, department)
    }

    // convert list To tree
    departmentTree := utils.BuildDepartmentTree(departments, "")
    return departmentTree, nil
}

/* 获取角色用户 */
func (rs *DepartmentServiceImpl) GetUserByDepartmentID(id *string) ([]models.SimpleUser, error) {
    var users []models.SimpleUser

    sql := `
    select b.id, b.user_name from sys_department a
        left join sys_user_department ab on a.id = ab.department_id
            join sys_user b on ab.user_id = b.id
    where a.id = ?
    order by b.user_name
   `

    rows, err := rs.mysqlClient.QueryContext(rs.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var user models.SimpleUser
        if err := rows.Scan(&user.ID, &user.Name); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}
