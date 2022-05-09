package services

import (
    "context"
    "database/sql"
    "time"

    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"

    "github.com/go-redis/redis/v8"
    "github.com/rs/xid"
)

type UserService interface {
    CreateUser(*models.User) error
    GetUser(*string) (*models.User, error)
    UpdateUser(*models.User) error
    DeleteUser(*string) error
    GetUserList(*string, *string, *string, *string, *string) (*int64, []*models.User, error)
    LoginUser(*models.UserLogin) (*models.Login, error)
    LogoutUser(string) error
    RefreshUser(string) (*models.Login, error)
    ReadFromRedis(string) (string, error)
    WriteToRedis(string, string, time.Duration) error
    RemoveFromRedis(string) error
}

type UserServiceImpl struct {
    mysqlClient *sql.DB
    redisClient *redis.Client
    ctx         context.Context
}

func NewUserService(mysqlClient *sql.DB, redisClient *redis.Client, ctx context.Context) UserService {
    return &UserServiceImpl{
        mysqlClient: mysqlClient,
        redisClient: redisClient,
        ctx:         ctx,
    }
}

/* 创建用户 */
func (us *UserServiceImpl) CreateUser(user *models.User) error {
    sql := `
    insert into sys_user
        (id, user_name, password, mobile, email, gender, avatar, status, admin_flag, create_at, create_user, update_at, update_user)
    values
        (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

    id := xid.New().String()
    create_at := time.Now().Local()

    _, err := us.mysqlClient.ExecContext(us.ctx, sql, id, user.Name, utils.HashPassword(user.Password), user.Mobile, user.Email, user.Gender, user.Avatar, user.Status, user.AdminFlag, create_at, user.CreateUser, create_at, user.CreateUser)
    if err != nil {
        return err
    }

    return nil
}

/* 获取用户 */
func (us *UserServiceImpl) GetUser(id *string) (*models.User, error) {
    var user models.User

    sql := `
    select
        a.id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.status, a.admin_flag, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_user a
    where a.id = ?
   `
    row := us.mysqlClient.QueryRowContext(us.ctx, sql, id)

    err := row.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Status, &user.AdminFlag, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)
    if err != nil {
        return nil, err
    }

    //获取用户关联部门
    departments, err := us.GetDepartmentByUserId(id)
    if err != nil {
        return nil, err
    }

    if departments == nil {
        user.Department = []models.SimpleDepartment{}
    } else {
        user.Department = departments
    }

    //获取用户关联角色
    roles, err := us.GetRoleByUserId(id)
    if err != nil {
        return nil, err
    }

    if roles == nil {
        user.Role = []models.SimpleRole{}
    } else {
        user.Role = roles
    }

    //获取用户关联菜单树
    menuTree, err := us.GetMenuByUserId(id)
    if err != nil {
        return nil, err
    }

    if menuTree == nil {
        user.Menu = []*models.MenuTree{}
    } else {
        user.Menu = menuTree
    }

    //获取用户关联按钮
    buttons, err := us.GetButtonByUserId(id)
    if err != nil {
        return nil, err
    }

    if buttons == nil {
        user.Button = []models.Button{}
    } else {
        user.Button = buttons
    }

    return &user, nil
}

/* 更新用户 */
func (us *UserServiceImpl) UpdateUser(user *models.User) error {
    sql := `
    update sys_user set
        user_name=?, mobile=?, email=?, gender=?, avatar=?, status=?, admin_flag=?, update_at=?, update_user=?
    where id = ?
    `

    update_at := time.Now()
    _, err := us.mysqlClient.ExecContext(us.ctx, sql, user.Name, user.Mobile, user.Email, user.Gender, user.Avatar, user.Status, user.AdminFlag, update_at, user.UpdateUser, user.ID)
    if err != nil {
        return err
    }

    return nil
}

func (us *UserServiceImpl) DeleteUser(id *string) error {
    sql := `
    delete from sys_user where id = ?
    `
    _, err := us.mysqlClient.ExecContext(us.ctx, sql, id)
    if err != nil {
        return err
    }

    return nil
}

/* 获取用户列表 */
func (us *UserServiceImpl) GetUserList(page *string, limit *string, sort *string, status *string, keyword *string) (*int64, []*models.User, error) {
    var users []*models.User
    var sql string
    var total *int64

    sql = `
    select count(*) from sys_user
    `
    row := us.mysqlClient.QueryRowContext(us.ctx, sql)
    row.Scan(&total)
    if *total == 0 {
        return total, nil, nil
    }

    sql = `
        select
            a.id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.status, a.admin_flag, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_user a
            where id >= (select id from sys_user limit ?, 1)
        order by ` + *sort +
        ` limit ?`

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var user models.User
        rows.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Status, &user.AdminFlag, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)

        //获取用户关联部门
        departments, err := us.GetDepartmentByUserId(&user.ID)
        if err != nil {
            return nil, nil, err
        }

        if departments == nil {
            user.Department = []models.SimpleDepartment{}
        } else {
            user.Department = departments
        }

        //获取用户关联角色
        roles, err := us.GetRoleByUserId(&user.ID)
        if err != nil {
            return nil, nil, err
        }

        if roles == nil {
            user.Role = []models.SimpleRole{}
        } else {
            user.Role = roles
        }

        users = append(users, &user)
    }

    return total, users, nil
}

/* 获取用户部门 */
func (us *UserServiceImpl) GetDepartmentByUserId(id *string) ([]models.SimpleDepartment, error) {
    var departments []models.SimpleDepartment
    sql := `
    select b.id, b.department_name from sys_user a
        left join sys_user_department ab on a.id = ab.user_id
            join sys_department b on ab.department_id = b.id
    where a.id = ?;
    `

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var department models.SimpleDepartment
        if err := rows.Scan(&department.ID, &department.Name); err != nil {
            return nil, err
        }
        departments = append(departments, department)
    }

    return departments, nil
}

/* 获取用户角色 */
func (us *UserServiceImpl) GetRoleByUserId(id *string) ([]models.SimpleRole, error) {
    var roles []models.SimpleRole

    sql := `
    select b.id, b.role_name from sys_user a
        left join sys_user_role ab on a.id = ab.user_id
            join sys_role b on ab.role_id = b.id
    where a.id = ?
   `

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var role models.SimpleRole
        if err := rows.Scan(&role.ID, &role.Name); err != nil {
            return nil, err
        }
        roles = append(roles, role)
    }

    return roles, nil
}

/* 获取用户菜单 */
func (us *UserServiceImpl) GetMenuByUserId(id *string) ([]*models.MenuTree, error) {
    var menus []*models.MenuTree

    sql := `
    select b.id, b.parent_id, b.title, b.name, b.path, b.component, b.redirect, b.icon, b.sort_id from sys_user a
        left join sys_user_role ac on a.id = ac.user_id
            join sys_role_permission bd on ac.role_id = bd.role_id
                join sys_permission b on b.id = bd.permission_id
    where b.permission_type = 0 and a.id = ?
    order by b.sort_id
    `

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        menu := &models.MenuTree{}
        meta := &models.Meta{}
        if err := rows.Scan(&menu.ID, &menu.ParentID, &meta.Title, &menu.Name, &menu.Path, &menu.Component, &menu.Redirect, &meta.Icon, &menu.SortID); err != nil {
            return nil, err
        }

        if menu.Path == "/" {
            menu.Meta = nil
        } else {
            menu.Meta = meta
        }
        menu.Children = nil
        menus = append(menus, menu)
    }

    // convert list To tree
    menuTree := utils.BuildMenuTree(menus, "")
    return menuTree, nil
}

/* 获取用户按钮 */
func (us *UserServiceImpl) GetButtonByUserId(id *string) ([]models.Button, error) {
    var buttons []models.Button
    sql := `
    select b.id, b.title, b.path from sys_user a
        left join sys_user_role ac on a.id = ac.user_id
            join sys_role_permission bd on ac.role_id = bd.role_id
                join sys_permission b on b.id = bd.permission_id
    where b.permission_type = 1 and a.id = ?
    order by b.sort_id
    `

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, id)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var button models.Button
        if err := rows.Scan(&button.ID, &button.Title, &button.Path); err != nil {
            return nil, err
        }
        buttons = append(buttons, button)
    }

    return buttons, nil
}
