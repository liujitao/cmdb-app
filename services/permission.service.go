package services

import (
    "cmdb-app-mysql/models"
    "context"
    "database/sql"
)

type PermissionService interface {
    CreatePermission(*models.Permission) error
    GetPermission(*string) (*models.Permission, error)
    UpdatePermission(*models.Permission) error
    DeletePermission(*string) error
    GetPermissionList(*string, *string, *string) (*int64, []*models.Permission, error)
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

func (ps *PermissionServiceImpl) CreatePermission(role *models.Permission) error {
    return nil
}

func (ps *PermissionServiceImpl) GetPermission(id *string) (*models.Permission, error) {
    return nil, nil
}

func (ps *PermissionServiceImpl) UpdatePermission(role *models.Permission) error {
    return nil
}

func (ps *PermissionServiceImpl) DeletePermission(id *string) error {
    return nil
}

/**/
func (ps *PermissionServiceImpl) GetPermissionList(page *string, limit *string, sort *string) (*int64, []*models.Permission, error) {
    var permissions []*models.Permission
    var sql string
    var total *int64

    sql = `
    select count(*) from sys_permission
    `

    row := ps.mysqlClient.QueryRowContext(ps.ctx, sql)
    row.Scan(&total)
    if *total == 0 {
        return total, nil, nil
    }

    sql = `
        select
            a.id, a.title, a.parent_id, a.name, a.path, a.component, a.redirect, a.icon, a.sort_id, a.permission_type, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_permission a
            where id >= (select id from sys_permission limit ?, 1)
        order by ` + *sort +
        ` limit ?`

    rows, err := ps.mysqlClient.QueryContext(ps.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var permission models.Permission
        rows.Scan(&permission.ID, &permission.ParentID, &permission.Title, &permission.Name, &permission.Path, &permission.Component, &permission.Icon, &permission.Redirect, &permission.SortID, &permission.CreateAt, &permission.CreateUser, &permission.UpdateAt, &permission.UpdateUser)
        permissions = append(permissions, &permission)
    }

    return total, permissions, nil
}
