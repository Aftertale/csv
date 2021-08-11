package main

import (
	"reflect"
	"testing"
)

func newUser() User {
	return User{
		Name:   "BillyBob",
		Age:    49,
		HasPet: false,
	}
}

func newUser2() User {
	return User{
		Name:   "Joe",
		Age:    72,
		HasPet: true,
	}
}

func TestMarshalHeader(t *testing.T) {
	user := newUser()

	want := []string{"name", "age", "has_pet"}

	got := marshalHeader(reflect.TypeOf(user))

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestMarshalOne(t *testing.T) {
	user := newUser()

	want := []string{"BillyBob", "49", "false"}
	got, err := marshalOne(reflect.ValueOf(user))
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestMarshal(t *testing.T) {
	user := []User{newUser(), newUser2(), {
		Name:   "Dillon",
		Age:    1000,
		HasPet: true,
	}}

	want := [][]string{{"name", "age", "has_pet"}, {"BillyBob", "49", "false"}, {"Joe", "72", "true"}, {"Dillon", "1000", "true"}}
	got, err := Marshal(user)
	if err != nil {
		t.Error("Got error when trying to marshal", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}
