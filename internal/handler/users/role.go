package users

import (
	"ether/internal/domain/model"
	"ether/internal/domain/service/users"
	"ether/internal/handler"
	"github.com/dormoron/mist"
	"time"
)

type RoleHandler interface {
	handler.Handler
	Insert(ctx *mist.Context)
	Update(ctx *mist.Context)
	Delete(ctx *mist.Context)
	Details(ctx *mist.Context)
}

type RoleHandlerStruct struct {
	svc users.RoleService
}

func (r RoleHandlerStruct) RegisterRoutes(server *mist.HTTPServer) {
	group := server.Group("/role")
	group.GET("/detail", r.Details)
	group.POST("/insert", r.Insert)
	group.POST("/delete", r.Delete)
	group.POST("/update", r.Update)
}

func (r RoleHandlerStruct) Insert(ctx *mist.Context) {
	type roleReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var req roleReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	insert, err := r.svc.Insert(ctx, model.Role{
		Name:        req.Name,
		Description: req.Description,
		CreateTime:  time.Now(),
	})
	if err != nil {
		return
	}
	_ = ctx.RespondSuccess(insert)
}

func (r RoleHandlerStruct) Update(ctx *mist.Context) {
	type roleReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var req roleReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	err := r.svc.UpdateById(ctx, model.Role{
		Name:        req.Name,
		Description: req.Description,
		UpdateTime:  time.Now(),
	})
	if err != nil {
		return
	}
	_ = ctx.RespondSuccess("修改成功")
}

func (r RoleHandlerStruct) Delete(ctx *mist.Context) {
	type roleReq struct {
		Id uint `json:"id"`
	}
	var req roleReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	err := r.svc.Delete(ctx, req.Id)
	if err != nil {
		return
	}
	_ = ctx.RespondSuccess("删除成功")
}

func (r RoleHandlerStruct) Details(ctx *mist.Context) {
	id, err := ctx.QueryValue("id").AsUint()
	if err != nil {
		return
	}
	detail, err := r.svc.Detail(ctx, id)
	if err != nil {
		return
	}
	_ = ctx.RespondSuccess(detail)
}

func NewRoleHandler(svc users.RoleService) RoleHandler {
	return &RoleHandlerStruct{
		svc: svc,
	}
}
