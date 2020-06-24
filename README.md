# Coj Specialist API

##### Golang v.1.14

Install:

```
git clone https://github.com/judascrow/cojspcl-api.git
cp .env.example .env
```

## Features

- RESTful API
- Gin
- Gorm
- Gin-Swagger
- Log-Zap
- Gin-Jwt
- Casbin

## COMMAND

Migrate & Seed :

```
go run main.go create seed
```

Run :

```
go run main.go
```

## API ROUTER URI

```
Host: localhost:8000
Base Path: /api/v1
```

### Healthcheck

| Done               | Method | URI          | authorize | comment              |
| ------------------ | ------ | ------------ | --------- | -------------------- |
| :white_check_mark: | GET    | /healthcheck | No        | Check Status Service |

### Auth

| Done               | Method | URI            | authorize | comment  |
| ------------------ | ------ | -------------- | --------- | -------- |
| :white_check_mark: | POST   | /auth/login    | No        | Log in   |
| :white_check_mark: | POST   | /auth/register | No        | Register |
| :white_check_mark: | GET    | /auth/me       | Yes       | Me User  |

### User

| Done               | Method | URI                   | authorize | comment          |
| ------------------ | ------ | --------------------- | --------- | ---------------- |
| :white_check_mark: | GET    | /users                | Yes       | List Users       |
| :white_check_mark: | GET    | /users/:slug          | Yes       | Get User by Slug |
| :white_check_mark: | POST   | /users                | Yes       | Create User      |
| :white_check_mark: | PUT    | /users/:slug          | Yes       | Update User      |
| :white_check_mark: | DELETE | /users/:slug          | Yes       | Delete User      |
| :white_check_mark: | POST   | /users/:slug/password | Yes       | Change Password  |
| :white_check_mark: | POST   | /users/:slug/avatar   | Yes       | Upload Avatar    |

### Role

| Done               | Method | URI    | authorize | comment    |
| ------------------ | ------ | ------ | --------- | ---------- |
| :white_check_mark: | GET    | /roles | Yes       | List Roles |

### Swaggo

```
Generate command: swag init

URL: http://localhost:8000/api/v1/swagger/index.html
```

### File Server

```

URL: http://localhost:8000/upload
```

## FOR DEV

### GIT

Update

```
./git.sh "COMMENT"
```
