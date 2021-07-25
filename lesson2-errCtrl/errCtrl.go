package main

import (
	"errCtrl/dao"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	myDb := dao.New()
	user, err := myDb.GetUser("xiaomi")
	if err != nil {
		fmt.Println(errors.Cause(err))
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
