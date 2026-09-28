[简体中文](README.md) | [English](README.en.md)

# coypus 海狸鼠

![版本](https://img.shields.io/badge/release-0.1.0-blue.svg)
![语言](https://img.shields.io/badge/language-golang1.12-blue.svg)
![base](https://img.shields.io/badge/env-goframe1.7-red.svg)

> **⚠️ 本项目已停止开发！** 因长时间未对代码进行维护，可能会造成项目在不同环境上无法部署、运行 BUG 等问题，请知晓！项目仅供参考！

> 基于 goframe 框架的 Go Web 后端示例，完成后台管理系统基本组件开发。

## 项目介绍

coypus（海狸鼠）是一个结构清晰的 Go Web 后端入门示例（api / model / service 分层），实现了后台管理系统最核心的能力：JWT 登录认证、基于 casbin 的 RBAC 接口权限校验、用户/角色/菜单三套资源的增删改查。代码量不大，适合作为 Go Web 开发的学习模板，可配合前端项目 [coypus-vue](https://github.com/hequan2017/coypus-vue) 使用。

## ✨ 功能特性

- 登录认证：`/token` 签发 Token（jwt），其余接口统一校验；`/userInfo` 获取用户信息，`/menu` 下发菜单
- 权限验证：利用 casbin 库将 user / role / menu 自动关联。用户关联角色，角色关联菜单，权限关系为：
  - `角色(role.name, menu.path, menu.method)`
  - `用户(user.username, role.name)`
- 项目启动时自动加载权限；如有更改，会删除对应权限并重新加载
- 用户 `admin` 拥有所有权限，不进行权限匹配；登录接口 `/token` 不进行验证
- 用户 user、权限组 role、菜单 menu 的增删改查
- 请求和接收均传递 JSON 格式数据

例如：`test` 角色拥有 `test  /api/v1/users  GET` 权限，用户 `hequan` 属于 `test` 组，则 `hequan` 请求 `GET /api/v1/users` 时校验通过。

## 🛠 技术栈（版本取自 go.mod）

- Go 1.12
- goframe v1.9.1（Web 框架）
- gorm v1.9.10（ORM，MySQL）
- casbin v1.9.1（RBAC 权限模型）
- jwt-go v3.2.0（JWT 认证）、sha1（密码加密）

## 🚀 快速开始

* 部署 MySQL，创建库 `coypus`
* 导入 `docfile/sql/coypus.sql`
* 修改配置文件 `config/config.toml`（数据库连接、JwtSecret、PageSize 等）

```bash
go run main.go

2019/05/08 18:10:38.395 [INFO] 更新角色权限关系 [[hequan test]]
2019-05-08 18:10:38.397 16856: http server started listening on [:8000]
```

默认账户密码：`admin / 123456`

**请求示例**

```json
// 访问 /token 获取 token
{
    "username": "admin",
    "password": "123456"
}
```

```text
// 访问业务接口，请求头设置 Token
GET http://127.0.0.1:8000/api/v1/users?page=1
Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

**响应码说明**

```text
200：请求成功    201：创建、修改成功    204：删除成功    400：参数错误
401：未登录      403：禁止访问          404：未找到      500：系统错误
```

## 📁 目录结构

```text
- app      业务逻辑层  所有的业务逻辑存放目录。
    - api     业务接口  接收/解析用户输入参数的入口/接口层。
    - model   数据模型  数据管理层，仅用于操作管理数据，如数据库操作。
    - service 逻辑封装  业务逻辑封装层，实现特定的业务需求，可供不同的包调用。
- boot     初始化包  用于项目初始化参数设置。
- config   配置管理  所有的配置文件存放目录。
- docfile  项目文档  DOC项目文档，如: 设计文档、脚本文件等等。
- library  公共库包  公共的功能封装包，往往不包含业务需求实现。
- log      日志
- router   路由注册  用于路由统一的注册管理。
- test     单元测试
- go.mod   依赖管理  使用Go Module包管理的依赖描述文件。
- main.go  入口文件  程序入口文件。
```

## 🔗 相关项目

- 前端：[coypus-vue](https://github.com/hequan2017/coypus-vue)（未完成，仅完成基本登录）

## 📄 许可证

[MIT](LICENSE)

## 作者

* 何全
