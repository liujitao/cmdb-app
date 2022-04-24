package utils

import (
    "cmdb-app-mysql/models"

    "golang.org/x/crypto/bcrypt"
)

/*
明文加密
*/
func HashPassword(password string) string {
    bytePassword := []byte(password)
    passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    return string(passwordHash)
}

/*
验证密文
*/
func VerifyPassword(passwordHash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

/*
构建菜单树
*/
func BuildMenuTree(items []*models.Menu, parentID string) []*models.Menu {
    tree := make([]*models.Menu, 0)

    for _, item := range items {
        if item.ParentID == parentID {
            item.Children = BuildMenuTree(items, item.ID)
            tree = append(tree, item)
        }
    }

    return tree
}
