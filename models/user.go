package models

import "time"

type User struct {
    ID         string     `json:"id"`
    Avatar     string     `json:"avatar"`
    Mobile     string     `json:"mobile"`
    Email      string     `json:"email"`
    Name       string     `json:"user_name"`
    Password   string     `json:"password,omitempty"`
    Gender     int8       `json:"gender"`
    Status     int8       `json:"status"`
    CreateAt   *time.Time `json:"create_at,omitempty"`
    CreateUser string     `json:"create_user,omitempty"`
    UpdateAt   *time.Time `json:"update_at,omitempty"`
    UpdateUser string     `json:"update_user,omitempty"`
    Department []string   `json:"departments"`
    Role       []string   `json:"roles"`
}

type UserResponse struct {
    ID         string             `json:"id"`
    Avatar     string             `json:"avatar"`
    Mobile     string             `json:"mobile"`
    Email      string             `json:"email"`
    Name       string             `json:"user_name"`
    Password   string             `json:"password"`
    Gender     int8               `json:"gender"`
    Status     int8               `json:"status"`
    CreateAt   *time.Time         `json:"create_at"`
    CreateUser string             `json:"create_user"`
    UpdateAt   *time.Time         `json:"update_at"`
    UpdateUser string             `json:"update_user"`
    Department []SimpleDepartment `json:"departments"`
    Role       []SimpleRole       `json:"roles"`
    Menu       []*MenuTree        `json:"menus,omitempty"`
    Button     []Button           `json:"buttons,omitempty"`
}

type SimpleUser struct {
    ID   string `json:"id"`
    Name string `json:"user_name"`
}

const DefaultPassword = "Abcd@1234"

type UserPassword struct {
    ID       string `json:"id"`
    Password string `json:"password"`
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
