# Pi-db

  本文使用go语言的server框架[gin](https://github.com/gin-gonic/gin) 实现在树莓派上提供远程访问数据库(sqlite)服务.

+ ***项目结构***

```
.
├── README.md
├── apis
│   └── cpuInfo.go
├── database
│   └── cpu.db
├── main.go
├── models
│   ├── cpu.go
│   └── sqlite.go
└── routers
    └── router.go
```

 在routers中设置好相关的路由(目前只有GET)，由apis中来实现，它经由models来实现。

+ ***编译与安装***

1. 在电脑平台上先安装 gin:

```
$ go get -u github.com/gin-gonic/gin
```

2. ***编译***

```
$ go build
```

3. ***交叉编译树莓派版本***

​      由于用sqlite 数据库，需要用cgo来编译sqlite对应的版本，需要安装gcc的交叉编译器[arm-linux-gcc](https://github.com/downloads/UnhandledException/ARMx/ARMx-2009q3-67.tar.bz2)，

​        把ARMx-2009q3-67.tar.bz2下载后解压放在/usr/local/arm-gcc下：

```
/usr/local/arm-gcc
├── arm-none-linux-gnueabi
├── bin
├── lib
├── libexec
└── share
```

b. 再声明环境变量：

```
$ export PATH="/usr/local/arm-gcc/bin:$PATH"
```

c. 检验一下是否安装成功：

```
$ which arm-none-linux-gnueabi-gcc
/usr/local/arm-gcc/bin/arm-none-linux-gnueabi-gcc
```

d. 编译树莓版本：

```
$ CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-none-linux-gnueabi-gcc go build
```

+ ***运行***

   把生成的目标文件pi-db和文件夹database一起传到树莓派上，启动程序：

```
$ ./pi-db
```



+ ***访问方法***

```
curl -i http://localhost:8080/api/cpu/T
```

