package main

import "fmt"

type Observer interface{
    Update(float32,float32,float32)
}

type Subject interface{
    RegisterObserver(Observer)
    RemoveObserver(Observer)
    NotifyObservers()
}

type DisplayElement interface{
    Display()
}

type CurrentConditionDisplay struct{
    Temperature float32
    Humidity float32
    Subject Subject
}

func (c *CurrentConditionDisplay) Display (){
    fmt.Println("Current Condition: ",c.Temperature, c.Humidity)
}

func (c *CurrentConditionDisplay) Update (temp, humi, press float32){
    fmt.Println("4. Update received...")
    c.Temperature = temp
    c.Humidity = humi

    c.Display()
}

func NewCurrentConditionDisplay (w *WaetherData) *CurrentConditionDisplay{
    ncd := &CurrentConditionDisplay{0.0,0.0,w}
    ncd.Subject.RegisterObserver(ncd)
    return ncd
}

type WaetherData struct{
        Observers []*Observer
        Humidity float32
        Temperature float32
        Pressure float32
}

func (w *WaetherData) RegisterObserver(o Observer){
    w.Observers = append(w.Observers, &o)
}

func (w *WaetherData) RemoveObserver (o Observer){
    j := 0

    for _, n := range w.Observers {
        if n != &o {
            w.Observers[j] = n
            j++
        }
    }
    w.Observers = w.Observers[:j]
}

func (w *WaetherData) NotifyObservers(){
        fmt.Println("3. notifying observers: ", w.Observers)
        for _ , o := range w.Observers{
             fmt.Println("Notifying : ", o)
             sub := *o
             sub.Update(w.Temperature, w.Humidity, w.Pressure)
        }
}


func (w *WaetherData) MeasurementsChanged(){
        fmt.Println("2. Processing new measurements")
        w.NotifyObservers()
}

func (w *WaetherData) SetMeasurements( temp, humi, press float32){
    w.Humidity = humi
    w.Temperature = temp
    w.Pressure = press
    fmt.Println("1. Received new Measurements")
    w.MeasurementsChanged()
}

func main(){
   wd := WaetherData{}

   cd := NewCurrentConditionDisplay(&wd)
   cd.Display()
   wd.SetMeasurements(11,22,12)
}

