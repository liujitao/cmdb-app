package models

import "time"

type User struct {
    ID           string                   `json:"user_id"`
    Avatar       string                   `json:"avatar"`
    Mobile       string                   `json:"mobile"`
    Email        string                   `json:"email"`
    Name         string                   `json:"user_name"`
    Password     string                   `json:"password"`
    Gender       int8                     `json:"gender"`
    DepartmentID string                   `json:"department_id"`
    Status       int8                     `json:"status"`
    AdminFlag    int8                     `json:"admin_flag"`
    CreateAt     time.Time                `json:"create_at"`
    CreateUser   string                   `json:"create_user"`
    UpdateAt     time.Time                `json:"update_at"`
    UpdateUser   string                   `json:"update_user"`
    Role         []map[string]string      `json:"roles"`
    Button       []map[string]string      `json:"buttons"`
    Menu         []map[string]interface{} `json:"menus"`
}

type UserPassword struct {
    OldPassword string `json:"oldPassword"`
    NewPassword string `json:"newPassword"`
}

type UserLogin struct {
    LoginID  string `json:"loginID"`
    Password string `json:"password"`
}
