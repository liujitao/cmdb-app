package models

import "time"

type SimpleRole struct {
    ID   string `json:"id"`
    Name string `json:"role_name"`
}

type Role struct {
    ID          string     `json:"id"`
    Name        string     `json:"role_name"`
    Description string     `json:"description"`
    CreateAt    *time.Time `json:"create_at,omitempty"`
    CreateUser  string     `json:"create_user,omitempty"`
    UpdateAt    *time.Time `json:"update_at,omitempty"`
    UpdateUser  string     `json:"update_user,omitempty"`
    Permission  []string   `json:"permissions"`
}

type RoleResponse struct {
    ID          string             `json:"id"`
    Name        string             `json:"role_name"`
    Description string             `json:"description"`
    CreateAt    *time.Time         `json:"create_at,omitempty"`
    CreateUser  string             `json:"create_user,omitempty"`
    UpdateAt    *time.Time         `json:"update_at,omitempty"`
    UpdateUser  string             `json:"update_user,omitempty"`
    User        []SimpleUser       `json:"users"`
    Permission  []SimplePermission `json:"permissions"`
}
