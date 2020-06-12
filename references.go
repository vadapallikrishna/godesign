package main

ref S

type I interface {
   func Hello() string
}

type E struct {

}

//immutable 

func main()  {
    ref S interface
    S = "456"
    S = S + 789
    S.Hello()
    
}
