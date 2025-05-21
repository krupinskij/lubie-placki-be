package validators

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Message struct {
	Message string
}

type Validator func(name string, value any, arg string) (Message, bool)

var required Validator = func(name string, value any, arg string) (Message, bool) {
	if value == nil {
		return Message{Message: fmt.Sprintf("Field %v is required", name)}, false
	}

	return Message{}, true
}

var maxStringLength Validator = func(name string, value any, arg string) (Message, bool) {
	str, ok := value.(string)
	if !ok {
		return Message{Message: fmt.Sprintf("Field %v is not string", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: "Arg max is not int"}, false
	}

	if len(str) > max {
		return Message{Message: fmt.Sprintf("Field %v is too long (max %v)", name, max)}, false
	}

	return Message{}, true
}

var minStringLength Validator = func(name string, value any, arg string) (Message, bool) {
	str, ok := value.(string)
	if !ok {
		return Message{Message: fmt.Sprintf("Field %v is not string", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: "Arg min is not int"}, false
	}

	if len(str) < min {
		return Message{Message: fmt.Sprintf("Field %v is too short (min %v)", name, min)}, false
	}

	return Message{}, true
}

var maxArrayLength Validator = func(name string, value any, arg string) (Message, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Message{Message: fmt.Sprintf("Field %v is not array", name)}, false
	}

	max, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: "Arg max is not int"}, false
	}

	if reflect.ValueOf(value).Len() > max {
		return Message{Message: fmt.Sprintf("Field %v is too long (max %v)", name, max)}, false
	}

	return Message{}, true
}

var minArrayLength Validator = func(name string, value any, arg string) (Message, bool) {
	var kind = reflect.TypeOf(value).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Message{Message: fmt.Sprintf("Field %v is not array", name)}, false
	}

	min, err := strconv.Atoi(arg)
	if err != nil {
		return Message{Message: "Arg min is not int"}, false
	}

	if reflect.ValueOf(value).Len() < min {
		return Message{Message: fmt.Sprintf("Field %v is too short (min %v)", name, min)}, false
	}

	return Message{}, true
}

var ValidatorMap = map[string]Validator{
	"required":        required,
	"maxStringLength": maxStringLength,
	"minStringLength": minStringLength,
	"maxArrayLength":  maxArrayLength,
	"minArrayLength":  minArrayLength,
}

type Animal struct {
	Breed string `validate:"required"`
	Name  string `validate:"maxStringLength:10"`
}

type Address struct {
	Street string `validate:"maxStringLength:5"`
}

type User struct {
	Name    string   `validate:"required,maxStringLength:10"`
	Email   string   `validate:"required"`
	Animals []Animal `validate:"required,minArrayLength:3"`
	Address Address
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
				seq := strings.Split(key, ":")
				fun, ok := ValidatorMap[seq[0]]
				if !ok {
					return Message{Message: fmt.Sprintf("Key %v not exist", key)}, false
				}

				var arg string
				if len(seq) == 1 {
					arg = ""
				} else {
					arg = seq[1]
				}

				if message, ok := fun(fmt.Sprintf("%v%v", name, field.Name), iface, arg); !ok {
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
