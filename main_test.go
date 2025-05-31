package main

import (
	"testing"

	"github.com/lubie-placki-be/configs"
)

type Address struct {
	Street string `validate:"required"`
}

type Pet struct {
	Name string
}

type User struct {
	Name        string  `validate:"required,minStringLength:3"`
	HomeAddress Address `validate:"deep"`
	WorkAddress Address
	Pets        []Pet `validate:"maxArrayLength:2"`
}

func Test_ValidateField_Required(t *testing.T) {
	user := User{}

	result, ok := configs.Validate(user)
	if ok || result.Key != "required" {
		t.Errorf(`Validate User.Name should fail (required)`)
	}
}

func Test_ValidateField_MinStringLength(t *testing.T) {
	user := User{Name: "a"}

	result, ok := configs.Validate(user)
	if ok || result.Key != "minStringLength" {
		t.Errorf(`Validate User.Name should fail (minStringLength)`)
	}
}

func Test_ValidateNested_Deep(t *testing.T) {
	user := User{Name: "Adam", HomeAddress: Address{}}

	_, ok := configs.Validate(user)
	if ok {
		t.Errorf(`Validate User.HomeAddress.Street should fail`)
	}
}

func Test_ValidateNested_NotDeep(t *testing.T) {
	user := User{Name: "Adam", HomeAddress: Address{Street: "Długa"}, WorkAddress: Address{}}

	_, ok := configs.Validate(user)
	if !ok {
		t.Errorf(`Validate User.WorkAddress.Street shouldn't fail`)
	}
}

func Test_ValidateArray(t *testing.T) {
	user := User{Name: "Adam", HomeAddress: Address{Street: "Długa"}, WorkAddress: Address{}, Pets: []Pet{{}, {}, {}}}

	result, ok := configs.Validate(user)
	if ok || result.Key != "maxArrayLength" {
		t.Errorf(`Validate User.Pets shouldn't fail (maxArrayLength)`)
	}
}
