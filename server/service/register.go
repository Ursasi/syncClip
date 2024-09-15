package service

import "strconv"

func GetOrAllocate(ip, port string) (string, error) {
	exist := B.Exist(ip, port)
	if exist {
		return B.Get(ip, port).ID, nil
	}
	id, err := allocate(ip, port)
	if err != nil {
		return "", err
	}
	return id, nil
}

func allocate(ip, port string) (string, error) {
	hash := Hash(ip + port)
	board := Board{
		IP:   ip,
		Port: port,
		ID:   strconv.Itoa(int(hash)),
	}
	_, err := B.Add(board)
	if err != nil {
		return "", err
	}
	return board.ID, nil
}
