# XXTEA Golang

## Introduction

[`xxtea`](https://en.wikipedia.org/wiki/XXTEA) is a block cipher. This project is the Golang implementation of the `xxtea` encryption algorithm. Refer to the implementation of the [xxtea](https://github.com/ifduyue/xxtea) project of Python. The encryption and decryption results are the same as the [xxtea](https://github.com/ifduyue/xxtea) project of Python.

version:

| xxtea-go |      xxtea-python      |
|:--------:| :--------------------: |
|  v1.0.x  | v1.2.0, v1.3.0, v2.0.0 |

## Installation

```shell
go get github.com/yang3yen/xxtea-go
```

## Usage

This module provides seven functions: `URandom`,`Encrypt`,`Decrypt`,`EncryptBase64`,`DecryptBase64`,`EncryptHex`,`DecryptHex`.

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