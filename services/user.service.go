package services

import (
    "context"
    "database/sql"
    "log"
    "time"

    "cmdb-app/models"
    "cmdb-app/utils"

    "github.com/rs/xid"
)

type UserService interface {
    CreateUser(*models.User) error
    GetUser(*string) (*models.User, error)
    UpdateUser(*models.User) error
    DeleteUser(*string) error
    GetAllUser() ([]*models.User, error)
}

type UserServiceImpl struct {
    mysqlClient *sql.DB
    ctx         context.Context
}

func NewUserService(mysqlClient *sql.DB, ctx context.Context) UserService {
    return &UserServiceImpl{
        mysqlClient: mysqlClient,
        ctx:         ctx,
    }
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
    sql := `
    insert into sys_user
        (user_id, user_name, password, mobile, email, gender, avatar, department_id, status, admin_flag, create_at, create_user, update_at, update_user)
    values
        (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

    user_id := xid.New().String()
    create_at := time.Now().Local()

    _, err := u.mysqlClient.ExecContext(u.ctx, sql, user_id, user.Name, utils.EncryptPassword(user.Password), user.Mobile, user.Email, user.Gender, user.Avatar, user.DepartmentID, user.Status, user.AdminFlag, create_at, user.CreateUser, create_at, user.CreateUser)
    if err != nil {
        return err
    }

    return nil
}

func (u *UserServiceImpl) GetUser(id *string) (*models.User, error) {
    var user models.User

    sql := `
    select
        a.user_id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.department_id, a.status, a.admin_flag, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_user a
    where a.user_id=?
   `
    row := u.mysqlClient.QueryRowContext(u.ctx, sql, id)
    log.Println(id, &id)

    err := row.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.DepartmentID, &user.Status, &user.AdminFlag, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
    sql := `
    update sys_user
    set
        user_name=?, mobile=?, email=?, gender=?, avatar=?, department_id=?, status=?, admin_flag=?, update_at=?, update_user=?
    where user_id = ?
    `

    update_at := time.Now()
    _, err := u.mysqlClient.ExecContext(u.ctx, sql, user.Name, user.Mobile, user.Email, user.Gender, user.Avatar, user.DepartmentID, user.Status, user.AdminFlag, update_at, user.UpdateUser, user.ID)
    if err != nil {
        return err
    }

    return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
    sql := `
    delete 
    from sys_user 
    where user_id=?
    `
    _, err := u.mysqlClient.ExecContext(u.ctx, sql, id)
    if err != nil {
        return err
    }

    return nil
}

func (u *UserServiceImpl) GetAllUser() ([]*models.User, error) {
    var users []*models.User

    sql := `
    select
        a.user_id, a.avatar, a.mobile, a.email, a.user_name, a.password, a.gender, a.department_id, a.status, a.admin_flag, a.create_at, a.create_user, a.update_at, a.update_user
    from sys_user a
   `
    rows, err := u.mysqlClient.QueryContext(u.ctx, sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var user models.User
        rows.Scan(&user.ID, &user.Avatar, &user.Mobile, &user.Email, &user.Name, &user.Password, &user.Gender, &user.DepartmentID, &user.Status, &user.AdminFlag, &user.CreateAt, &user.CreateUser, &user.UpdateAt, &user.UpdateUser)
        users = append(users, &user)
    }
    return users, nil
}
