# id-maker
Go 开发的一款分布式唯一 ID 生成系统

## 使用

有两种方式来调用接口：

1. HTTP 方式
2. gRPC 方式

### HTTP 方式

1、健康检查：

```
curl http://127.0.0.1:8080/ping
```

2、获取 ID：

获取 tag 是 test 的 ID：

```
curl http://127.0.0.1:8080/v1/id/test
```

3、获取雪花 ID：

```
curl http://127.0.0.1:8080/v1/snowid
```

### gRPC 方式

1、获取 ID：

```
grpcurl -plaintext -d '{"tag":"test"}' -import-path $HOME/src/id-maker/internal/controller/rpc/proto -proto segment.proto localhost:50051 proto.Gid/GetId
```

2、获取雪花 ID：

```
grpcurl -plaintext -import-path $HOME/src/id-maker/internal/controller/rpc/proto -proto segment.proto localhost:50051 proto.Gid/GetSnowId
```

## 本地开发

```
# Run MySQL
$ make compose-up

# Run app with migrations
$ make run
```

## 推荐阅读

- [go-clean-template](https://github.com/evrone/go-clean-template)
- [hwholiday/gid](https://github.com/hwholiday/gid)
- [Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)