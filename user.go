package main

type Color int

const (
	Brown Color = iota + 1
	Black
	Green
	Blue
)

type User struct {
	Name      string
	Age       int
	HairColor Color
	EyeColor  Color
}

func NewUser(name string, age int, hair, eye Color) *User {
	return &User{
		Name:      name,
		Age:       age,
		HairColor: hair,
		EyeColor:  eye,
	}
}
