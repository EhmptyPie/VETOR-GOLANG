package main

import ("fmt"
"strings"
"sort"
)

func main(){
	greeting := "Hello my friends!"
    fmt.Println(strings.Contains(greeting, "friends"))
    fmt.Println(strings.ReplaceAll(greeting, 
        "Hello", "Hi"))
        fmt.Println(strings.ToUpper(greeting))
        fmt.Println(strings.Index(greeting,"my"))
        fmt.Println(strings.Split(greeting, " "))
        ages := []int {50, 70, 20}
        sort.Ints(ages)
        fmt.Println(ages)
        index := sort.SearchInts(ages, 50)
        fmt.Println(index)
        names := []string{"Azul", "Verde", "Vermelho"}
        sort.Strings(names)
        fmt.Println(names)
        fmt.Println(sort.SearchStrings(names, "Azul"))

}
