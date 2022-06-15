package services

import (
    "cmdb-app-mysql/models"
    "context"
    "database/sql"
    "fmt"
    "strconv"
    "strings"
    "time"
)

type RoleService interface {
    CreateRole(*models.Role) error
    GetRole(*string) (*models.RoleResponse, error)
    UpdateRole(*models.Role) error
    DeleteRole(*string) error
    GetRoleList(*string, *string, *string) (*int64, []*models.RoleResponse, error)
    GetRoleOptions() ([]*models.SimpleRole, error)
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

/* 创建 */
func (rs *RoleServiceImpl) CreateRole(role *models.Role) error {
    var id string
    var sql string
    var permission []string

    // 获得最后记录的ID值
    sql = `select id from sys_role order by id DESC limit 1`
    row := rs.mysqlClient.QueryRowContext(rs.ctx, sql)
    if err := row.Scan(&id); err != nil {
        return err
    }

    newID, _ := strconv.Atoi(id)
    id = fmt.Sprintf("%06d", newID+1)
    create_at := time.Now().Local()

    // 插入角色
    sql = `
        insert into sys_role
            (id, role_name, description, create_at, create_user)
        values
            (?, ?, ?, ?, ?)
        `

    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql, id, role.Name, role.Description, create_at, role.CreateUser); err != nil {
        return err
    }

    // 插入角色权限关联
    if len(role.Permission) == 0 {
        return nil
    }

    for _, item := range role.Permission {
        permission = append(permission, "('"+id+"', '"+item+"')")
    }

    sql = `insert into sys_role_permission values ` + strings.Join(permission, ",")

    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql); err != nil {
        return err
    }

    return nil
}

/* 获取 */
func (rs *RoleServiceImpl) GetRole(id *string) (*models.RoleResponse, error) {
    var role models.RoleResponse

    sql := `
    select
        a.id, a.role_name, a.description, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_role a
    where a.id = ?
   `
    row := rs.mysqlClient.QueryRowContext(rs.ctx, sql, id)

    err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreateAt, &role.CreateUser, &role.UpdateAt, &role.UpdateUser)
    if err != nil {
        return nil, err
    }

    //获取角色关联用户
    users, err := rs.GetUserByRoleID(id)
    if err != nil {
        return nil, err
    }

    if users == nil {
        role.User = []models.SimpleUser{}
    } else {
        role.User = users
    }

    //获取角色关联权限
    permissions, err := rs.GetPermissionByRoleID(id)
    if err != nil {
        return nil, err
    }

    if permissions == nil {
        role.Permission = []models.SimplePermission{}
    } else {
        role.Permission = permissions
    }

    return &role, nil
}

/* 更新 */
func (rs *RoleServiceImpl) UpdateRole(role *models.Role) error {
    var sql string
    var permission []string

    sql = `
    update sys_role set
        role_name= ?, description = ?, update_at= ?, update_user= ?
    where id = ?
    `
    id := role.ID
    update_at := time.Now().Local()

    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql, role.Name, role.Description, update_at, role.UpdateUser, id); err != nil {
        return err
    }

    // 删除角色权限关联
    sql = `delete from sys_role_permission where role_id = ?`
    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql, id); err != nil {
        return err
    }

    // 更新角色权限关联
    if len(role.Permission) == 0 {
        return nil
    }

    for _, item := range role.Permission {
        permission = append(permission, "('"+id+"', '"+item+"')")
    }

    sql = `insert into sys_role_permission values ` + strings.Join(permission, ",")
    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql); err != nil {
        return err
    }

    return nil
}

/* 删除 */
func (rs *RoleServiceImpl) DeleteRole(id *string) error {
    var sql string

    // 删除用户部门角色关联
    sql = `delete from sys_role_permission where role_id = ?`
    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql, id); err != nil {
        return err
    }

    // 删除角色
    sql = `delete from sys_role where id = ?`
    if _, err := rs.mysqlClient.ExecContext(rs.ctx, sql, id); err != nil {
        return err
    }

    return nil
}

/* 获取列表 */
func (rs *RoleServiceImpl) GetRoleList(page *string, limit *string, sort *string) (*int64, []*models.RoleResponse, error) {
    var roles []*models.RoleResponse
    var sql string
    var total *int64

    sql = `select count(*) from sys_role`

    row := rs.mysqlClient.QueryRowContext(rs.ctx, sql)
    row.Scan(&total)
    if *total == 0 {
        return total, nil, nil
    }

    sql = `
    select
        a.id, a.role_name, a.description, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_role a
    order by ` + *sort +
        ` limit ?, ?`

    rows, err := rs.mysqlClient.QueryContext(rs.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var role models.RoleResponse
        rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreateAt, &role.CreateUser, &role.UpdateAt, &role.UpdateUser)

        // 获取角色关联用户
        users, err := rs.GetUserByRoleID(&role.ID)
        if err != nil {
            return nil, nil, err
        }

        if users == nil {
            role.User = []models.SimpleUser{}
        } else {
            role.User = users
        }

        // 获取角色关联权限
        permissions, err := rs.GetPermissionByRoleID(&role.ID)
        if err != nil {
            return nil, nil, err
        }

        if permissions == nil {
            role.Permission = []models.SimplePermission{}
        } else {
            role.Permission = permissions
        }

        roles = append(roles, &role)
    }

    return total, roles, nil
}

/* 获取选择项 */
func (rs *RoleServiceImpl) GetRoleOptions() ([]*models.SimpleRole, error) {
    var roles []*models.SimpleRole

    sql := `select id, role_name from sys_role order by role_name`
    rows, err := rs.mysqlClient.QueryContext(rs.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        role := &models.SimpleRole{}
        if err := rows.Scan(&role.ID, &role.Name); err != nil {
            return nil, err
        }
        roles = append(roles, role)
    }

    return roles, nil
}

/* 获取角色关联用户 */
func (rs *RoleServiceImpl) GetUserByRoleID(id *string) ([]models.SimpleUser, error) {
    var users []models.SimpleUser

    sql := `
    select b.id, b.user_name from sys_role a
        left join sys_user_role ab on a.id = ab.role_id
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

/* 获取角色关联权限 */
func (rs *RoleServiceImpl) GetPermissionByRoleID(id *string) ([]models.SimplePermission, error) {
    var permissions []models.SimplePermission

    sql := `
    select b.id, b.title from sys_role a
        left join sys_role_permission ab on a.id = ab.role_id
            join sys_permission b on ab.permission_id = b.id
    where a.id = ?
    order by b.parent_id, b.id
   `

    rows, err := rs.mysqlClient.QueryContext(rs.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var permission models.SimplePermission
        if err := rows.Scan(&permission.ID, &permission.Title); err != nil {
            return nil, err
        }
        permissions = append(permissions, permission)
    }

    return permissions, nil
}
