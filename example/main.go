package main

import (
	"fmt"
	oracle "github.com/linxlib/gorm_oracle"
	"gorm.io/gorm"
)

type AS struct {
	A int
}

func main() {
	//gorm.Open(oracle.Open("oracle://user:password@ip:port/service"))

	db, err := gorm.Open(oracle.Open("oracle://user:password@ip:port/service"))
	if err != nil {
		fmt.Println(err)
		return
	}

	result := new(int64)
	s := db.Debug().Table(`"USERNAME"."TABLENAME"`).
		Where("A=?", 1001).
		Count(result)
	if s.Error != nil {
		fmt.Println(s.Error)
		return
	}
	fmt.Println(*result)

}
