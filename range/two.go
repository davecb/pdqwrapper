type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 30},
        {Name: "Charlie", Age: 35},
    }

    // Range by value
    for _, person := range people {
        fmt.Printf("%s is %d years old\n", person.Name, person.Age)
    }

    // Range by index and value
    for i, person := range people {
        fmt.Printf("Person %d: %s\n", i, person.Name)
    }

    // Range with pointer receiver
    for i := range people {
        people[i].Age++ // Modify the actual slice element
    }
}
