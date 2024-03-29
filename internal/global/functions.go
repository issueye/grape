package global

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/pkg/utils"
)

// CreateToken
// 创建 Token
func CreateToken(user *model.UserInfo) (signedToken string, success bool) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = utils.Struct2Json(user)
	expire := Auth.TimeFunc().Add(Auth.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = Auth.TimeFunc().Unix()
	tokenString, err := token.SignedString(Auth.Key)
	if err != nil {
		Log.Errorf("生成TOKEN失败，失败原因：%s", err.Error())
		return "", false
	}
	return fmt.Sprintf("%s %s", TokenHeadName, tokenString), true
}

// ParseToken
// 解析token
func ParseToken(token string) (*model.UserInfo, error) {
	t, err := Auth.ParseTokenString(token)
	if err != nil {
		return nil, err
	}

	claims := t.Claims.(jwt.MapClaims)
	user := new(model.UserInfo)
	utils.Json2Struct(claims["user"].(string), user)
	return user, nil
}
