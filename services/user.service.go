package services

import (
    "context"
    "database/sql"
    "fmt"
    "strconv"
    "strings"
    "time"

    "cmdb-app-mysql/models"
    "cmdb-app-mysql/utils"

    "github.com/go-redis/redis/v8"
)

type UserService interface {
    CreateUser(*models.User) error
    GetUser(*string) (*models.UserResponse, error)
    UpdateUser(*models.User) error
    DeleteUser(*string) error
    GetUserList(*string, *string, *string, *string, *string) (*int64, []*models.UserResponse, error)
    ChangeUserPassword(*models.UserPassword) error
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

/* 创建 */
func (us *UserServiceImpl) CreateUser(user *models.User) error {
    var id string
    var sql string
    var department, role []string

    // 获得最后记录的ID值
    sql = `select id from sys_user order by id DESC limit 1`
    row := us.mysqlClient.QueryRowContext(us.ctx, sql)
    if err := row.Scan(&id); err != nil {
        return err
    }

    newID, _ := strconv.Atoi(id)
    id = fmt.Sprintf("%09d", newID+1)
    create_at := time.Now().Local()
    password := utils.HashPassword(models.DefaultPassword)

    // 插入用户
    sql = `
        insert into sys_user
            (id, user_name, password, mobile, email, gender, avatar, status, create_at, create_user)
        values
            (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        `

    if _, err := us.mysqlClient.ExecContext(us.ctx, sql, id, user.Name, password, user.Mobile, user.Email, user.Gender, user.Avatar, user.Status, create_at, user.CreateUser); err != nil {
        return err
    }

    // 插入用户部门关联
    for _, item := range user.Department {
        department = append(department, "('"+id+"', '"+item+"')")
    }

    sql = `insert into sys_user_department values ` + strings.Join(department, ",")
    if _, err := us.mysqlClient.ExecContext(us.ctx, sql); err != nil {
        return err
    }

    // 插入用户角色关联
    if len(user.Role) == 0 {
        return nil
    }

    for _, item := range user.Role {
        role = append(role, "('"+id+"', '"+item+"')")
    }

    sql = `insert into sys_user_role values ` + strings.Join(role, ",")
    if _, err := us.mysqlClient.ExecContext(us.ctx, sql); err != nil {
        return err
    }

    return nil
}

/* 获取 */
func (us *UserServiceImpl) GetUser(id *string) (*models.UserResponse, error) {
    var user models.UserResponse

    sql := `
    select
        a.id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.status, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_user a
    where a.id = ?
   `
    row := us.mysqlClient.QueryRowContext(us.ctx, sql, id)

    err := row.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Status, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)
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

    //获取用户关联菜单
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

/* 更新 */
func (us *UserServiceImpl) UpdateUser(user *models.User) error {
    var sql string
    var err error
    var department, role []string

    sql = `
    update sys_user set
        user_name=?, mobile=?, email=?, gender=?, avatar=?, status=?, update_at=?, update_user=?
    where id = ?
    `

    update_at := time.Now().Local()
    id := user.ID

    _, err = us.mysqlClient.ExecContext(us.ctx, sql, user.Name, user.Mobile, user.Email, user.Gender, user.Avatar, user.Status, update_at, user.UpdateUser, id)
    if err != nil {
        return err
    }

    // 删除用户部门角色关联
    sql = `delete from sys_user_role where user_id = ?`
    _, err = us.mysqlClient.ExecContext(us.ctx, sql, id)
    if err != nil {
        return err
    }

    sql = `delete from sys_user_department where user_id = ?`
    _, err = us.mysqlClient.ExecContext(us.ctx, sql, id)
    if err != nil {
        return err
    }

    // 更新用户部门关联
    for _, item := range user.Department {
        department = append(department, "('"+id+"', '"+item+"')")
    }
    sql = `insert into sys_user_department values ` + strings.Join(department, ",")

    _, err = us.mysqlClient.ExecContext(us.ctx, sql)
    if err != nil {
        return err
    }

    // 更新用户角色关联
    if len(user.Role) == 0 {
        return nil
    }

    for _, item := range user.Role {
        role = append(role, "('"+id+"', '"+item+"')")
    }
    sql = `insert into sys_user_role values ` + strings.Join(role, ",")

    _, err = us.mysqlClient.ExecContext(us.ctx, sql)
    if err != nil {
        return err
    }

    return nil
}

/* 删除 */
func (us *UserServiceImpl) DeleteUser(id *string) error {
    var sql string

    // 删除用户角色关联
    sql = `delete from sys_user_role where user_id = ?`
    if _, err := us.mysqlClient.ExecContext(us.ctx, sql, id); err != nil {
        return err
    }

    // 删除用户部门关联
    sql = `delete from sys_user_department where user_id = ?`
    if _, err := us.mysqlClient.ExecContext(us.ctx, sql, id); err != nil {
        return err
    }

    // 删除用户
    sql = `delete from sys_user where id = ?`
    if _, err := us.mysqlClient.ExecContext(us.ctx, sql, id); err != nil {
        return err
    }

    return nil
}

/* 获取列表 */
func (us *UserServiceImpl) GetUserList(page *string, limit *string, sort *string, status *string, keyword *string) (*int64, []*models.UserResponse, error) {
    var users []*models.UserResponse
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
            a.id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.status, a.create_at, a.create_user, a.update_at, a.update_user
        from sys_user a
        order by ` + *sort +
        ` limit ?, ?`

    rows, err := us.mysqlClient.QueryContext(us.ctx, sql, page, limit)
    if err != nil {
        return nil, nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var user models.UserResponse
        rows.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.Status, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)

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

/* 获取用户关联部门 */
func (us *UserServiceImpl) GetDepartmentByUserId(id *string) ([]models.SimpleDepartment, error) {
    var departments []models.SimpleDepartment
    sql := `
    select b.id, b.department_name from sys_user a
        left join sys_user_department ab on a.id = ab.user_id
            join sys_department b on ab.department_id = b.id
    where a.id = ?
    order by b.department_name
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

/* 获取用户关联角色 */
func (us *UserServiceImpl) GetRoleByUserId(id *string) ([]models.SimpleRole, error) {
    var roles []models.SimpleRole

    sql := `
    select b.id, b.role_name from sys_user a
        left join sys_user_role ab on a.id = ab.user_id
            join sys_role b on ab.role_id = b.id
    where a.id = ?
    order by b.role_name
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

/* 获取用户关联菜单 */
func (us *UserServiceImpl) GetMenuByUserId(id *string) ([]*models.MenuTree, error) {
    var menus []*models.MenuTree

    sql := `
    select b.id, b.parent_id, b.title, b.name, b.path, b.component, b.redirect, b.icon, b.sort_id from sys_user a
        left join sys_user_role ac on a.id = ac.user_id
            join sys_role_permission bd on ac.role_id = bd.role_id
                join sys_permission b on b.id = bd.permission_id
    where b.permission_type in (0, 1) and a.id = ?
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

/* 获取用户关联按钮 */
func (us *UserServiceImpl) GetButtonByUserId(id *string) ([]models.Button, error) {
    var buttons []models.Button
    sql := `
    select b.id, b.title, b.path from sys_user a
        left join sys_user_role ac on a.id = ac.user_id
            join sys_role_permission bd on ac.role_id = bd.role_id
                join sys_permission b on b.id = bd.permission_id
    where b.permission_type = 2 and a.id = ?
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

/* 变更用户密码 */
func (us *UserServiceImpl) ChangeUserPassword(user *models.UserPassword) error {
    var password string
    sql := `update sys_user set password = ? where id = ?`
    if user.Password == "" {
        password = utils.HashPassword(models.DefaultPassword)
    } else {
        password = user.Password
    }
    _, err := us.mysqlClient.ExecContext(us.ctx, sql, password, user.ID)
    if err != nil {
        return err
    }

    return nil
}
