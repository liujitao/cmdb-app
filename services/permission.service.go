package services

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "context"
    "database/sql"
    "fmt"
    "strconv"
    "time"
)

type PermissionService interface {
    CreatePermission(*models.Permission) error
    GetPermission(*string) (*models.PermissionResponse, error)
    UpdatePermission(*models.Permission) error
    DeletePermission(*string) error
    GetPermissionList() ([]*models.PermissionResponse, error)
    GetPermissionTree() ([]*models.PermissionTree, error)
    GetPermissionOption() ([]*models.PermissionTree, error)
}

type PermissionServiceImpl struct {
    mysqlClient *sql.DB
    ctx         context.Context
}

func NewPermissionService(mysqlClient *sql.DB, ctx context.Context) PermissionService {
    return &PermissionServiceImpl{
        mysqlClient: mysqlClient,
        ctx:         ctx,
    }
}

/* 创建 */
func (ps *PermissionServiceImpl) CreatePermission(permission *models.Permission) error {
    var id string
    var sql string

    // 获得最后记录的ID值
    sql = `select id from sys_permission order by id DESC limit 1`
    row := ps.mysqlClient.QueryRowContext(ps.ctx, sql)
    if err := row.Scan(&id); err != nil {
        return err
    }

    newID, _ := strconv.Atoi(id)
    id = fmt.Sprintf("%06d", newID+1)
    create_at := time.Now().Local()

    // 插入权限
    sql = `
    insert into sys_permission
        (id, parent_id, title, name, path, component, redirect, icon, permission_type, sort_id, create_user, create_at)
    values
        (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

    if _, err := ps.mysqlClient.ExecContext(ps.ctx, sql, id, permission.ParentID, permission.Title, permission.Name, permission.Path, permission.Component, permission.Redirect, permission.Icon, permission.Type, permission.SortID, permission.CreateUser, create_at); err != nil {
        return err
    }

    return nil
}

/* 获取 */
func (ps *PermissionServiceImpl) GetPermission(id *string) (*models.PermissionResponse, error) {
    var permission models.PermissionResponse

    sql := `
    select
        id, parent_id, title, name, path, component, redirect, icon, permission_type, sort_id, create_at, create_user, update_at, update_user
    from sys_permission
    where id = ?
   `
    row := ps.mysqlClient.QueryRowContext(ps.ctx, sql, id)

    err := row.Scan(&permission.ID, &permission.ParentID, &permission.Title, &permission.Name, &permission.Path, &permission.Component, &permission.Redirect, &permission.Icon, &permission.Type, &permission.SortID, &permission.CreateAt, &permission.CreateUser, &permission.UpdateAt, &permission.UpdateUser)
    if err != nil {
        return nil, err
    }

    return &permission, nil
}

/* 更新 */
func (ps *PermissionServiceImpl) UpdatePermission(permission *models.Permission) error {
    sql := `
    update sys_permission set
        parent_id = ?, title = ?, name = ?, path = ?, component = ?, redirect = ?, icon = ?, permission_type = ?, sort_id = ?, update_at = ?, update_user = ?
    where id = ?
    `

    update_at := time.Now().Local()
    id := permission.ID

    _, err := ps.mysqlClient.ExecContext(ps.ctx, sql, permission.ParentID, permission.Title, permission.Name, permission.Path, permission.Component, permission.Redirect, permission.Icon, permission.Type, permission.SortID, update_at, permission.UpdateUser, id)
    if err != nil {
        return err
    }

    return nil
}

/* 删除 */
func (ps *PermissionServiceImpl) DeletePermission(id *string) error {
    // 判断权限是否关联其他数据

    // 删除权限
    sql := `delete from sys_permission where id = ?`
    _, err := ps.mysqlClient.ExecContext(ps.ctx, sql, id)
    if err != nil {
        return err
    }

    return nil
}

/* 获取列表 */
func (ps *PermissionServiceImpl) GetPermissionList() ([]*models.PermissionResponse, error) {
    var permissions []*models.PermissionResponse

    sql := `
        select
            id, parent_id, title, name, path, component, redirect, icon, permission_type, sort_id, create_at, create_user, update_at, update_user, create_at, create_user, update_at, update_user
        from sys_permission
        order by parent_id, id`

    rows, err := ps.mysqlClient.QueryContext(ps.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var permission models.PermissionResponse
        rows.Scan(&permission.ID, &permission.ParentID, &permission.Title, &permission.Name, &permission.Path, &permission.Component, &permission.Redirect, &permission.Icon, &permission.Type, &permission.SortID, &permission.CreateAt, &permission.CreateUser, &permission.UpdateAt, &permission.UpdateUser)

        permissions = append(permissions, &permission)
    }

    return permissions, nil
}

/* 获取树 */
func (ps *PermissionServiceImpl) GetPermissionTree() ([]*models.PermissionTree, error) {
    var permissions []*models.PermissionTree

    sql := `
    select
        a.id, a.title, a.parent_id, a.name, a.path, a.component, a.redirect, a.icon, a.sort_id, a.permission_type, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_permission a
    order by a.permission_type, a.sort_id
    `

    rows, err := ps.mysqlClient.QueryContext(ps.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        permission := &models.PermissionTree{}
        if err := rows.Scan(&permission.ID, &permission.Title, &permission.ParentID, &permission.Name, &permission.Path, &permission.Component, &permission.Redirect, &permission.Icon, &permission.SortID, &permission.Type, &permission.CreateAt, &permission.CreateUser, &permission.UpdateAt, &permission.UpdateUser); err != nil {
            return nil, err
        }
        permission.Children = nil
        permissions = append(permissions, permission)
    }

    // convert list To tree
    permissionTree := utils.BuildPremissionTree(permissions, "")
    return permissionTree, nil
}

/* 获取选择项 */
func (ps *PermissionServiceImpl) GetPermissionOption() ([]*models.PermissionTree, error) {
    var permissions []*models.PermissionTree

    sql := `select id, parent_id, title from sys_permission order by parent_id, id`
    rows, err := ps.mysqlClient.QueryContext(ps.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        permission := &models.PermissionTree{}
        if err := rows.Scan(&permission.ID, &permission.ParentID, &permission.Title); err != nil {
            return nil, err
        }
        permission.Children = nil
        permissions = append(permissions, permission)
    }

    // convert list To tree
    permissionTree := utils.BuildPremissionTree(permissions, "")
    return permissionTree, nil
}
