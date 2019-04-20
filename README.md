golang-enum-example
===================

* Stringer 面倒。。。
* Thrift を使うと良い感じに生成してくれる

How to run
=====
```
% make clean test
```

Description
=====

```go:enum.go
enum Gender {
    Unknown = 0,
    Male = 1,
    Female = 2,
}
```


Precondition
=====
thrift バイナリーがインストールされているようにしてください。brew とか。

```
% thrift --version
Thrift version 0.12.0
```
