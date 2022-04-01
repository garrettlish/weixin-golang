# wxcloudrun-golang
[![GitHub license](https://img.shields.io/github/license/WeixinCloud/wxcloudrun-express)](https://github.com/WeixinCloud/wxcloudrun-express)
![GitHub package.json dependency version (prod)](https://img.shields.io/badge/golang-1.17.1-green)

微信云托管 golang 模版，实现简单的计数器读写接口，使用云托管 MySQL 读写、记录计数值。

![](https://qcloudimg.tencent-cloud.cn/raw/be22992d297d1b9a1a5365e606276781.png)


## 快速开始
前往 [微信云托管快速开始页面](https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/basic/guide.html)，选择相应语言的模板，根据引导完成部署。


## 目录结构说明
~~~
.
├── Dockerfile                Dockerfile 文件
├── LICENSE                   LICENSE 文件
├── README.md                 README 文件
├── container.config.json     微信云托管流水线配置
├── db                        数据库逻辑目录
├── go.mod                    go.mod 文件
├── go.sum                    go.sum 文件
├── index.html                主页 html 
├── main.go                   主函数入口
└── service                   接口服务逻辑目录
~~~


## 服务 API 文档

### `GET /api/stock`

获取当前股票价格

#### 请求参数

无

#### 响应结果

- `code`：错误码
- `data`：当前股票价格

##### 响应结果示例

```json
{
  "code": 0,
  "data": 42
}
```

#### 调用示例

```
curl https://<云托管服务域名>/api/stock
```


## License

[MIT](./LICENSE)
