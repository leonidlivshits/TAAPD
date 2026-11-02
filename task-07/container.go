package main

type ChannelContainer []Channel

func (container *ChannelContainer) Add(channel Channel) {
	*container = append(*container, channel)
}
