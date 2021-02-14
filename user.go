package main

type User struct {
	Name      string
	Age       int
	HigherEducation bool
	Languages []string
}

func (u User) IsAdult() bool  {
	return u.Age >= 18
}

func NewUser(name string, age int, higherEducation bool, languages []string) *User {
	return &User{
		Name:      name,
		Age:       age,
		HigherEducation: higherEducation,
		Languages:  languages,
	}
}
