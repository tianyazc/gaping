# Gaping ---  A remote host port stability test tool inspired by paping


## Install ↯↯↯
### Compile and install
  - Golang 1.11 or Golang 1.12
```
git clone https://github.com/tianyazc/gaping.git && cd gaping && export GO111MODULE=on go build
```

### Release Download
Download form [**Release**](https://github.com/tianyazc/gaping/releases/tag/0.1)

## Usage ↯↯↯

```
$ ./gaping_darwin_amd64 -h
flag needs an argument: -h
Usage of ./gaping_darwin_amd64:
  -V	Show Version and exit
  -c int
    	Tcp test conuts (default 100000000)
  -h string
    	Dest host ipaddress
  -help
    	Show gaping help
  -p string
    	Dest host port
  -v	Show Version and exit
```

## Example ↯↯↯

- always ping

```
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
- ping once
```
$ ./gaping_darwin_amd64 -h 115.239.210.27 -p 80 -c 1
2019-03-20 17:27:16 Connected to 115.239.210.27:80: time=13.396ms protocol=TCP port=80
```
