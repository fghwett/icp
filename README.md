# ICP备案信息查询API接口

项目为 [https://github.com/yitd/ICP-API](https://github.com/yitd/ICP-API) go语言版本，支持独立部署及SDK使用

采用 [管局官网](https://beian.miit.gov.cn/#/Integrated/recordQuery) 备案接口，同步最新ICP备案数据

**主要功能：**
- [x] 通过主域名查询
- [ ] 通过子域查询
- [ ] 通过链接🔗 查询
- [ ] 通过备案号反差域名
- [ ] 通过注册人（注册单位）反差域名

## 开发流程
查询备案信息主要流程如下：
调用 `auth` 接口获取token，利用token调用 `icpAbbreviateInfo/queryByCondition`接口查询相关信息

> `icpAbbreviateInfo/queryByCondition`该接口可以实现所有主要功能

另外，管局官方开放了`icpAbbreviateInfo/queryDetailByServiceIdAndDomainId`接口，可以查询详细信息，但是通过数据对比，该接口与上一个接口提供的信息完全一致

甚至上一个接口可以得知域名是否被禁止访问，所以这个接口就没有使用（搞不清楚为什么做这个接口）

欢迎有能力者为其开发其他语言的版本

## 接口
接口域名地址：`https://hlwicpfwc.miit.gov.cn/icpproject_query/api/`
接口通用header：
```shell
"Content-Type: $Content", // 根据参数传入
"Origin: https://beian.miit.gov.cn/",
"Referer: https://beian.miit.gov.cn/",
"token: $token", // 根据参数传入
"User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36", // 可以获取自己浏览器的
"CLIENT-IP: $ip", // 建议随机生成 同一个IP会有访问频率限制
"X-FORWARDED-FOR: $ip" // 获取token和请求域名信息的IP可以不同 但是CLIENT-IP和X-FORWARDED-FOR的IP一定要一样
```

### 获取token
接口地址：`auth`

```shell
header -> "token": "0"
header -> "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8"

body -> authKey=md5("testtest" . timestamp)&timeStamp=timestamp
timestamp精确到秒

```

**接口返回值：**

失败-header中两个ip不同时出现
```json
{
  "success": false,
  "code": 500,
  "msg": "服务器异常"
}
```

成功
```json
{
  "code": 200,
  "msg": "操作成功",
  "params": {
    "bussiness": "eyJ0eXBlIjoxLCJ1IjoiMDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjYiLCJzIjoxNjMzOTUyOTI4NzIxLCJlIjoxNjMzOTUzNDA4NzIxfQ.9jJWrc2L1IwD4I_vs8p9O0oFFG6RUjIqda5Ubz2nZn4",
    "expire": 300000,
    "refresh": "eyJ0eXBlIjoyLCJ1IjoiMDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjYiLCJzIjoxNjMzOTUyOTI4NzIxLCJlIjoxNjMzOTUzNzA4NzIxfQ.r1vTT-MN3EquWVdshOlehr7caK4X2D59FAz3vjZjkNc"
  },
  "success": true
}
```
`params.bussiness`即获取到的token，通过base64解码token第一段参数，既可获取到实际token的过期时间

1. 将token第一段补齐到base64标准格式
    ```shell
    eyJ0eXBlIjoxLCJ1IjoiMDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjYiLCJzIjoxNjMzOTUyOTI4NzIxLCJlIjoxNjMzOTUzNDA4NzIxfQ==
    ```
2. 解码
    ```json
    {
        "type": 1,
        "u": "098f6bcd4621d373cade4e832627b4f6",
        "s": 1633952928721,
        "e": 1633953408721
    }
    ```
s为生效时间e为过期时间，时间很短只有8分钟。但是根据接口中的refresh接口得知，token应该是可以刷新的，但是暂时还没有找到该接口

### 根据域名、备案号、企业名等信息查询
接口地址：`icpAbbreviateInfo/queryByCondition`

```shell
header -> "token": 上个接口中获取到的token
header -> "Content-Type": "application/json;charset=UTF-8"

body -> 
{
    "pageNum": "1", // 可为空
    "pageSize": "10", // 可为空
    "unitName": "mi.cn" // 必填 要查询的域名备案号等
}
```

**接口返回值：**
```json
{
   "code": 200,
   "msg": "操作成功",
   "params": {
       "endRow": 0,
       "firstPage": 1,
       "hasNextPage": false,
       "hasPreviousPage": false,
       "isFirstPage": true,
       "isLastPage": true,
       "lastPage": 1,
       "list": [
           {
               "contentTypeName": "",
               "domain": "mi.cn",
               "domainId": 10004219290,
               "homeUrl": "www.mi.cn",
               "leaderName": "",
               "limitAccess": "否",
               "mainId": 6504864,
               "mainLicence": "京ICP备10046444号",
               "natureName": "企业",
               "serviceId": 10001392064,
               "serviceLicence": "京ICP备10046444号-9",
               "serviceName": "小米科技",
               "unitName": "小米科技有限责任公司",
               "updateRecordTime": "2021-08-16 13:55:56"
           }
       ],
       "navigatePages": 8,
       "navigatepageNums": [
           1
       ],
       "nextPage": 1,
       "pageNum": 1,
       "pageSize": 10,
       "pages": 1,
       "prePage": 1,
       "size": 1,
       "startRow": 0,
       "total": 1
   },
   "success": true
}
```

### 获取域名详细信息接口
接口地址：`icpAbbreviateInfo/queryDetailByServiceIdAndDomainId`

```shell
header -> "token": 上个接口中获取到的token
header -> "Content-Type": "application/json;charset=UTF-8"

body -> 
{
   "mainId": 6504864,
   "domainId": 10004219290,
   "serviceId": 10001392064
}
// 以上三个参数全部必填 从上个接口中获取
```

**接口返回值：**
```shell
{
   "code": 200,
   "msg": "操作成功",
   "params": {
       "contentTypeName": "",
       "domain": "mi.cn",
       "domainId": 10004219290,
       "homeUrl": "www.mi.cn",
       "leaderName": "",
       "mainId": 6504864,
       "mainLicence": "京ICP备10046444号",
       "natureName": "企业",
       "serviceId": 10001392064,
       "serviceLicence": "京ICP备10046444号-9",
       "serviceName": "小米科技",
       "unitName": "小米科技有限责任公司",
       "updateRecordTime": "2021-08-16 13:55:56"
   },
   "success": true
}
```