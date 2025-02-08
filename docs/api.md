# API

后端 API，所有 API 起始端点都为 `/v0/api`

### 请求

- URL: `/v0/api/account/login`
- 方法: `POST`
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