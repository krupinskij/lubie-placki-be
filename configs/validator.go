package configs

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Message struct {
	Message string
}

type validator func(name string, value any, arg string) (Message, bool)

var required validator = func(name string, value any, arg string) (Message, bool) {
	if value == nil {
		return Message{Message: fmt.Sprintf("Field \"%v\" is required", name)}, false
	}

	return Message{}, true
}

var max validator = func(name string, value any, arg string) (Message, bool) {
	num, ok := value.(int)
	if !ok {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not int", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if num > max {
		return Message{Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too large (max %v)", name, value, max)}, false
	}

	return Message{}, true
}

var min validator = func(name string, value any, arg string) (Message, bool) {
	num, ok := value.(int)
	if !ok {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not int", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if num < min {
		return Message{Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too small (min %v)", name, value, min)}, false
	}

	return Message{}, true
}

var maxStringLength validator = func(name string, value any, arg string) (Message, bool) {
	str, ok := value.(string)
	if !ok {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not string", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if len(str) > max {
		return Message{Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too long (max %v)", name, value, max)}, false
	}

	return Message{}, true
}

var minStringLength validator = func(name string, value any, arg string) (Message, bool) {
	str, ok := value.(string)
	if !ok {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not string", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if len(str) < min {
		return Message{Message: fmt.Sprintf("Field \"%v\" of value \"%v\" is too short (min %v)", name, value, min)}, false
	}

	return Message{}, true
}

var maxArrayLength validator = func(name string, value any, arg string) (Message, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not an array", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg max for field \"%v\" is not of type int", name)}, false
	}

	if reflect.ValueOf(value).Len() > max {
		return Message{Message: fmt.Sprintf("Field \"%v\" is too long (max %v)", name, max)}, false
	}

	return Message{}, true
}

var minArrayLength validator = func(name string, value any, arg string) (Message, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Message{Message: fmt.Sprintf("Field \"%v\" is not an array", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: fmt.Sprintf("Arg min for field \"%v\" is not of type int", name)}, false
	}

	if reflect.ValueOf(value).Len() < min {
		return Message{Message: fmt.Sprintf("Field \"%v\" is too short (min %v)", name, min)}, false
	}

	return Message{}, true
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

func Validate(u any, optional_name ...string) (Message, bool) {
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

		if tag != "" {
			keys := strings.SplitSeq(tag, ",")

			for key := range keys {
				re := regexp.MustCompile(`(\w*):?(\d*)?`)
				seq := re.FindAllStringSubmatch(key, -1)[0]

				fun, ok := validatorMap[seq[1]]
				if !ok {
					return Message{Message: fmt.Sprintf("Key %v not exist", seq[1])}, false
				}

				if message, ok := fun(fmt.Sprintf("%v%v", name, field.Name), iface, seq[2]); !ok {
					return message, false
				}

			}
		}

		var kind = reflect.TypeOf(iface).Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			for iter, item := range value.Seq2() {
				if message, ok := Validate(item.Interface(), fmt.Sprintf("%v%v[%v]", name, field.Name, iter)); !ok {
					return message, false
				}
			}
		}

		if kind == reflect.Struct {
			if message, ok := Validate(iface, fmt.Sprintf("%v%v", name, field.Name)); !ok {
				return message, false
			}
		}
	}

	return Message{}, true
}
