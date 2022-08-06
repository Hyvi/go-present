记录 protobuf 的学习内容

## Note
`cd` the `protobuf` dir, then run the command

```bash
 protoc --go_out=out/ --go_opt=paths=source_relative addressbook.proto
 ```

 generate the go (example.proto) file in the dir `out`

在 nvim 中执行单元测试(:copen 打开 quickfix 窗口）
```
:AsyncRun go test ./protobuf  -run TestWrite -v
```

## Encoding
https://developers.google.com/protocol-buffers/docs/encoding


### Base 128 Varints

Varints are a method of serializing integers using one or more bytes. Smaller numbers take a smaller number of bytes. 用于序列化整数

Each byte in a varint, except the last byte, has the most significant bit (MSB) set – this indicates that there are further bytes to come.  除了最后一个字节之外每个字节都会有 MSB 位（还有字节）

The lower 7 bits of each byte are used to store the two's complement representation of the number in groups of 7 bits, least significant group first.

详解 varint 编码原理：https://segmentfault.com/a/1190000020500985

### Message Structure
A protocol buffer message is a series of key-value pairs.

The binary version of a message just uses the `field's number` as the key – the name and declared type for each field can only be determined on the decoding end by referencing the message type's definition (i.e. the .proto file). 在二进制存储里，field number 将作为 key。而 name 以及类型用在解码端




### More Value Type


#### Signed Integers
If you use int32 or int64 as the type for a negative number, the resulting varint is always ten bytes long – it is, effectively, treated like a very large unsigned integer. 占满了字节：

If you use one of the signed types, the resulting varint uses ZigZag encoding, which is much more efficient. 所以使用 zigzag 编码

#### Enbedded Messages
Embedded messages are treated in exactly the same way as strings (wired type = 2) 

#### Packed Repeated Fields
These function like repeated fields, but are encoded differently. A packed repeated field containing zero elements does not appear in the encoded message


## TODO

[Buf Tour](https://docs.buf.build/tour/introduction)

## 参考
1. Google: Go Generated Code

2. Google: Protocol Buffer Basics: Go

3. [A google protobuf decoder by posting data on web page based on Golang](https://github.com/superryanguo/postDataWebGo)

4. [Uber Buf](https://github.com/bufbuild/buf) Defining APIs using an IDL provides a number of benefits over simply exposing JSON/REST services, and today, Protobuf is the most stable, widely-adopted IDL in the industry.

