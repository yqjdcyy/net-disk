# FOR

- 模拟网盘进行本机数据定时备份

# TODO

## Backend

### Server

- ~~上传~~
- ~~日志~~
- ~~优化~~
- ~~接口~~
  - ~~列表~~
  - ~~删除~~

### Client

- ~~配置~~
- ~~任务~~
- 优化
  - 支持多目录
  - 上传异常时，中断删除文件操作

## Shell(需安装 imagemagick)

- ~~目录优化~~
- ~~指定目录的定时优化~~

## Front

- 目录展现
- 列表展现
  - 按文件类型展示
  - 支持大图

## Bug/Optimize

- ~~windows.seperator 2 linux.seperator~~
  - ~~异常目录处理~~
  - ~~mkdirAll permission denied~~
- ~~输出转移到 logger~~
  - ~~保持文件操作、网络请求的独立性，放弃 log 模块的侵入~~
- ~~上传图片无法正常观看~~
- 支持缩略图展现

# Command

## CMD

- `set gopath`
- `set gopath=%gopath%;"..."`

## PowerShell

- `$env:gopath`
- `$env:gopath="..."`

# 参考

## ImageMagick

- [使用 ImageMagick 如何对图片进行全面压缩](https://blog.csdn.net/Shijun_Zhang/article/details/6702752)
- [imagemagick.option - quality](https://www.imagemagick.org/script/command-line-options.php#quality)

## 其它

- [golang 中使用 url encoding 遇到的小坑](http://weakyon.com/2017/05/04/something-of-golang-url-encoding.html)
