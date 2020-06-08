## 流媒体网站

### 简介

这是一个流媒体网站，主要实现功能：
- 播放视频，上传视频
- 用户登录，管理自己的视频，添加和删除自己的评论。
- im推送，可以推送im消息到个人或者房间。

### 安装

已升级到golang1.12，基于gomod管理依赖。

* 安装依赖

```
export GOPROXY=goproxy.io
go mod download
```

* 编译gateway服务

```
cd gateway/cli && go build && cd -
```

* 编译logic服务

```
cd logic/cli && go build && cd -
```

### 架构

* api: api接口
    * 对外暴露的api接口

* scheduler: 定时任务
    * 异步处理评论删除的任务
   
* streamserver: 流媒体服务
    * 负责视频流的主要逻辑

* gateway: 长连接网关
    * 海量长连接按BUCKET打散, 减小推送遍历的锁粒度
    * 按广播/房间粒度的消息前置合并, 减少编码CPU损耗, 减少系统网络调用, 巨幅提升吞吐

* logic: 逻辑服务器
    * 本身无状态, 负责将推送消息分发到所有gateway节点
    * 对调用方暴露HTTP/1接口, 方便业务对接
    * 采用HTTP/2长连接RPC向gateway集群分发消息

### 潜在问题

* 推送主要瓶颈是gateway层而不是内部通讯, 所以gateway和logic之间仍旧采用了小包通讯(对网卡有PPS压力), 同时logic为业务提供了批量推送接口来缓解特殊需求.
