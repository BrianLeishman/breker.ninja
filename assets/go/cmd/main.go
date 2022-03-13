package main

import (
	"fmt"
	"log"
	"time"

	breker "github.com/BrianLeishman/breker.ninja/assets/go"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/ads1x15"
	"periph.io/x/periph/host"
)

func main() {
	// if err := ui.Init(); err != nil {
	// 	log.Fatalf("failed to initialize termui: %v", err)
	// }
	// defer ui.Close()

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
	pin, err := adc.PinForChannel(ads1x15.Channel0Minus1, 1024*physic.MilliVolt, 860*physic.Hertz, ads1x15.BestQuality)
	if err != nil {
		log.Fatalln(err)
	}
	defer pin.Halt()

	// Read values continuously from ADC.
	fmt.Println("Continuous reading")
	c := pin.ReadContinuous()

	// avg := breker.NewMovingAverage[physic.ElectricPotential](1000)

	var readings [860]physic.ElectricPotential

	start := time.Now()

	// stop := make(chan (struct{}))
	// go func() {
	for i := 0; i < len(readings); i++ {
		reading := <-c
		readings[i] = reading.V
	}
	// }()

	// time.Sleep(time.Second)
	// stop <- struct{}{}

	duration := time.Since(start)

	// sl := avg.Slice()
	log.Println(readings)
	log.Println("took", duration)

	// return

	// // avg := breker.NewMovingAverage[float64](860)

	// ads1, err := breker.NewADS("I2C1", uint16(0x4a), "")
	// if err != nil {
	// 	panic(err)
	// }

	// ads1.SetConfigGain(breker.ConfigGain4)
	// ads1.SetConfigDataRate(breker.ConfigDataRate860)

	// ads1.SetConfigInputMultiplexer(breker.ConfigInputMultiplexerDifferential01)

	// if err := ads1.WriteConfig(); err != nil {
	// 	panic(err)
	// }

	// // wait for config to take effect?
	// time.Sleep(100 * time.Millisecond)

	// var reading uint16

	// start := time.Now()

	// // recorder := time.NewTicker(time.Second / 860)
	// // go func() {
	// for i := 0; i < 860; i++ {
	// 	// <-recorder.C

	// 	reading, err = ads1.Read()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	log.Println(reading)

	// 	// avg.Record(float64(int16(reading)))
	// 	time.Sleep(time.Second / 860)
	// }
	// // }()

	// log.Println(time.Since(start))
	// log.Println(avg.Slice())

	// return

	// // tm.Clear() // Clear current screen

	// p0 := widgets.NewPlot()
	// p0.Title = "Voltage Readings"
	// p0.SetRect(0, 0, 100, 15)
	// p0.AxesColor = ui.ColorWhite
	// p0.LineColors[0] = ui.ColorBlue

	// reporter := time.NewTicker(1 * time.Second)
	// go func() {
	// 	for {
	// 		<-reporter.C

	// 		p0.Data = [][]float64{avg.Slice()}
	// 	}
	// }()
	// time.Sleep(1 * time.Second)
	// ui.Render(p0)
	// fmt.Println(p0.Data, len(p0.Data[0]))

	// for e := range ui.PollEvents() {
	// 	if e.Type == ui.KeyboardEvent {
	// 		break
	// 	}
	// }
}

func getVoltageFromReading(n int16, max float64) float64 {
	return breker.ConvertRange(float64(int16(n)), 0, 65535, 0, max)
}

func getAmperage(v float64) float64 {
	return v / 30
}
