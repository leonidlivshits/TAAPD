package main

func TransmissionTime(channel Channel) float64 {
	data := channel.commonData()
	return 8 * data.MessageMB / data.SpeedMbps
}
