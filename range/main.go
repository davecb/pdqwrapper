package main

import "fmt"

func main() {
    list := &List{items: []interface{}{1, 2, 3}}
    
    // Using the custom Range
    // for item := range list.Range() {
    //     fmt.Println(item)
    // }
    
    // Standard Go range
    for i, item := range list.items {
        fmt.Printf("item[%d] = %d\n", i,item)
    }
}


type List struct {
    items []interface{}
}

