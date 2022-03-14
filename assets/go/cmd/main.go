package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/ads1x15"
	"periph.io/x/periph/host"

	tm "github.com/buger/goterm"

	ui "github.com/gizak/termui/v3"
)

func main() {
	tm.Clear()

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create a new ADS1115 ADC.
	adc, err := ads1x15.NewADS1115(bus, &ads1x15.Opts{
		I2cAddress: 0x4a,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Obtain an analog pin from the ADC.
	pin, err := adc.PinForChannel(ads1x15.Channel0Minus1, physic.Volt, 860*physic.Hertz, ads1x15.BestQuality)
	if err != nil {
		log.Fatalln(err)
	}
	defer pin.Halt()

	// Read values continuously from ADC.
	fmt.Println("Continuous reading")
	c := pin.ReadContinuous()

	records := make([]float64, 860)

	go func() {
		for {
			reading := <-c

			records = append(records[1:], float64(reading.V)/float64(physic.Volt))
		}
	}()

	chart := tm.NewLineChart(100, 20)

	data := new(tm.DataTable)
	data.AddColumn("Time")
	data.AddColumn("V")

	reporter := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-reporter.C

			tm.Println(chart.Draw(data))
		}
	}()

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

}
