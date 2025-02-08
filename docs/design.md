# Design

API 接口规范

使用响应体中的 status 字段设置真实的状态码。所有数据放在 data 字段中

```js
{
  status: 200, // 真实的状态码
  msg: '获取 Token 成功', // 警惕 AI 插件，插件会把 “msg” 写成 “message”
  data: {
      token: 'xxx',
      exp: 1145141919810
  }
}
```

```js
{
  status: 403,
  msg: '鉴权失败',
  data: {} // data 数据为空时可不传
}
```

