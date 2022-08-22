# GORM Oracle driver

## Description
GORM-ORACLE is a GORM driver that does not depend on the Oracle client, Based on [github.com/wdrabbit/gorm-oracle](https://github.com/wdrabbit/gorm-oracle), upgrade https://github.com/sijms/go-ora to v2 version (which supports more recently version of oracle)
, not thoroughly tested and not recommended for production use.

升级了依赖版本到v2，用于支持更高版本的oracle

项目集成go-ora驱动,无需安装oracle客户端。

## DB Driver
[go-ora v2](https://github.com/sijms/go-ora/v2)
A pure golang development of Oracle driver, do not need to install Oracle client.

## Quick Start
### how to install 
```bash
go get github.com/wdrabbit/gorm-oracle
```
### usage
```go
import (
    "gorm.io/gorm"
    "oracle "github.com/linxlib/gorm_oracle"
    "log"
)

//ORM bean
type TestBean struct {
	Id string `gorm:"column:ID;not null;primaryKey;size:36"`
	Field1 string `gorm:"column:FIELD1;size:255"`
	Field2 string `gorm:"column:FIELD2;size:255"`
	State string `gorm:"column:STATE;size:2"`
	CreateAt time.Time `gorm:"column:CREATE_AT"`
}

func main(){

    databaseURL := "oracle://username:password@host:port/db"
    db, err := gorm.Open(oracle.Open(databaseURL),&gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    
    //Insert
    datas := []bean.TestBean{
        {"a","a","a","01",time.Now()},
        {"b","a","a","01",time.Now()},
    }
    db = db.Debug().Create(&datas)

    //Update
    db = db.Debug().Where("id=?","a").Model(&bean.TestBean{}).Update("state","02")

    //Delete
    db := db.Where("id = ?","a").Delete(&bean.TestBean{})

    //Select
    var rows []bean.TestBean
    db = db.Debug().Find(&rows)
    
    //do somethings

}
```
