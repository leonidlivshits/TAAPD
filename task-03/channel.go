package main

type Channel interface {
	commonData() ChannelData
	kindAndDetail() (string, string)
}

type ChannelData struct {
	Name      string
	SpeedMbps float64
	MessageMB float64
}

func (data ChannelData) commonData() ChannelData {
	return data
}

type WiredChannel struct {
	ChannelData
	CableLengthM float64
}

type RadioChannel struct {
	ChannelData
	FrequencyMHz float64
}

var (
	_ Channel = WiredChannel{}
	_ Channel = RadioChannel{}
)
