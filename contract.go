package main

//contract
type Property contract{

}

type Extends contract{

}

//incomplete type
type P

//use incomplete type as substitute
func (*p P) sayHello() {

}

func callsayHello() {
    contract(sayHello)
}
//type P based on X
type P(X)

type T contract{
    Property(string)
    Property(int32)
    Extends(type P)
}

type T contract{
    
}

type E contract{

}

type Hello struct {

}

func (t *T) say() {

}

func sayHello() {

}

func main() {
    i, am := T(sayHello)
}
