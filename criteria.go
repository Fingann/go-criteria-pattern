package main

type Criteria interface {
	Satisfied(user User) bool
}

type AgeCriteria struct {
	Age int
}

func (a *AgeCriteria) Satisfied(user User) bool {
	return user.Age == a.Age
}
