package models

import "time"

type SimplePermission struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}

type Permission struct {
    ID         string     `json:"id"`
    ParentID   string     `json:"parent_id"`
    Title      string     `json:"title"`
    Name       string     `json:"name"`
    Path       string     `json:"path"`
    Component  string     `json:"component"`
    Icon       string     `json:"icon"`
    Redirect   string     `json:"redirect"`
    Type       int8       `json:"permission_type"`
    SortID     int8       `json:"sort_id"`
    CreateAt   *time.Time `json:"create_at,omitempty"`
    CreateUser string     `json:"create_user,omitempty"`
    UpdateAt   *time.Time `json:"update_at,omitempty"`
    UpdateUser string     `json:"update_user,omitempty"`
}

type PermissionResponse struct {
    ID         string     `json:"id"`
    ParentID   string     `json:"parent_id"`
    Title      string     `json:"title"`
    Name       string     `json:"name"`
    Path       string     `json:"path"`
    Component  string     `json:"component"`
    Icon       string     `json:"icon"`
    Redirect   string     `json:"redirect"`
    Type       int8       `json:"permission_type"`
    SortID     int8       `json:"sort_id"`
    CreateAt   *time.Time `json:"create_at,omitempty"`
    CreateUser string     `json:"create_user,omitempty"`
    UpdateAt   *time.Time `json:"update_at,omitempty"`
    UpdateUser string     `json:"update_user,omitempty"`
}

type PermissionTree struct {
    ID         string            `json:"id"`
    ParentID   string            `json:"parent_id"`
    Title      string            `json:"title"`
    Name       string            `json:"name,omitempty"`
    Path       string            `json:"path,omitempty"`
    Component  string            `json:"component,omitempty"`
    Icon       string            `json:"icon,omitempty"`
    Redirect   string            `json:"redirect,omitempty"`
    Type       int8              `json:"permission_type"`
    SortID     int8              `json:"sort_id,omitempty"`
    CreateAt   *time.Time        `json:"create_at,omitempty"`
    CreateUser *string           `json:"create_user,omitempty"`
    UpdateAt   *time.Time        `json:"update_at,omitempty"`
    UpdateUser *string           `json:"update_user,omitempty"`
    Children   []*PermissionTree `json:"children,omitempty"`
}

type MenuTree struct {
    ID        string      `json:"id"`
    ParentID  string      `json:"parent_id"`
    Name      string      `json:"name"`
    Path      string      `json:"path"`
    Component string      `json:"component"`
    Meta      *Meta       `json:"meta,omitempty"`
    Redirect  string      `json:"redirect"`
    SortID    int8        `json:"sort_id"`
    Children  []*MenuTree `json:"children,omitempty"`
}

type Meta struct {
    Title string `json:"title,omitempty"`
    Icon  string `json:"icon,omitempty"`
}

type Button struct {
    ID    string `json:"id"`
    Title string `json:"title"`
    Path  string `json:"path"`
}
