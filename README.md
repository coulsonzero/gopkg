# gopkg
A module repository for Golang

[![Go gopkg/pro](https://pkg.go.dev/badge/github.com/coulsonzero/gopkg/pro.svg)](https://pkg.go.dev/github.com/coulsonzero/gopkg/pro)

[//]: # (<a href="https://pkg.go.dev/github.com/coulsonzero/gopkg/pro"><img src="https://pkg.go.dev/badge/github.com/coulsonzero/gopkg/pro.svg" alt="Go Reference"></a>)


### What I've done ?  (v0.4.0)
remove all subdirectories to update the import package

before
```sh
$ tree
.
├── encrypt
├   ├── md5.go          # package gopkg
├   └── bcrypt.go       # package gopkg
└── fileconfig
    ├── env.go          # package gopkg
    ├── ini.go          # package gopkg
    └── yml.go          # package gopkg
```
before usage (v0.2.0)
```go
import (
    gopkg1 "github.com/coulsonzero/gopkg/encrypt"
    gopkg2 "github.com/coulsonzero/gopkg/fileconfig"
)
```
```go
gopkg1.HashPassword("admin123")
gopkg2.ConfigEnv()
```

v2 (v0.4.0)

```sh
$ tree
.
├── bcrypt.go
├── md5.go
├── env.go
├── ini.go
├── yml.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

v3 (v0.8.1-beta)

> most different: import `unsafe`: `//go:linkname readSql .../pro.ReadSql`

```sh
$ tree
.
├── LICENSE
├── README.md
├── arrays
│   ├── set.go
│   └── slice.go
├── color
│   └── color.go
├── conf
│   ├── config.ini
│   ├── config.yml
│   └── user.sql
├── demo
│   ├── file_demo.go
│   ├── jsondemo.go
│   ├── reflect_demo.go
│   ├── slice_demo.go
│   ├── sql_demo.go
│   └── string_demo.go
├── files
│   └── file.go
├── go.mod
├── go.sum
├── pro
│   ├── arm.s
│   ├── cmd.go
│   ├── dsn.go
│   ├── json.go
│   ├── password.go
│   ├── reflect.go
│   ├── sort.go
│   ├── sql.go
│   ├── string.go
│   ├── time.go
│   └── zero-arm.go
├── push.sh
├── runtime
│   └── unsafe
│       └── pro
│           ├── cmd.go
│           ├── dsn.go
│           ├── json.go
│           ├── password.go
│           ├── reflect.go
│           ├── sort.go
│           ├── sql.go
│           ├── string.go
│           ├── time.go
│           ├── zero-dsn.go
│           └── zero-log.go
├── shell
│   ├── dev-push.sh
│   ├── master-push.sh
│   ├── tag-delete.sh
│   └── tag-release.sh
├── test.sh
└── testing
├── cmd_test.go
├── json_test.go
├── pro_example.go
├── slice_test.go
└── string_test.go
```

### Usage (v0.4.0)
#### Install module
```sh
$ go get github.com/coulsonzero/gopkg
```

#### How to import it ?

```go
import "github.com/coulsonzero/gopkg"
```

#### How to use it ?
```go
gopkg.HashPassword("admin123")
gopkg.ConfigEnv(testEnvArr)
```


### Usage(v0.8.1-beta)


#### Install module
```sh
$ go get github.com/coulsonzero/gopkg/pro
```

#### How to import it ?

```go
import "github.com/coulsonzero/gopkg/pro"
```

#### How to use it ?
```go
pro.HashPassword("admin123")
pro.ConfigEnv(testEnvArr)
```


