package main

import ("fmt")

type FlyBehavior interface {
	Fly()
}

type Duck struct {
	FlyBehavior FlyBehavior
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

type FlyNoWay struct{}

func (FlyNoWay) Fly(){
        fmt.Println("I can't fly")
}

func main(){
    duck := Duck{FlyNoWay{}}
    duck.PerformFly()
}

