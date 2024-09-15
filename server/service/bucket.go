package service

import (
	"errors"
	"hash/fnv"
	"strconv"
	constant "syncClip"
	"time"
)

var B Bucket

type Bucket struct {
	m map[string]Board
}

type Board struct {
	IP         string
	Port       string
	ID         string
	ProbeStamp time.Time
}

func InitBucket() {
	B = Bucket{
		m: make(map[string]Board),
	}
}

func (b *Bucket) Add(board Board) (int, error) {
	if b.Exist(board.IP, board.Port) {
		return 0, errors.New("board already exist")
	}
	hashValue := Hash(board.IP + board.Port)
	b.m[strconv.Itoa(int(hashValue))] = board
	return len(b.m), nil
}

func (b *Bucket) Get(ip, port string) Board {
	hashValue := Hash(ip + port)
	return b.m[strconv.Itoa(int(hashValue))]
}

func (b *Bucket) Del(ip, port string) (int, error) {
	if b.Exist(ip, port) {
		return 0, errors.New("board not exist")
	}
	hashValue := Hash(ip + port)
	delete(b.m, strconv.Itoa(int(hashValue)))
	return len(b.m), nil
}

func (b *Bucket) Exist(ip, port string) bool {
	hashValue := Hash(ip + port)
	_, ok := b.m[strconv.Itoa(int(hashValue))]
	return ok
}

func Hash(s string) uint32 {
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

func (b *Bucket) L2M(l []Board) {
	for _, v := range l {
		hashValue := Hash(v.IP + v.Port)
		b.m[strconv.Itoa(int(hashValue))] = v
	}
}

func (b *Bucket) Clean() {
	for k, v := range b.m {
		if time.Since(v.ProbeStamp) > constant.DefaultProbeInterval {
			delete(b.m, k)
		}
	}
}
