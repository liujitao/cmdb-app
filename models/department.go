package models

import "time"

type SimpleDepartment struct {
    ID   string `json:"id"`
    Name string `json:"department_name"`
}

type Department struct {
    ID          string       `json:"id"`
    ParentID    string       `json:"parent_id,omitempty"`
    Name        string       `json:"department_name"`
    Description string       `json:"description"`
    CreateAt    *time.Time   `json:"create_at,omitempty"`
    CreateUser  string       `json:"create_user,omitempty"`
    UpdateAt    *time.Time   `json:"update_at,omitempty"`
    UpdateUser  string       `json:"update_user,omitempty"`
    User        []SimpleUser `json:"users"`
}

type DepartmentTree struct {
    ID       string            `json:"id"`
    ParentID string            `json:"parent_id"`
    Name     string            `json:"department_name"`
    Children []*DepartmentTree `json:"children,omitempty"`
}
