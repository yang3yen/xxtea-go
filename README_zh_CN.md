# XXTEA Golang

## 简介

[`xxtea`](https://en.wikipedia.org/wiki/XXTEA) 是一种分组密码。本项目是`xxtea`加密算法的Golang实现，参考Python的[xxtea](https://github.com/ifduyue/xxtea) 项目实现，加密解密结果同Python的[xxtea](https://github.com/ifduyue/xxtea) 项目相同。

版本对应：

| xxtea-go |      xxtea-python      |
|:--------:| :--------------------: |
|  v1.0.x  | v1.2.0, v1.3.0, v2.0.0 |

## 安装

```shell
go get github.com/yang3yen/xxtea-go
```

## 使用

本模块提供7个函数：`URandom`，`Encrypt`，`Decrypt`，`EncryptBase64`，`DecryptBase64`，`EncryptHex`，`DecryptHex`。

```go
package main

import (
	"bytes"
	"fmt"
	"time"
	"github.com/yang3yen/xxtea-go/xxtea"
)

func main() {
	data := []byte("xxtea-test-case")
	key, _ := xxtea.URandom(16, time.Now().UnixNano())

	enc, _ := xxtea.Encrypt(data, key, true, 0)
	dec, _ := xxtea.Decrypt(enc, key, true, 0)
	if bytes.Equal(data, dec) {
		fmt.Println("encrypt success!")
	} else {
		fmt.Println("encrypt fail!")
	}

	encHex, _ := xxtea.EncryptHex(data, key, true, 0)
	decHex, _ := xxtea.DecryptHex(encHex, key, true, 0)
	if bytes.Equal(data, decHex) {
		fmt.Println("encrypt hex success!")
	} else {
		fmt.Println("encrypt hex fail!")
	}

	encB64, _ := xxtea.EncryptBase64(data, key, true, 0)
	decB64, _ := xxtea.DecryptBase64(encB64, key, true, 0)
	if bytes.Equal(data, decB64) {
		fmt.Println("encrypt base64 success!")
	} else {
		fmt.Println("encrypt base64 fail!")
	}
}
```