package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/hequan2017/coypus/library/e"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func JwtSecret() []byte {
	return []byte(g.Config().GetString("setting.PageSize"))
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "https://github.com/hequan2017/go-admin/",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret())
	return token, err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 根据 token 获取用户名
func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}

func JWT(r *ghttp.Request) {
	if r.URL.Path == "/token" {
		r.Exit()
	} else {
		Authorization := r.Header.Get("Authorization")
		token := strings.Split(Authorization, " ")
		if Authorization == "" {
			_ = r.Response.WriteJson(g.Map{
				"code": http.StatusForbidden,
				"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
				"data": nil,
			})
			r.ExitAll()
		} else {
			_, err := ParseToken(token[1])
			if err != nil {
				_ = r.Response.WriteJson(g.Map{
					"code": http.StatusForbidden,
					"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
					"data": nil,
				})
				r.ExitAll()
			}
		}
		return
	}

}
