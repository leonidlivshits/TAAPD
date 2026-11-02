package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type channelRecord struct {
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	SpeedMbps    *float64 `json:"speed_mbps"`
	MessageMB    *float64 `json:"message_mb"`
	CableLengthM *float64 `json:"cable_length_m"`
	FrequencyMHz *float64 `json:"frequency_mhz"`
}

func LoadChannels(path string) (ChannelContainer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("открыть %q: %w", path, err)
	}
	defer file.Close()

	channels, err := ReadChannels(file)
	if err != nil {
		return nil, fmt.Errorf("файл %q: %w", path, err)
	}
	return channels, nil
}

func ReadChannels(input io.Reader) (ChannelContainer, error) {
	var channels ChannelContainer
	decoder := json.NewDecoder(input)
	decoder.DisallowUnknownFields()

	for number := 1; ; number++ {
		var record channelRecord
		err := decoder.Decode(&record)
		if err == io.EOF {
			return channels, nil
		}
		if err != nil {
			return nil, fmt.Errorf("запись %d: %w", number, err)
		}

		channel, err := record.toChannel()
		if err != nil {
			return nil, fmt.Errorf("запись %d: %w", number, err)
		}
		channels.Add(channel)
	}
}

func (record channelRecord) toChannel() (Channel, error) {
	if strings.TrimSpace(record.Name) == "" {
		return nil, fmt.Errorf("name: нужно непустое имя канала")
	}
	if record.SpeedMbps == nil || *record.SpeedMbps <= 0 {
		return nil, fmt.Errorf("speed_mbps: нужна положительная скорость")
	}
	if record.MessageMB == nil || *record.MessageMB < 0 {
		return nil, fmt.Errorf("message_mb: нужен неотрицательный объём")
	}

	data := ChannelData{
		Name:      record.Name,
		SpeedMbps: *record.SpeedMbps,
		MessageMB: *record.MessageMB,
	}
	switch record.Type {
	case "wired":
		if record.CableLengthM == nil || *record.CableLengthM < 0 {
			return nil, fmt.Errorf("cable_length_m: нужна неотрицательная длина кабеля")
		}
		if record.FrequencyMHz != nil {
			return nil, fmt.Errorf("frequency_mhz: поле не относится к проводному каналу")
		}
		return WiredChannel{ChannelData: data, CableLengthM: *record.CableLengthM}, nil
	case "radio":
		if record.FrequencyMHz == nil || *record.FrequencyMHz <= 0 {
			return nil, fmt.Errorf("frequency_mhz: нужна положительная частота")
		}
		if record.CableLengthM != nil {
			return nil, fmt.Errorf("cable_length_m: поле не относится к радиоканалу")
		}
		return RadioChannel{ChannelData: data, FrequencyMHz: *record.FrequencyMHz}, nil
	default:
		return nil, fmt.Errorf("неизвестный вид канала %q", record.Type)
	}
}
