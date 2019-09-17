package permission

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/hequan2017/coypus/library/e"
	"github.com/hequan2017/coypus/library/inject"
	jwtGet "github.com/hequan2017/coypus/library/jwt"
	"net/http"
	"strings"
)

func CasbinMiddleware(r *ghttp.Request) {
	Authorization := r.Header.Get("Authorization")
	token := strings.Split(Authorization, " ")
	t, _ := jwt.Parse(token[1], func(*jwt.Token) (interface{}, error) {
		return jwtGet.JwtSecret(), nil
	})
	glog.Info("-----权限验证-----", jwtGet.GetIdFromClaims("username", t.Claims), r.Request.URL.Path, r.Request.Method)

	if b, err := inject.Obj.Enforcer.EnforceSafe(jwtGet.GetIdFromClaims("username", t.Claims), r.Request.URL.Path, r.Request.Method); err != nil {
		_ = r.Response.WriteJson(g.Map{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
			"data": nil,
		})
		r.ExitAll()
	} else if !b {
		_ = r.Response.WriteJson(g.Map{
			"code": http.StatusForbidden,
			"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
			"data": nil,
		})
		r.ExitAll()
	}
}
