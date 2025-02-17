package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/ctype"
	"server/service/common/response"
	"server/service/user"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Param data body UserCreateRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	err := user.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return
}
