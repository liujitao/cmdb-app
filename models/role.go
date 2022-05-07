package models

import "time"

type Role struct {
    ID         string     `json:"id"`
    Name       string     `json:"role_name"`
    CreateAt   *time.Time `json:"create_at,omitempty"`
    CreateUser string     `json:"create_user,omitempty"`
    UpdateAt   *time.Time `json:"update_at,omitempty"`
    UpdateUser string     `json:"update_user,omitempty"`
}
