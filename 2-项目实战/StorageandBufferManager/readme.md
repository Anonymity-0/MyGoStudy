# Storage and Buffer Manager


## 项目描述

这是一个使用 Go 语言实现的缓冲区管理器。它使用 LRU 算法来管理缓冲区中的页面，并提供了读取和写入页面的方法。此外，它还提供了统计信息，如缓冲区的命中率和输入/输出操作的数量。

## 安装

首先，你需要安装 Go。你可以从 [Go 的官方网站](https://golang.org/) 下载并安装它。


## 使用

你可以使用 `go run` 命令来运行这个项目：

```bash
go run *.go
```

或

```bash
go run main.go BCB.go BFrame.go BMgr.go DSMgr.go LRU.go Trace.go
```

项目会创建一个新的文件 `data.dbf`，并执行测试文件中的操作。在所有操作执行完成后，它会打印出缓冲区的命中率和输入/输出操作的数量。

