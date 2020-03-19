package mysql

import (
	"fmt"
	"testing"
)

func T1() error {
	fmt.Println("Do T1 function")
	return nil
}

func T2(args ...interface{}) error {
	fmt.Println("Do T2 fucntion: ", args)
	return nil
}

func TestDoTransaction(t *testing.T) {
	OpenDB()

	trans := Transaction{}
	err1 := trans.AddTransactionFunc(T1)
	err2 := trans.AddTransactionFunc(T2, "a", "b")
	err := trans.Process()
	fmt.Println(err)

	fmt.Println(err1)
	fmt.Println(err2)
}
