# dockerproxy
Docker 镜像加速站集合

# Docker

环境变量参数：

|参数|描述|
|:--|:--|
|AUTO_UPDATE_TIME|设置自动更新时间/分钟|
|CRON_SPEC|定时任务 spec|
|GIST_URL|存放数据 gist 地址|

### 使用环境变量

```docker
docker run --name=dockerproxy -e GIST_URL="xxx" -p 9101:9101 -d ghcr.io/xiaoxuan6/dockerproxy:latest
```

### 使用 `.env` 文件

```docker
docker run --name=dockerproxy -v $(pwd)/.env:/src/.env -p 9101:9101 -d ghcr.io/xiaoxuan6/dockerproxy:latest
```
