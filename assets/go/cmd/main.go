package main

import (
	"fmt"
	"log"
	"math"
	"time"

	breker "github.com/BrianLeishman/breker.ninja/assets/go"
	"github.com/MichaelS11/go-ads"
)

func main() {
	err := ads.HostInit()
	if err != nil {
		panic(fmt.Errorf("failed to init ads: %w", err))
	}

	avg := breker.NewMovingAverage[int16](860)

	ads1, err := ads.NewADS("I2C1", uint16(0x4a), "")
	if err != nil {
		panic(err)
	}

	ads1.SetConfigGain(ads.ConfigGain4)
	ads1.SetConfigDataRate(ads.ConfigDataRate860)

	ads1.SetConfigInputMultiplexer(ads.ConfigInputMultiplexerDifferential01)

	var reading uint16

	recorder := time.NewTicker(time.Second / 860)
	go func() {
		for {
			select {
			case <-recorder.C:
				reading, err = ads1.Read()
				if err != nil {
					panic(err)
				}

				go avg.Record(int16(reading))
			}
		}
	}()

	reporter := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-reporter.C:
				readingVoltage := getVoltageFromReading(int16(avg.Average()), 1.024)
				v := 120
				a := getAmperage(readingVoltage)
				w := a * float64(v)
				log.Printf("%fV %fA %fW\n", readingVoltage, a, w)
			}
		}
	}()

	select {}
}

func getVoltageFromReading(n int16, max float64) float64 {
	return math.Abs(breker.ConvertRange(float64(int16(n)), 0, 65535, 0, max))
}

func getAmperage(v float64) float64 {
	return v / 30
}

// v = current * resistance
