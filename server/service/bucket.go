package service

import (
	"errors"
	"hash/fnv"
	"strconv"
)

var B Bucket

type Bucket struct {
	m map[string]Board
}

type Board struct {
	IP   string
	Port string
	MAC  string
	ID   string
}

func InitBucket() {
	B = Bucket{
		m: make(map[string]Board),
	}
}

func (b *Bucket) Add(board Board) (int, error) {
	if b.Exist(board.IP, board.Port, board.MAC) {
		return 0, errors.New("board already exist")
	}
	hashValue := b.Hash(board.IP + board.Port + board.MAC)
	b.m[strconv.Itoa(int(hashValue))] = board
	return len(b.m), nil
}

func (b *Bucket) Get(ip, port, mac string) Board {
	hashValue := b.Hash(ip + port + mac)
	return b.m[strconv.Itoa(int(hashValue))]
}

func (b *Bucket) Del(ip, port, mac string) (int, error) {
	if b.Exist(ip, port, mac) {
		return 0, errors.New("board not exist")
	}
	hashValue := b.Hash(ip + port + mac)
	delete(b.m, strconv.Itoa(int(hashValue)))
	return len(b.m), nil
}

func (b *Bucket) Exist(ip, port, mac string) bool {
	hashValue := b.Hash(ip + port + mac)
	_, ok := b.m[strconv.Itoa(int(hashValue))]
	return ok
}

func (b *Bucket) Hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0
	}
	return h.Sum32()
}

func (b *Bucket) M2L() []Board {
	var l []Board
	for _, v := range B.m {
		l = append(l, v)
	}
	return l
}
