## 证书生成

程序使用双向验证，所以需要 `CA`, `Server`, `Client` 这三种证书，`Server` 和 `Client` 的证书使用 `CA` 私钥签名生成。
`CA` 即相当与根证书，为 `Server` 和 `Client` 做信任担保。


### 生成 `CA`

1. 私钥

`openssl genrsa -out CA.key 2048`

2. 数字证书

`openssl req -x509 -new -nodes -key CA.key -subj "/CN=localhost" -days 5000 -out CA.crt`


### 生成 `Server`

1. 私钥

`openssl genrsa -out Server.key 2048`

2. 证书签名请求(Certificate Sign Request)

`openssl req -new -key Server.key -subj "/CN=localhost" -days 5000 -out Server.csr`

3. 证书

使用 `CA` 进行签名:

`openssl x509 -req -in Server.csr -CA CA.crt -CAkey CA.key -CAcreateserial -out Server.crt -days 5000`


### 生成 `Client`

1. 私钥

`openssl genrsa -out Client.key 2048`

2. 证书签名请求(Certificate Sign Request)

`openssl req -new -key Client.key -subj "/CN=localhost" -days 5000 -out Client.csr`

3. 证书

使用 `CA` 进行签名:

`openssl x509 -req -in Client.csr -CA CA.crt -CAkey CA.key -CAcreateserial -out Client.crt -days 5000`


### 查看证书信息

`openssl x509 -text -in <crt file> -noout`


### ECC 证书

相比上面的 `RSA` 证书，`ECC` 证书体积更小，但兼容性没有 `RSA` 好，不支持部分老旧操作系统、浏览器。

创建命令如下:

``` shell
# secp256r1
openssl ecparam -genkey -name secp256r1 | openssl ec -out CA.key

#secp384r1
openssl ecparam -genkey -name secp384r1 | openssl ec -out CA.key
```

使用交互的方式创建 `CSR`，其中 `Common Name` 必须为你的域名：

`openssl req -new -key CA.key -out CA.csr`
