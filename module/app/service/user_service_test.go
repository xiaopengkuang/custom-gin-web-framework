package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	userservice := &UserService{}

	us := reflect.ValueOf(userservice)
	usType := reflect.TypeOf(userservice)
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf("1"))
	fmt.Println(us.MethodByName("Name2").Call(args))
	mm := usType.NumMethod()
	for i := 0; i < mm; i++ {
		meth := usType.Method(i)
		fmt.Println(meth.Name)

		if meth.Name == "Name2" {

			metttt := meth.Func
			fmt.Println(metttt.IsValid())
			fmt.Println(metttt.IsNil())
			if metttt.IsValid() || !metttt.IsNil() {
				fmt.Println("nooo")
			}

			fmt.Println(metttt.Kind())

			fmt.Println(metttt.Call(args))
		}
	}

	fmt.Println(us)

}
