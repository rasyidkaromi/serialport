# serialport
Windows Serial Port with GO

Install
---------------
go get github.com/rasyidkaromi/serialport


Usage
---------------


        package main

        import (
          "log"
          "github.com/rasyidkaromi/serialport"
          "fmt"
          "time"
        )


        func main (){
          ports, err := serialport.GetPortsList()

          if err != nil{
            log.Fatal(err)
          }

          log.Println("Port mana yang ingin Anda gunakan?")

          for _, port := range ports{
            fmt.Println("*", port)
          }

          var portt string

          fmt.Scanln(&portt)

          log.Println("Port:", portt)

          mode := &serialport.Mode{
            BaudRate: 115200,
          }

          connection := serialport.NewConnection(portt, mode)

          defer connection.Close()

          go func (){
            buffer := make([]byte, 100)
            for {
              time.Sleep(1000 * time.Millisecond)
              n, buf, text, isOk := connection.Read(&buffer)

              if isOk{
                fmt.Println(buf)
                fmt.Println(n)
                fmt.Println(text)
              }
            }
          } ()

          var text string
          for {
            fmt.Scanln(&text)

            connection.Write(text)
          }
        }

