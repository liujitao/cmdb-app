package models

// type Menu struct {
//     ID         string    `json:"id"`
//     ParentID   string    `json:"parent_id"`
//     Title      string    `json:"title"`
//     Perms      string    `json:"perms"`
//     Icon       string    `json:"icon"`
//     SortID     int8      `json:"sort_id"`
//     CreateAt   time.Time `json:"create_at,omitempty"`
//     CreateUser string    `json:"create_user,omitempty"`
//     UpdateAt   time.Time `json:"update_at,omitempty"`
//     UpdateUser string    `json:"update_user,omitempty"`
// }

type Menu struct {
    ID       string  `json:"id"`
    ParentID string  `json:"parent_id"`
    Title    string  `json:"title"`
    Perms    string  `json:"perms"`
    Icon     string  `json:"icon"`
    SortID   int8    `json:"sort_id"`
    Children []*Menu `json:"children,omitempty"`
}

type Button struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Resource string `json:"resources"`
}
