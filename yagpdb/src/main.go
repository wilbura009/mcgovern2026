package main
import (
    "log"
    "text/template"
)

type User struct {
    name string
}

func main() {
    // Define the template
    const welcomeMsg = `
Welcome {{.name}}!
`
    // Set the data to be inserted into the template
    var user = []User{
        {"Guy"},
    }
    println(user)
}
