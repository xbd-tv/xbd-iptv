# xbd-iptv
小板凳 优质 IPTV 节目频道资源

### CND
使用 [jsdelivr gh](https://www.jsdelivr.com/github) CDN 加速

*注意使用以下方式刷新 CDN*
```
curl https://purge.jsdelivr.net/{....地址}  // 刷新CDN
```
##### M3U8
|  名称   | 地址  |
|  ----  | ----  |
|cn|https://cdn.jsdelivr.net/gh/xbd-tv/xbd-iptv@main/channels/cn.m3u8|
##### JSON 频道信息
|  名称   | 地址  |
|  ----  | ----  |
|cn|https://cdn.jsdelivr.net/gh/xbd-tv/xbd-iptv@main/channels/cn.json|

### 引用资源
[IPTV-ORG](https://github.com/iptv-org/iptv)


## 开发调试
1. 测试 m3u 解析
```bash
go run main.go m3u --in="C:\Users\feng\Downloads\cn.m3u"
```
