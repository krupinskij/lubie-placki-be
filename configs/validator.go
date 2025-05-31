package configs

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Result struct {
	Key     string
	Message string
}

type validator func(name string, value any, arg string) (Result, bool)

var required validator = func(name string, value any, arg string) (Result, bool) {
	if value == "" {
		return Result{Key: "required", Message: fmt.Sprintf("Field \"%v\" is required", name)}, false
	}

	return Result{}, true
}

var max validator = func(name string, value any, arg string) (Result, bool) {
	num, ok := value.(int)
	if !ok {
		return Result{Key: "max", Message: fmt.Sprintf("Field \"%v\" is not int", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "max", Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if num > max {
		return Result{Key: "max", Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too large (max %v)", name, value, max)}, false
	}

	return Result{}, true
}

var min validator = func(name string, value any, arg string) (Result, bool) {
	num, ok := value.(int)
	if !ok {
		return Result{Key: "min", Message: fmt.Sprintf("Field \"%v\" is not int", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "min", Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if num < min {
		return Result{Key: "min", Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too small (min %v)", name, value, min)}, false
	}

	return Result{}, true
}

var maxStringLength validator = func(name string, value any, arg string) (Result, bool) {
	str, ok := value.(string)
	if !ok {
		return Result{Key: "maxStringLength", Message: fmt.Sprintf("Field \"%v\" is not string", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "maxStringLength", Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if len(str) > max {
		return Result{Key: "maxStringLength", Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too long (max %v)", name, value, max)}, false
	}

	return Result{}, true
}

var minStringLength validator = func(name string, value any, arg string) (Result, bool) {
	str, ok := value.(string)
	if !ok {
		return Result{Key: "minStringLength", Message: fmt.Sprintf("Field \"%v\" is not string", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "minStringLength", Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if len(str) < min {
		return Result{Key: "minStringLength", Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too short (min %v)", name, value, min)}, false
	}

	return Result{}, true
}

var maxArrayLength validator = func(name string, value any, arg string) (Result, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Result{Key: "maxArrayLength", Message: fmt.Sprintf("Field \"%v\" is not an array", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "maxArrayLength", Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if reflect.ValueOf(value).Len() > max {
		return Result{Key: "maxArrayLength", Message: fmt.Sprintf("Field \"%v\" is too long (max %v)", name, max)}, false
	}

	return Result{}, true
}

var minArrayLength validator = func(name string, value any, arg string) (Result, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Result{Key: "minArrayLength", Message: fmt.Sprintf("Field \"%v\" is not an array", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Result{Key: "minArrayLength", Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if reflect.ValueOf(value).Len() < min {
		return Result{Key: "minArrayLength", Message: fmt.Sprintf("Field \"%v\" is too short (min %v)", name, min)}, false
	}

	return Result{}, true
}

var validatorMap = map[string]validator{
	"required":        required,
	"max":             max,
	"min":             min,
	"maxStringLength": maxStringLength,
	"minStringLength": minStringLength,
	"maxArrayLength":  maxArrayLength,
	"minArrayLength":  minArrayLength,
}

func Validate(u any, optional_name ...string) (Result, bool) {
	name := ""
	if len(optional_name) > 0 {
		name = fmt.Sprintf("%v.", optional_name[0])
	}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := range v.NumField() {
		field := t.Field(i)
		value := v.Field(i)
		iface := value.Interface()

		tag := field.Tag.Get("validate")

		deep := false

		if tag != "" {
			keys := strings.SplitSeq(tag, ",")

			for key := range keys {
				re := regexp.MustCompile(`(\w*):?(\d*)?`)
				seq := re.FindAllStringSubmatch(key, -1)[0]

				if seq[1] == "deep" {
					deep = true
					continue
				}

				fun, ok := validatorMap[seq[1]]
				if !ok {
					return Result{Message: fmt.Sprintf("Key %v not exist", seq[1])}, false
				}

				if message, ok := fun(fmt.Sprintf("%v%v", name, field.Name), iface, seq[2]); !ok {
					return message, false
				}

			}
		}

		var kind = reflect.TypeOf(iface).Kind()
		if deep && (kind == reflect.Array || kind == reflect.Slice) {
			for iter, item := range value.Seq2() {
				if result, ok := Validate(item.Interface(), fmt.Sprintf("%v%v[%v]", name, field.Name, iter)); !ok {
					return result, false
				}
			}
		}

		if deep && kind == reflect.Struct {
			if result, ok := Validate(iface, fmt.Sprintf("%v%v", name, field.Name)); !ok {
				return result, false
			}
		}
	}

	return Result{}, true
}
