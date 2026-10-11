package main

import "fmt"

type FiberMode string

const (
	SingleMode FiberMode = "одномодовое"
	MultiMode  FiberMode = "многомодовое"
)

type OpticalChannel struct {
	ChannelData
	Mode       FiberMode
	FiberCount int
}

func (channel OpticalChannel) kindAndDetail() (string, string) {
	return "Оптоволоконный канал", fmt.Sprintf("вид волокна %s; число волокон %d", channel.Mode, channel.FiberCount)
}

var _ Channel = OpticalChannel{}
