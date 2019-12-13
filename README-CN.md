# Gaping ---  一个golang实现的测试远程主机端口稳定性的工具，灵感来自paping

[![Build Status](https://travis-ci.org/tianyazc/gaping.svg?branch=master)](https://github.com/tianyazc/gaping)
[![](https://badge.juejin.im/entry/5c96245ce51d450a0e0a5c08/likes.svg?style=flat)](https://juejin.im/post/5c9623246fb9a070d14c33fc)

[English](https://github.com/tianyazc/gaping)

![]( http://img.chulinx.xyz/termtosvg_wvabfoi8.svg)


## 安装 ↯↯↯
### 编译安装
  - Golang 1.11 or Golang 1.12
```bash
git clone https://github.com/tianyazc/gaping.git && cd gaping && export GO111MODULE=on go build
```

### 二进制文件下载
Download form [**Release**](https://github.com/tianyazc/gaping/releases/tag/0.1)

## 使用 ↯↯↯

```bash
Gaping  v0.02- Copyright (c) 2019 Mike chulinx
Example: pping -h 127.0.0.1,localhost -p 80 -type tcp
  -V	Show Version and exit
  -c int
    	Tcp test conuts (default 100000000)
  -h string
    	Dest host ipaddress
  -help
    	Show gaping help
  -p int
    	Dest host port
  -types string
    	Set network protocol (default "tcp")
  -v	Show Version and exit
```

## 举个例子 ↯↯↯

- 一直ping

```bash
$ ./gaping_darwin_amd64 -h 115.239.210.27 -p 80
2019-03-20 17:25:00 Connected to 115.239.210.27:80: time=10.505ms protocol=TCP port=80
2019-03-20 17:25:01 Connected to 115.239.210.27:80: time=11.18ms protocol=TCP port=80
2019-03-20 17:25:02 Connected to 115.239.210.27:80: time=10.814ms protocol=TCP port=80
2019-03-20 17:25:03 Connected to 115.239.210.27:80: time=12.213ms protocol=TCP port=80
2019-03-20 17:25:04 Connected to 115.239.210.27:80: time=11.144ms protocol=TCP port=80
2019-03-20 17:25:05 Connected to 115.239.210.27:80: time=11.545ms protocol=TCP port=80
2019-03-20 17:25:06 Connected to 115.239.210.27:80: time=10.984ms protocol=TCP port=80
2019-03-20 17:25:07 Connected to 115.239.210.27:80: time=10.63ms protocol=TCP port=80
2019-03-20 17:25:08 Connected to 115.239.210.27:80: time=10.768ms protocol=TCP port=80
2019-03-20 17:25:09 Connected to 115.239.210.27:80: time=11.549ms protocol=TCP port=80
2019-03-20 17:25:10 Connected to 115.239.210.27:80: time=12.01ms protocol=TCP port=80
2019-03-20 17:25:11 Connected to 115.239.210.27:80: time=11.086ms protocol=TCP port=80
```
- ping 一次
```bash
$ ./gaping_darwin_amd64 -h 115.239.210.27 -p 80 -c 1
2019-03-20 17:27:16 Connected to 115.239.210.27:80: time=13.396ms protocol=TCP port=80
```
- ping 多台主机一个端口
```bash
$ ./gaping -h 127.0.0.1,localhost -p 1087 -types udp
2019-12-13 16:02:44 127.0.0.1 Connected to  127.0.0.1:1087: time=251µs protocol=udp port=1087
2019-12-13 16:02:45 localhost Connected to  localhost:1087: time=6.779ms protocol=udp port=1087
2019-12-13 16:02:46 127.0.0.1 Connected to  127.0.0.1:1087: time=283µs protocol=udp port=1087
2019-12-13 16:02:47 localhost Connected to  localhost:1087: time=1.343ms protocol=udp port=1087
```
