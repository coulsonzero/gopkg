# gopkg
A module repository for Golang 


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

now

```sh
➡︎  🍭  tree
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

### Usage (v0.4.0)
#### Install module
```go
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
