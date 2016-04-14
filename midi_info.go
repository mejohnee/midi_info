package main

import (
    "github.com/rakyll/portmidi"
    "fmt"
)

func main() {
    portmidi.Initialize()
    //portmidi.CountDevices() // returns the number of MIDI devices
    //var i portmidi.DeviceID
    fmt.Println("ID\tInterface\tName\t\t\tInput\t\tOutput\t\tOpened")
    for i := 0; i < portmidi.CountDevices(); i++ {
        
        device := portmidi.Info(portmidi.DeviceID(i))
        fmt.Printf("%d\t" +device.Interface + "\t\t" + device.Name + "\t", i)
        if(device.IsInputAvailable){
            fmt.Print("available\t")
        } else {
            fmt.Print("not available\t")
        }

        if(device.IsOutputAvailable){
            fmt.Print("available\t")
        } else {
            fmt.Print("not available\t")
        }

        if(device.IsOpened){
            fmt.Print("yes\t")
        } else {
            fmt.Print("no\t")
        }

        if(int(portmidi.DefaultInputDeviceID()) == i || int(portmidi.DefaultOutputDeviceID()) == i) {
            fmt.Print("default")
        }

        fmt.Println()
    }
}

