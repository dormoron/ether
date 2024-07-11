package users

import (
	"ether/internal/handler"
	"ether/pkg/casbin"
	"github.com/dormoron/mist"
)

var _ handler.Handler = (*CasbinHandlerStruct)(nil)

type CasbinHandler interface {
	Insert(ctx *mist.Context)
	InsertRole(ctx *mist.Context)
	UpdateRoleForUser(ctx *mist.Context)
	Update(ctx *mist.Context)
	Delete(ctx *mist.Context)
	DeleteRole(ctx *mist.Context)
}

type CasbinHandlerStruct struct {
	access casbin.Access
}

func (c *CasbinHandlerStruct) RegisterRoutes(server *mist.HTTPServer) {
	accessGroup := server.Group("/access")
	accessGroup.POST("/create", c.Insert)
}

func (c *CasbinHandlerStruct) Insert(ctx *mist.Context) {
	var req AccessRequest
	policy, err := c.access.AddPolicy(req.Subject, req.Object, req.Action)
	if err != nil && !policy {
		return
	}
	_ = ctx.RespondSuccess("success")
}

func (c *CasbinHandlerStruct) UpdateRoleForUser(ctx *mist.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *CasbinHandlerStruct) Update(ctx *mist.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *CasbinHandlerStruct) Delete(ctx *mist.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *CasbinHandlerStruct) DeleteRole(ctx *mist.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *CasbinHandlerStruct) InsertRole(ctx *mist.Context) {

}

func NewCasbinHandler(access casbin.Access) CasbinHandler {
	return &CasbinHandlerStruct{
		access: access,
	}
}
