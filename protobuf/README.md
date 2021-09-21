记录 protobuf 的学习内容

## Note
`cd` the `protobuf` dir, then run the command 

```bash 
 protoc --go_out=out/ --go_opt=paths=source_relative addressbook.proto
 ```

 generate the go (example.proto) file in the dir `out`

## TODO 

[Buf Tour](https://docs.buf.build/tour/introduction)

## 参考
1. Google: Go Generated Code

2. Google: Protocol Buffer Basics: Go

3. [A google protobuf decoder by posting data on web page based on Golang](https://github.com/superryanguo/postDataWebGo)

4. [Uber Buf](https://github.com/bufbuild/buf) Defining APIs using an IDL provides a number of benefits over simply exposing JSON/REST services, and today, Protobuf is the most stable, widely-adopted IDL in the industry.

