package utils

import (
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
// func BuildMenuTree(menus []*models.Menu, ParentID string) []*models.Menu {
//     result := []*models.Menu{}
//     for _, v := range menus {
//         if v.ParentID == ParentID {
//             v.Children = BuildMenuTree(v, v.ID)
//             result = append(result, v)
//         }
//     }
//     return result
// }
