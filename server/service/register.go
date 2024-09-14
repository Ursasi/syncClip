package service

import "strconv"

func GetOrAllocate(ip, port, mac string) (string, error) {
	exist := B.Exist(ip, port, mac)
	if exist {
		return B.Get(ip, port, mac).ID, nil
	}
	id, err := allocate(ip, port, mac)
	if err != nil {
		return "", err
	}
	return id, nil
}

func allocate(ip, port, mac string) (string, error) {
	hash := B.Hash(ip + port + mac)
	board := Board{
		IP:   ip,
		Port: port,
		MAC:  mac,
		ID:   strconv.Itoa(int(hash)),
	}
	_, err := B.Add(board)
	if err != nil {
		return "", err
	}
	return board.ID, nil
}
