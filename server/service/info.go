package service

import "time"

func All() []Board {
	return B.M2L()
}

func Get(ip, port string) Board {
	return B.Get(ip, port)
}

func Probe(ID string) []Board {
	board, ok := B.m[ID]
	if !ok {
		return nil
	}
	var res []Board
	for _, b := range B.m {
		if b == board {
			continue
		}
		res = append(res, b)
	}
	board.ProbeStamp = time.Now()
	return res
}
