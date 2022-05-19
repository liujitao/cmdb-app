package services

import (
    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"
    "context"
    "database/sql"
)

type PermissionService interface {
    CreatePermission(*models.Permission) error
    GetPermission(*string) (*models.Permission, error)
    UpdatePermission(*models.Permission) error
    DeletePermission(*string) error
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
