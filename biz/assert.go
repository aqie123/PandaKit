package biz

import (
	"fmt"
	"github.com/XM-GO/PandaKit/utils"
	"reflect"
)

func ErrIsNil(err error, msg string, params ...any) {
	if err != nil {
		if err.Error() == "record not found" {
			return
		}
		panic(any(NewBizErr(fmt.Sprintf(msg, params...))))
	}
}

func ErrIsNil2(err error) {
	if err != nil {
		if err.Error() == "record not found" {
			return
		}
		panic(any(NewBizErr(fmt.Sprintf(err.Error()))))
	}
}

func ErrIsNil3(msg string, params ...any) {
	panic(any(NewBizErr(fmt.Sprintf(msg, params...))))
}

func ErrIsNil4(code int16, msg string, params ...any) {
	panic(any(NewBizErrCode(code, fmt.Sprintf(msg, params...))))
}

func ErrIsNilAppendErr(err error, msg string) {
	if err != nil {
		panic(any(NewBizErr(fmt.Sprintf(msg, err.Error()))))
	}
}

func IsNil(err error) {
	switch t := err.(type) {
	case *BizError:
		panic(any(t))
	case error:
		panic(any(NewBizErr(fmt.Sprintf("非业务异常: %s", err.Error()))))
	}
}

func IsTrue(exp bool, msg string, params ...any) {
	if !exp {
		panic(any(NewBizErr(fmt.Sprintf(msg, params...))))
	}
}

func IsTrueBy(exp bool, err BizError) {
	if !exp {
		panic(any(err))
	}
}

func NotEmpty(str string, msg string, params ...any) {
	if str == "" {
		panic(any(NewBizErr(fmt.Sprintf(msg, params...))))
	}
}

func NotNil(data any, msg string) {
	if reflect.ValueOf(data).IsNil() {
		panic(any(NewBizErr(msg)))
	}
}

func NotBlank(data any, msg string) {
	if utils.IsBlank(reflect.ValueOf(data)) {
		panic(any(NewBizErr(msg)))
	}
}

func IsEquals(data any, data1 any, msg string) {
	if data != data1 {
		panic(any(NewBizErr(msg)))
	}
}

func Nil(data any, msg string) {
	if !reflect.ValueOf(data).IsNil() {
		panic(any(NewBizErr(msg)))
	}
}
