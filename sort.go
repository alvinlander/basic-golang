package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}
type UserSlice []User

func (user UserSlice) Len() int{
	return len(user)
}

func (user UserSlice) Less(i,j int) bool{
	return user[i].Age < user[j].Age
}
func (user UserSlice) Swap(i,j int){
	user[i],user[j] = user[j],user[i]
}
func main() {

	Users := []User{
		{"Lans",22},
		{"Eko",14},
		{"Budi",25},
		{"Rudi",21},
	}
	sort.Sort(UserSlice(Users))
	fmt.Println(Users)

}