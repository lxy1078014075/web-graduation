package jwt

import (
	"errors"
	"strconv"
	"time"
	"web-graduation/dao/redis"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 24 * 365

var MySecret = []byte("Liang XinYuan")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	PositionId int64 `json:"identity"`
	jwt.StandardClaims
}

// GetToken 生成JWT
func GetToken(userID int64, positionId int64) (string, error) {
	c := MyClaims{
		userID,
		positionId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "LXY-subject",
		},
	}
	// 使用指定的签名方法生成签名对象
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)
	if err != nil {
		return "", err
	}
	// 将userId：token 存储在redis中，用来判断当前的token和redis存储的是否相同，来判断是否是同一个用户在使用
	err = redis.SetKey(strconv.FormatInt(userID, 10), token)

	return token, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	var c = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return c, nil
	}
	return nil, errors.New("invalid token")
}
