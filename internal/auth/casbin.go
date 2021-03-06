package auth

import (
	"github.com/casbin/casbin/v2"
	"gopkg.in/mgo.v2/bson"
	"orse/internal/database"
	//"github.com/allegro/bigcache"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CasbinEnforcer *casbin.Enforcer

func init() {
	var err error
	a, _ := gormadapter.NewAdapterByDB(database.Db)
	CasbinEnforcer,err = casbin.NewEnforcer("config/casbin/casbin_model.ini", a)
	if err != nil {
		panic(any(err))
	}
	CasbinEnforcer.LoadPolicy()
}

type CasbinModel struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Uid    string        `json:"uid"`
	Ptype  string        `json:"ptype" bson:"ptype"`
	Role   string        `json:"role" bson:"v0"`
	Path   string        `json:"path" bson:"v1"`
	Method string        `json:"method" bson:"v2"`
}

//GetRolesForUser 获取用户具有的角色。
func (c *CasbinModel) GetRolesForUser() ([]string, error) {
	return CasbinEnforcer.GetRolesForUser("alice")
}

//GetUsersForRole 获取具有角色的用户。
func (c *CasbinModel) GetUsersForRole() ([]string, error) {
	return CasbinEnforcer.GetUsersForRole(c.Role)
}

//AddRoleForUser 为用户添加角色。 如果用户已经拥有该角色（aka不受影响），则返回false。
func (c *CasbinModel) AddRoleForUser()(bool, error)  {
	return CasbinEnforcer.AddRoleForUser(string(c.Uid), c.Role)
}

//DeleteRoleForUser 删除用户的角色。 如果用户没有该角色，则返回false。(不会产生影响)
func (c *CasbinModel) DeleteRoleForUser() (bool, error) {
	return CasbinEnforcer.DeleteRoleForUser("alice", "data1_admin")
}

//DeleteUser 删除一个用户。 如果用户不存在，则返回false（也就是说不受影响）。
func (c *CasbinModel) DeleteUser() (bool, error) {
	return CasbinEnforcer.DeleteUser("alice")
}

//DeleteRole 删除一个角色。
func (c *CasbinModel) DeleteRole() (bool, error) {
	return CasbinEnforcer.DeleteRole("data2_admin")
}

//DeletePermission 删除权限。 如果权限不存在，则返回false（也就是说不受影响）。
func (c *CasbinModel) DeletePermission() (bool, error) {
	return CasbinEnforcer.DeletePermission("read")
}

//AddPermissionForUser 为用户或角色添加权限。 如果用户或角色已经拥有该权限（也就是不受影响），则返回false。
func (c *CasbinModel) AddPermissionForUser() (bool, error) {
	return CasbinEnforcer.AddPermissionForUser("bob", "read")
}

//AddPermissionForUser 为用户或角色添加多个权限。 如果用户或角色已经有一个权限，则返回 false (不会受影响)。
func (c *CasbinModel) addPermissionForUser() (bool, error) {
	return CasbinEnforcer.AddPermissionsForUser("bob", []string{"read", "read"})

}

//DeletePermissionForUser 删除用户或角色的权限。 如果用户或角色没有权限则返回 false(不会受影响)。
func (c *CasbinModel) DeletePermissionForUser() (bool, error) {
	return CasbinEnforcer.DeletePermissionForUser("bob", "read")
}

//DeletePermissionsForUser 删除用户或角色的权限。 如果用户或角色没有任何权限（也就是不受影响），则返回false。
func (c *CasbinModel) DeletePermissionsForUser() (bool, error) {
	return CasbinEnforcer.DeletePermissionsForUser("bob")
}

//GetPermissionsForUser 获取用户或角色的权限。
func (c *CasbinModel) GetPermissionsForUser() [][]string {
	return CasbinEnforcer.GetPermissionsForUser("bob")
}

//HasPermissionForUser 确定用户是否具有权限。
func (c *CasbinModel) hasPermissionForUser() bool {
	return CasbinEnforcer.HasPermissionForUser(c.Role, c.Path, c.Method)
}
