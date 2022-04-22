package models

import "time"

type Department struct {
    ID         string     `json:"id"`
    ParentID   string     `json:"parent_id"`
    Name       string     `json:"department_name"`
    Children   []*Menu    `json:"children,omitempty"`
    CreateAt   *time.Time `json:"create_at,omitempty"`
    CreateUser string     `json:"create_user,omitempty"`
    UpdateAt   *time.Time `json:"update_at,omitempty"`
    UpdateUser string     `json:"update_user,omitempty"`
}
