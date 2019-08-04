# cat-slave
wechat public 

<!-- 格式化代码 -->
gofmt -w .   
<!-- 检查静态错误 -->
go tool vet .
<!-- 编译 -v打印包名 -->
go build -v .

项目路径
/data/cat-slave
docker路径
/data/docker-config

目录和文章都放在微信上

go项目只负责全文索引
全文索引存在es中

架构
    router 