package models

type SimpleServer struct {
    ID   string `json:"id"`
    Name string `json:"server_name"`
}

type SimpleNetwork struct {
    ID   string `json:"id"`
    Name string `json:"network_name"`
}

type SimpleNetstorage struct {
    ID   string `json:"id"`
    Name string `json:"storage_name"`
}
