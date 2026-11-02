package main

import "io"

func WriteWiredChannels(output io.Writer, channels ChannelContainer) error {
	var wired ChannelContainer
	for _, channel := range channels {
		if _, ok := channel.(WiredChannel); ok {
			wired.Add(channel)
		}
	}
	return WriteChannels(output, wired)
}
