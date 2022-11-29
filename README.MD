# go-oss

模拟阿里云 OSS 服务,支持前端签名上传.

DockerHub: https://hub.docker.com/r/gcslaoli/go-oss

Github: https://github.com/gcslaoli/go-oss

## 拉取镜像

```bash
docker pull gcslaoli/go-oss
```

## 配置文件

```yaml
#  config.yaml
server:
  address: ":8000" # 服务监听地址
  clientMaxBodySize: 810241024 # 客户端最大Body上传限制大小，影响文件上传大小(Byte)。默认为8*1024*1024=8MB
logger:
  level: "all" # 日志级别
  stdout: true # 是否输出到控制台

oss:
  allowAnonymous: true # 是否允许匿名上传,开启后不需要签，但上传的文件将会被替换为随机文件名
  baseUrl: "https://oss.dev.lidong.xin" # OSS访问域名
  keySecrets:
    yourKey: "your-secrets" # key 和 secret 用于前端签名上传, 请自行修改,支持多个
    key2: "secrets2"
```

## docker-compose

```yaml
# docker-compose.yaml
version: "3"
services:
  go-oss:
    image: gcslaoli/go-oss:latest
    restart: always
    ports:
      - 8000:8000
    volumes:
      - /data/go-oss:/app/upload
      - config.yaml:/app/config.yaml
```
