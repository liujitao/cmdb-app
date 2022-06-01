package services

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "context"
    "database/sql"
    "strings"
    "time"

    "github.com/rs/xid"
)

type DepartmentService interface {
    CreateDepartment(*models.Department) error
    GetDepartment(*string) (*models.DepartmentResponse, error)
    UpdateDepartment(*models.Department) error
    DeleteDepartment(*string) error
    GetDepartmentList() ([]*models.DepartmentResponse, error)
    GetDepartmentTree() ([]*models.DepartmentTree, error)
    GetDepartmentOption() ([]*models.DepartmentTree, error)
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

/* 创建 */
func (ds *DepartmentServiceImpl) CreateDepartment(department *models.Department) error {
    sql := `
    insert into sys_department
        (id, parent_id, department_name, description, sort_id, create_user, create_at)
    values
        (?, ?, ?, ?, ?, ?, ?)
    `

    id := strings.ToUpper(xid.New().String())
    create_at := time.Now().Local()

    _, err := ds.mysqlClient.ExecContext(ds.ctx, sql, id, department.ParentID, department.Name, department.Description, department.SortID, department.CreateUser, create_at)
    if err != nil {
        return err
    }

    return nil
}

/* 获取 */
func (ds *DepartmentServiceImpl) GetDepartment(id *string) (*models.DepartmentResponse, error) {
    var department models.DepartmentResponse

    sql := `
    select
        id, parent_id, department_name, description, sort_id, create_at, create_user, update_at, update_user
    from sys_department
    where id = ?
   `
    row := ds.mysqlClient.QueryRowContext(ds.ctx, sql, id)

    err := row.Scan(&department.ID, &department.ParentID, &department.Name, &department.Description, &department.SortID, &department.CreateAt, &department.CreateUser, &department.UpdateAt, &department.UpdateUser)
    if err != nil {
        return nil, err
    }

    return &department, nil
}

/* 更新 */
func (ds *DepartmentServiceImpl) UpdateDepartment(department *models.Department) error {
    sql := `
    update sys_department set
        parent_id = ? , department_name = ?, description = ?, sort_id = ?, update_at = ?, update_user = ?
    where id = ?
    `

    update_at := time.Now().Local()
    id := department.ID

    _, err := ds.mysqlClient.ExecContext(ds.ctx, sql, department.ParentID, department.Name, department.Description, department.SortID, update_at, department.UpdateUser, id)
    if err != nil {
        return err
    }

    return nil
}

/* 删除 */
func (ds *DepartmentServiceImpl) DeleteDepartment(id *string) error {
    // 判断部门是否关联其他数据

    // 删除部门
    sql := `delete from sys_department where id = ?`
    _, err := ds.mysqlClient.ExecContext(ds.ctx, sql, id)
    if err != nil {
        return err
    }

    return nil
}

/* 获取列表 */
func (ds *DepartmentServiceImpl) GetDepartmentList() ([]*models.DepartmentResponse, error) {
    var departments []*models.DepartmentResponse
    sql := `
        select
            a.id, a.department_name, a.description, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_department a
        order by parent_id, id`

    rows, err := ds.mysqlClient.QueryContext(ds.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var department models.DepartmentResponse
        rows.Scan(&department.ID, &department.Name, &department.Description, &department.CreateAt, &department.CreateUser, &department.UpdateAt, &department.UpdateUser)

        //获取用户关联角色
        users, err := ds.GetUserByDepartmentID(&department.ID)
        if err != nil {
            return nil, err
        }

        if users == nil {
            department.User = []models.SimpleUser{}
        } else {
            department.User = users
        }

        departments = append(departments, &department)
    }

    return departments, nil
}

/* 获取树 */
func (ds *DepartmentServiceImpl) GetDepartmentTree() ([]*models.DepartmentTree, error) {
    var departments []*models.DepartmentTree

    sql := `select id, parent_id, department_name, description, sort_id from sys_department order by parent_id, id`
    rows, err := ds.mysqlClient.QueryContext(ds.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        department := &models.DepartmentTree{}
        if err := rows.Scan(&department.ID, &department.ParentID, &department.Name, &department.Description, &department.SortID); err != nil {
            return nil, err
        }
        department.Children = nil
        departments = append(departments, department)
    }

    // convert list To tree
    departmentTree := utils.BuildDepartmentTree(departments, "")
    return departmentTree, nil
}

/* 获取选择项 */
func (ds *DepartmentServiceImpl) GetDepartmentOption() ([]*models.DepartmentTree, error) {
    var departments []*models.DepartmentTree

    sql := `select id, parent_id, department_name from sys_department order by parent_id, id`
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
