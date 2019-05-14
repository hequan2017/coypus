package e

var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "请求参数错误",

	ERROR_USER_EXIST:      "已存在该用户名称",
	ERROR_USER_NOT_EXIST:  "该用户不存在",
	ERROR_USER_EXIST_FAIL: "获取已存在用户失败",

	ERROR_USER_GET_S_FAIL:  "获取所有用户失败",
	ERROR_USER_ADD_FAIL:    "新增用户失败",
	ERROR_USER_EDIT_FAIL:   "修改用户失败",
	ERROR_USER_DELETE_FAIL: "删除用户失败",

	ERROR_ROLE_EXIST:      "已存在该权限组名称",
	ERROR_ROLE_NOT_EXIST:  "该权限组不存在",
	ERROR_ROLE_EXIST_FAIL: "获取已存在权限组失败",

	ERROR_ROLE_GET_S_FAIL:  "获取所有权限组失败",
	ERROR_ROLE_ADD_FAIL:    "新增权限组失败",
	ERROR_ROLE_EDIT_FAIL:   "修改权限组失败",
	ERROR_ROLE_DELETE_FAIL: "删除权限组失败",

	ERROR_MENU_EXIST:      "已存在该菜单名称",
	ERROR_MENU_NOT_EXIST:  "该菜单不存在",
	ERROR_MENU_EXIST_FAIL: "获取已存在菜单失败",

	ERROR_MENU_GET_S_FAIL:  "获取所有菜单失败",
	ERROR_MENU_ADD_FAIL:    "新增菜单失败",
	ERROR_MENU_EDIT_FAIL:   "修改菜单失败",
	ERROR_MENU_DELETE_FAIL: "删除菜单失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
	ERROR_LOAD_CASBIN_FAIL:      "加载用户权限失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
