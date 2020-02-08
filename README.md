# 文件服务器

在当前目录启动文件服务器。

## 帮助

```
$ fileserver -h
```

```
文件服务器

Usage:
  fileserver [flags]

Flags:
  -d, --dir string    指定文件服务器的根目录
  -h, --help          help for fileserver
      --host string   指定服务器的IP (default "0.0.0.0")
  -p, --port string   指定服务器的端口 (default "18020")
```

### 使用

```
$ ./fileserver      // 在当前目录启动一个文件服务器
```

```bash
文件服务器已启动:
  监听IP: 0.0.0.0
  本机IP: 192.168.0.6
  端口: 18020
  目录: /Users/fenggese/Workspace/golang/mod/fileserver

使用：
  查看文件列表:
  curl "http://192.168.0.6:18020/"

  下载文件: curl -o 文件名 主机:端口/文件名
  curl -o text.txt 192.168.0.6:18020/text.txt
```

#### 应用场景

在docker容器中下载自己电脑中的文件.

docker容器中只需要有curl命令即可。

<img src="https://github.com/FengGeSe/fileserver/blob/master/static/fileserver.jpg">




ps: 好用记得点star ^.^~

