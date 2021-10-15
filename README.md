# go lambda
serverless project  golang + aws lambda  + aws dynamodb

# 启动项目

参照[Developer Guide](https://docs.aws.amazon.com/zh_cn/sdk-for-go/v1/developer-guide/configuring-sdk.html)、[aws-sdk-go](https://github.com/aws/aws-sdk-go#configuring-credentials)配置安全凭证到本地环境

```bash
$ go mod tidy
```

设置环境变量 dev sls prod
```bash
$ export GOENV_MODE=prod
```

发布到serverless
```bash
$ make deploy
```

部署到服务器
```bash
$ make publish
# 将bin下的文件上传到服务器
# 用nohub 启动
$ nohub ./main &
# 关闭服务
$ ps -aux | grep "./mian"
# 找到PID杀死进程
$ kill 1
```