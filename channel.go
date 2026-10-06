package main

type Channel interface {
	isChannel()
}

type ChannelData struct {
	Name      string
	SpeedMbps float64
	MessageMB float64
}

type WiredChannel struct {
	ChannelData
	CableLengthM float64
}

func (WiredChannel) isChannel() {}

type RadioChannel struct {
	ChannelData
	FrequencyMHz float64
}

func (RadioChannel) isChannel() {}

var (
	_ Channel = WiredChannel{}
	_ Channel = RadioChannel{}
)
