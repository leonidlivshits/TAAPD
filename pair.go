package main

import "fmt"

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
