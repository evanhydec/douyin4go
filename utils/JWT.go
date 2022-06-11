package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"unsafe"
)

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

type Claims struct {
	ID        int `json:"user_id"`
	Authority int `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id int, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(effectTime).Unix()
	UserClaims := Claims{
		ID:        id,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "tiktok",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

// ParseToken 解析Token
func ParseToken(tokenString string) *Claims {
	if tokenString == "" {
		return &Claims{
			ID:             0,
			Authority:      1,
			StandardClaims: jwt.StandardClaims{},
		}
	}
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		LogrusObj.Info(err)
		return nil
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		//panic("token is valid")
		fmt.Println("token is valid")
		return nil
	}
	return claims
}

func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

func StringToByteSlice(s string) []byte {

	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))

	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}

	return *(*[]byte)(unsafe.Pointer(&tmp2))

}
