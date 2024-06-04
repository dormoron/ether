package users

import (
	"errors"
	"ether/internal/domain/model"
	"ether/internal/domain/service/users"
	regexp "github.com/dlclark/regexp2"
	"github.com/dormoron/mist"
	"go.opentelemetry.io/otel/trace"
)

const biz = "login"

type AuthHandler struct {
	svc         users.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewAuthHandler(svc users.UserService) *AuthHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &AuthHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *AuthHandler) SignUp(ctx *mist.Context) {
	var req SignUpReq

	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Username)
	if err != nil {
		_ = ctx.RespJSONOK("系统错误")
		return
	}
	if !ok {
		_ = ctx.RespJSONOK("你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		_ = ctx.RespJSONOK("两次输入的密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		_ = ctx.RespJSONOK("系统错误")
		return
	}
	if !ok {
		_ = ctx.RespJSONOK("密码必须大于8位，包含数字、特殊字符")
		return
	}

	err = u.svc.SignUp(ctx.Request.Context(), model.Auth{
		Username: req.Username,
		Password: req.Password,
	})
	if errors.Is(err, users.ErrUserDuplicateEmail) {
		// 这是复用
		span := trace.SpanFromContext(ctx.Request.Context())
		span.AddEvent("邮件冲突")
		_ = ctx.RespJSONOK("邮箱冲突")
		return
	}
	if err != nil {
		_ = ctx.RespJSONOK("系统异常")
		return
	}

	_ = ctx.RespJSONOK("注册成功")

}