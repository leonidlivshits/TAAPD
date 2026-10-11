package main

import (
	"fmt"
	"io"
)

func TransmissionTime(channel Channel) float64 {
	data := channel.commonData()
	return 8 * data.MessageMB / data.SpeedMbps
}

func WriteTransmissionTimes(output io.Writer, channels ChannelContainer) error {
	for _, channel := range channels {
		if _, err := fmt.Fprintf(output, "Время передачи сообщения %q: %g с\n", channel.commonData().Name, TransmissionTime(channel)); err != nil {
			return err
		}
	}
	return nil
}
