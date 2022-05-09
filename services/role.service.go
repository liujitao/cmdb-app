package services

import (
    "cmdb-app-mysql/models"
    "context"
    "database/sql"
)

type RoleService interface {
    CreateRole(*models.Role) error
    GetRole(*string) (*models.Role, error)
    UpdateRole(*models.Role) error
    DeleteRole(*string) error
    GetRoleList(*string, *string, *string) (*int64, []*models.Role, error)
}

type RoleServiceImpl struct {
    mysqlClient *sql.DB
    ctx         context.Context
}

func NewRoleService(mysqlClient *sql.DB, ctx context.Context) RoleService {
    return &RoleServiceImpl{
        mysqlClient: mysqlClient,
        ctx:         ctx,
    }
}

func (rs *RoleServiceImpl) CreateRole(role *models.Role) error {
    return nil
}

func (rs *RoleServiceImpl) GetRole(id *string) (*models.Role, error) {
    return nil, nil
}

func (rs *RoleServiceImpl) UpdateRole(role *models.Role) error {
    return nil
}

func (rs *RoleServiceImpl) DeleteRole(id *string) error {
    return nil
}

/**/
func (rs *RoleServiceImpl) GetRoleList(page *string, limit *string, sort *string) (*int64, []*models.Role, error) {
    var roles []*models.Role
    var sql string
    var total *int64

    sql = `
    select count(*) from sys_role
    `

    row := rs.mysqlClient.QueryRowContext(rs.ctx, sql)
    row.Scan(&total)
    if *total == 0 {
        return total, nil, nil
    }

    sql = `
        select
            a.id, a.role_name, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_role a
            where id >= (select id from sys_role limit ?, 1)
        order by ` + *sort +
        ` limit ?`

    rows, err := rs.mysqlClient.QueryContext(rs.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var role models.Role
        rows.Scan(&role.ID, &role.Name, &role.CreateAt, &role.CreateUser, &role.UpdateAt, &role.UpdateUser)

        //获取用户关联角色
        users, err := rs.GetUserByRoleID(&role.ID)
        if err != nil {
            return nil, nil, err
        }

        if users == nil {
            role.User = []models.SimpleUser{}
        } else {
            role.User = users
        }

        roles = append(roles, &role)
    }

    return total, roles, nil
}

/* 获取角色用户 */
func (rs *RoleServiceImpl) GetUserByRoleID(id *string) ([]models.SimpleUser, error) {
    var users []models.SimpleUser

    sql := `
    select b.id, b.user_name from sys_role a
        left join sys_user_role ab on a.id = ab.role_id
            join sys_user b on ab.user_id = b.id
    where a.id = ?
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
