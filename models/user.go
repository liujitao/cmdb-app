package models

import "time"

type User struct {
    ID         string             `json:"id"`
    Avatar     string             `json:"avatar"`
    Mobile     string             `json:"mobile"`
    Email      string             `json:"email"`
    Name       string             `json:"user_name"`
    Password   string             `json:"password"`
    Gender     int8               `json:"gender"`
    Status     int8               `json:"status"`
    AdminFlag  int8               `json:"admin_flag"`
    CreateAt   *time.Time         `json:"create_at"`
    CreateUser string             `json:"create_user"`
    UpdateAt   *time.Time         `json:"update_at"`
    UpdateUser string             `json:"update_user"`
    Department []SimpleDepartment `json:"department"`
    Role       []SimpleRole       `json:"roles"`
    Menu       []*MenuTree        `json:"menus,omitempty"`
    Button     []Button           `json:"buttons,omitempty"`
}

type SimpleUser struct {
    ID   string `json:"id"`
    Name string `json:"user_name"`
}

type PasswordChange struct {
    ID          string `json:"id"`
    OldPassword string `json:"old_password"`
    NewPassword string `json:"new_password"`
}

type PasswordReset struct {
    ID          string `json:"id"`
    NewPassword string `json:"new_password"`
}

type UserLogin struct {
    LoginID  string `json:"login_id"`
    Password string `json:"password"`
}

type Login struct {
    Header               string `json:"header"`
    Type                 string `json:"token_type"`
    Token                string `json:"token"`
    RefreshToken         string `json:"refresh_token"`
    TokenExpireAt        int64  `json:"token_expire_at"`
    RefreshTokenExpireAT int64  `json:"refresh_token_expire_at"`
}
