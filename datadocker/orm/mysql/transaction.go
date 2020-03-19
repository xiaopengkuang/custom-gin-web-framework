package mysql

import (
	"fmt"
)

// 事务类
type Transaction struct {
	FunctionList []Function
}

type Function struct {
	Func       interface{}
	Args       []interface{}
	IsFunction bool
}

func (f *Function) executeFunc() error {
	switch f.Func.(type) {
	case func(args ...interface{}) error:
		return f.Func.(func(args ...interface{}) error)(f.Args)
	case func() error:
		return f.Func.(func() error)()
	default:
		return fmt.Errorf("excute not support")

	}

	return f.Func.(func(args ...interface{}) error)(f.Args)
}

func (f *Function) IsValidFunction() bool {
	switch f.Func.(type) {
	case func(args ...interface{}) error:
		return true
	case func() error:
		return f.Args == nil
	default:
		return false
	}
}

// 添加function
func (t *Transaction) AddTransactionFunc(function interface{}, args ...interface{}) error {
	t.checkFunctionList()

	functionItem := Function{Func: function, Args: args}
	functionItem.IsFunction = functionItem.IsValidFunction()
	if !functionItem.IsFunction {
		return fmt.Errorf("invalid function %+v", function)
	}

	t.FunctionList = append(t.FunctionList, functionItem)

	return nil
}

// 检查functionList
func (t *Transaction) checkFunctionList() {
	if t.FunctionList == nil {
		t.FunctionList = make([]Function, 0)
	}
}

func (t *Transaction) Process() error {
	if t.FunctionList == nil || len(t.FunctionList) == 0 {
		return nil
	}

	DB.Begin()

	for i := 0; i < len(t.FunctionList); i++ {
		err := t.FunctionList[i].executeFunc()
		if err != nil {
			DB.Rollback()
			return err
		}
	}

	DB.Commit()

	return nil
}

//
//type FF func(args ...interface{}) error
//
//func DoTransaction(argsList [][]interface{}, funcList ...func(args ...interface{}) error) error {
//	if funcList == nil || len(funcList) == 0 || argsList == nil || len(argsList) != len(funcList) {
//		return fmt.Errorf("Error: transaction error %+v\n", argsList)
//	}
//
//	DB.Begin()
//
//	for i := 0; i < len(funcList); i++ {
//		err := funcList[i](argsList[i])
//		if err != nil {
//			// 回滚事务
//			DB.Rollback()
//			return fmt.Errorf("err at %d : %s", i, err.Error())
//		}
//	}
//
//	// 确认事务
//	DB.Commit()
//	return nil
//}
