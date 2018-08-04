package models

import (
	"goAdmin/components"
	"goAdmin/connections/mysql"
	"strings"
)

func GetManagerTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "角色",
			Field:    "roles",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				labelModel, _ := mysql.Query("select r.name from goadmin_role_users as u left join goadmin_roles as r on "+
					"u.role_id = r.id where user_id = ?", model.ID)
				return components.Label.GetContent(labelModel[0]["name"].(string))
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "goadmin_users"
	userTable.Info.Title = "管理员管理"
	userTable.Info.Description = "管理员管理"

	var roles, permissions []map[string]string
	rolesModel, _ := mysql.Query("select `name`, `slug` from goadmin_roles where id > ?", 0)
	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": v["slug"].(string),
		})
	}
	permissionsModel, _ := mysql.Query("select `name`, `slug` from goadmin_permissions where id > ?", 0)
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": v["slug"].(string),
		})
	}

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "头像",
			Field:    "avatar",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "file",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "密码",
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "password",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "角色",
			Field:    "roles",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  roles,
			ExcuFun: func(model RowModel) interface{} {
				roleModel, _ := mysql.Query("select r.id, r.name, r.slug from goadmin_role_users as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", model.ID)
				var roles []string
				for _, v := range roleModel {
					roles = append(roles, v["slug"].(string))
				}
				return roles
			},
		}, {
			Head:     "权限",
			Field:    "permissions",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  permissions,
			ExcuFun: func(model RowModel) interface{} {
				permissionModel, _ := mysql.Query("select r.id, r.name, r.slug from goadmin_user_permissions as u left join goadmin_permissions as r on u.permission_id = r.id where user_id = ?", model.ID)
				var permissions []string
				for _, v := range permissionModel {
					permissions = append(permissions, v["slug"].(string))
				}
				return permissions
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "goadmin_users"
	userTable.Form.Title = "管理员管理"
	userTable.Form.Description = "管理员管理"

	return
}

func GetPermissionTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "http_method",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "http_path",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "goadmin_permissions"
	userTable.Info.Title = "权限管理"
	userTable.Info.Description = "权限管理"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "方法",
			Field:    "http_method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options: []map[string]string{
				{"value": "GET", "field": "GET"},
				{"value": "PUT", "field": "PUT"},
				{"value": "POST", "field": "POST"},
				{"value": "DELETE", "field": "DELETE"},
				{"value": "PATCH", "field": "PATCH"},
				{"value": "OPTIONS", "field": "OPTIONS"},
				{"value": "HEAD", "field": "HEAD"},
			},
			ExcuFun: func(model RowModel) interface{} {
				return strings.Split(model.Value, ",")
			},
		}, {
			Head:     "路径",
			Field:    "http_path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "textarea",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "goadmin_permissions"
	userTable.Form.Title = "权限管理"
	userTable.Form.Description = "权限管理"

	return
}

func GetRolesTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "goadmin_roles"
	userTable.Info.Title = "角色管理"
	userTable.Info.Description = "角色管理"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "goadmin_roles"
	userTable.Form.Title = "角色管理"
	userTable.Form.Description = "角色管理"

	return
}

func GetOpTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "goadmin_operation_log"
	userTable.Info.Title = "操作日志"
	userTable.Info.Description = "操作日志"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "goadmin_operation_log"
	userTable.Form.Title = "操作日志"
	userTable.Form.Description = "操作日志"

	return
}
