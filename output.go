package main

import (
	"fmt"
	"io"
)

func WriteChannels(output io.Writer, channels ChannelContainer) error {
	if _, err := fmt.Fprintf(output, "Всего каналов: %d\n", len(channels)); err != nil {
		return err
	}

	for i, item := range channels {
		data := item.commonData()
		kind, detail := item.kindAndDetail()
		if _, err := fmt.Fprintf(output,
			"%d. %s %q: скорость %g Мбит/с; объём сообщения %g Мбайт; %s\n",
			i+1, kind, data.Name, data.SpeedMbps, data.MessageMB, detail); err != nil {
			return err
		}
	}
	return nil
}

func (channel WiredChannel) kindAndDetail() (string, string) {
	return "Проводной канал", fmt.Sprintf("длина кабеля %g м", channel.CableLengthM)
}

func (channel RadioChannel) kindAndDetail() (string, string) {
	return "Радиоканал", fmt.Sprintf("частота %g МГц", channel.FrequencyMHz)
}
