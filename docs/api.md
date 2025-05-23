# API

后端 API，所有 API 起始端点都为 `/v0/api`

## 用户登录

### 请求
- URL: `/account/auth/login`
- 方法: `POST`
- 查询参数：无
- 请求体:

```json
{
    "username": "string",
    "password": "string",
    "turnstile_secretkey": "string"
}
```

### 响应

#### 成功

```json
{
    "status": 200,
    "msg": "获取 Token 成功",
    "data": {
        "token": "string",
        "exp": 86400
    }
}
```

#### 用户不存在或密码不正确
```json
{
      "status": 404,
      "msg": "用户不存在或密码不正确",
}
```

#### 后端/数据库 抛出了奇怪的错误
```json
{
    "status": 500,
    "msg": "对不起，线路依然繁忙，请再等一下，或者稍后再打过来",
}
```



## 发送验证码

### 请求
- URL: `/account/sendcode`

- 方法: `POST`

- 查询参数：usefor

  目前仅支持 register 和 reset_password，以后如果需要鉴权需要带上 JWT

  | 参数           | 效果               |
  | -------------- | ------------------ |
  | register       | 发送注册验证码     |
  | reset_password | 发送重置密码验证码 |
- 请求体:

```json
{
    "useremail": "string",
    "turnstile_secretkey": "string"
}
```

### 响应

#### 成功

```json
{
    "status": 200,
    "msg": "发送验证码成功! 请注意查收~"
}
```

#### 后端/数据库 抛出了奇怪的错误

```json
{
    "status": 500,
    "msg": "对不起，线路依然繁忙，请再等一下，或者稍后再打过来",
}
```



## 用户注册

### 请求

- URL: `/account/auth/register`

- 方法: POST

- 查询参数: 无

- 请求体:

  ```json
  {
      "username": "string",
      "password": "string",
      "useremail": "string",
      "code": "string",
      "turnstile_secretkey": "string"
  }
  ```

  

### 响应

#### 成功

```json
{
    "status": 200,
    "msg": "用户注册成功"
}
```

#### 验证码错误

```json
{
    "status": 400,
    "msg": "验证码错误"
}
```



# 以后的 API 文档都使用 ApiFox 一键生成

# 因为我好懒qnq

