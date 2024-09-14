package service

func All() []Board {
	return B.M2L()
}

func Get(ip, port, mac string) Board {
	return B.Get(ip, port, mac)
}
