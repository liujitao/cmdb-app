package models

type Response struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Error   string                 `json:"error,omitempty"`
    Data    map[string]interface{} `json:"data,omitempty"`
}

type ResponseList struct {
    Status  int                      `json:"status"`
    Message string                   `json:"message"`
    Error   string                   `json:"error,omitempty"`
    Data    []map[string]interface{} `json:"data,omitempty"`
}
