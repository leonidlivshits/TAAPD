package main

import (
	"fmt"
	"io"
)

func PairPhrase(first, second Channel) (string, error) {
	switch first.(type) {
	case WiredChannel:
		switch second.(type) {
		case WiredChannel:
			return "Встретились два проводных канала.", nil
		case RadioChannel:
			return "Проводной канал идёт перед радиоканалом.", nil
		}
	case RadioChannel:
		switch second.(type) {
		case WiredChannel:
			return "Сначала радиоканал, за ним проводной канал.", nil
		case RadioChannel:
			return "Пара радиоканалов оказалась рядом.", nil
		}
	}
	return "", fmt.Errorf("неподдерживаемая пара каналов: %T и %T", first, second)
}

func WritePair(output io.Writer, first, second Channel) error {
	phrase, err := PairPhrase(first, second)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(output, "Мультиметод для %T и %T: %s\n", first, second, phrase); err != nil {
		return err
	}
	return WriteChannels(output, ChannelContainer{first, second})
}

func WriteAllPairs(output io.Writer, channels ChannelContainer) error {
	if _, err := fmt.Fprintf(output, "Всего пар: %d\n", len(channels)*len(channels)); err != nil {
		return err
	}
	for _, first := range channels {
		for _, second := range channels {
			if err := WritePair(output, first, second); err != nil {
				return err
			}
		}
	}
	return nil
}
