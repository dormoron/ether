package users

import (
	"errors"
	"ether/internal/domain/model"
	"ether/internal/domain/service/users"
	"ether/internal/handler"
	"ether/pkg/logger"
	regexp "github.com/dlclark/regexp2"
	"github.com/dormoron/mist"
	"github.com/dormoron/mist/security"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

type AuthHandler interface {
	handler.Handler
	Login(ctx *mist.Context)
	SignUp(ctx *mist.Context)
	Profile(ctx *mist.Context)
	Logout(ctx *mist.Context)
}

type AuthHandlerStruct struct {
	svc         users.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	logger      logger.Logger
	security.Provider
}

func NewAuthHandler(svc users.UserService, sp security.Provider, logger logger.Logger) AuthHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &AuthHandlerStruct{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
		logger:      logger,
		Provider:    sp,
	}
}

func (u *AuthHandlerStruct) RegisterRoutes(server *mist.HTTPServer) {
	userGroup := server.Group("/auth")
	userGroup.POST("/signup", u.SignUp)
	userGroup.POST("/login", u.Login)
	userGroup.GET("/profile", u.Profile)
	userGroup.POST("/logout", u.Logout)
}

func (u *AuthHandlerStruct) Logout(ctx *mist.Context) {
	err := u.ClearToken(ctx)
	if err != nil {
		u.logger.Error("退出登录失败", logger.Field{
			Key:   "error",
			Value: err,
		})
		_ = ctx.RespondSuccess("退出登录失败")
		return
	}
	_ = ctx.RespondSuccess("退出登录成功")
}

func (u *AuthHandlerStruct) Login(ctx *mist.Context) {
	type Req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	auth, err := u.svc.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	_, err = u.InitSession(ctx, int64(auth.UserId), map[string]any{
		"authId": auth.Id,
	}, map[string]any{
		"authId": auth.Id,
	})
	if err != nil {
		u.logger.Error("登录失败", logger.Field{
			Key:   "error",
			Value: err,
		})
		return
	}
	_ = ctx.RespondSuccess("登录成功")
}

func (u *AuthHandlerStruct) SignUp(ctx *mist.Context) {
	var req SignUpReq

	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Username)
	if err != nil {
		_ = ctx.RespondSuccess("系统错误")
		return
	}
	if !ok {
		_ = ctx.RespondSuccess("你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		_ = ctx.RespondSuccess("两次输入的密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		_ = ctx.RespondSuccess("系统错误")
		return
	}
	if !ok {
		_ = ctx.RespondSuccess("密码必须大于8位，包含数字、特殊字符")
		return
	}

	err = u.svc.SignUp(ctx.Request.Context(), model.Auth{
		Username:   req.Username,
		Password:   req.Password,
		CreateTime: time.Now(),
	})
	if errors.Is(err, users.ErrUserDuplicateEmail) {
		// 这是复用
		span := trace.SpanFromContext(ctx.Request.Context())
		span.AddEvent("邮件冲突")
		_ = ctx.RespondSuccess("邮箱冲突")
		return
	}
	if err != nil {
		_ = ctx.RespondSuccess("系统异常")
		return
	}

	_ = ctx.RespondSuccess("注册成功")

}

func (u *AuthHandlerStruct) Profile(ctx *mist.Context) {
	session, err := u.Get(ctx)
	if err != nil {
		return
	}
	id, err := session.Get(ctx, "uId").AsUint()
	user, err := u.svc.Profile(ctx, id)
	if err != nil {
		return
	}
	_ = ctx.RespondSuccess(user)
}

func (u *AuthHandlerStruct) RefreshToken(ctx *mist.Context) {
	err := u.RenewAccessToken(ctx)
	if err != nil {
		_ = ctx.RespondWithJSON(http.StatusUnauthorized, "刷新错误")
	}
	_ = ctx.RespondSuccess("ok")
}
