# Sample HTTPS Server

一个 `Go` 实现的简单的 `HTTPS` 服务器，使用双向证书验证。


## Prepare

### Dependencies
+ `openssl`

### 证书

参照：[证书生成](./cert.md)


## Compile

``` shell
cd ./src
go build server.go
go build client.go
```

## Usage

+ Launch server

    `./server <CA.cert> <Server.crt> <Server.key> <port>`

+ Launch client

    `./client <CA.cert> <Client.crt> <Client.key>`
