[简体中文](README.md) | [English](README.en.md)

# coypus

![Version](https://img.shields.io/badge/release-0.1.0-blue.svg)
![Language](https://img.shields.io/badge/language-golang1.12-blue.svg)
![Base](https://img.shields.io/badge/env-goframe1.7-red.svg)

> **⚠️ Development of this project has been discontinued!** The code has not been maintained for a long time, so it may fail to deploy or run in different environments. Please be aware! For reference only!

> A Go web back-end example based on the goframe framework, implementing the basic components of an admin system.

## Introduction

coypus is a well-structured Go web back-end example (layered as api / model / service) that implements the core capabilities of an admin system: JWT login authentication, casbin-based RBAC API permission checks, and CRUD for users, roles and menus. The codebase is small and works well as a learning template for Go web development, and can be paired with the [coypus-vue](https://github.com/hequan2017/coypus-vue) front-end.

## ✨ Features

- Authentication: `/token` issues a JWT, all other endpoints require it; `/userInfo` returns user info, `/menu` delivers the menu
- Permission checks: the casbin library links user / role / menu automatically. Users are bound to roles and roles to menus:
  - `role(role.name, menu.path, menu.method)`
  - `user(user.username, role.name)`
- Permissions are loaded automatically at startup; on change, the affected policies are removed and reloaded
- The `admin` user has all permissions and bypasses matching; the login endpoint `/token` is not verified
- CRUD for users, roles (permission groups) and menus
- Requests and responses are both JSON

Example: the `test` role holds `test /api/v1/users GET`; since user `hequan` belongs to the `test` group, a `GET /api/v1/users` request from `hequan` passes the check.

## 🛠 Tech Stack (versions from go.mod)

- Go 1.12
- goframe v1.9.1 (web framework)
- gorm v1.9.10 (ORM, MySQL)
- casbin v1.9.1 (RBAC permission model)
- jwt-go v3.2.0 (JWT authentication), sha1 (password hashing)

## 🚀 Quick Start

- Deploy MySQL and create the `coypus` database
- Import `docfile/sql/coypus.sql`
- Edit the config file `config/config.toml` (database connection, JwtSecret, PageSize, etc.)

```bash
go run main.go

2019/05/08 18:10:38.395 [INFO] 更新角色权限关系 [[hequan test]]
2019-05-08 18:10:38.397 16856: http server started listening on [:8000]
```

Default account: `admin / 123456`

**Request example**

```json
// POST /token to obtain a token
{
    "username": "admin",
    "password": "123456"
}
```

```text
// Call a business API with the token header
GET http://127.0.0.1:8000/api/v1/users?page=1
Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

**Response codes**

```text
200: success    201: created / updated    204: deleted    400: bad request
401: not logged in    403: forbidden    404: not found    500: internal error
```

## 📁 Directory Structure

```text
- app      business logic layer, holding all business logic
    - api     API layer, entry point that receives/parses user input
    - model   data model layer, only for data access such as database operations
    - service service layer, encapsulates business logic for reuse across packages
- boot     initialization package for project setup
- config   configuration directory
- docfile  project documents, e.g. design docs and script files
- library  common/shared library packages, usually business-free
- log      logs
- router   unified route registration
- test     unit tests
- go.mod   Go Module dependency file
- main.go  program entry
```

## 🔗 Related Projects

- Front-end: [coypus-vue](https://github.com/hequan2017/coypus-vue) (unfinished, only basic login completed)

## 📄 License

[MIT](LICENSE)

## Author

* He Quan (何全)
