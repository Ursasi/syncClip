package util

import (
	"fmt"
	"syncClip/peer"
	"syncClip/server/service"
)

func WatchBoard() {
	displayBoard(peer.PeerState.Boards)
	for {
		peer.PeerState.Cond.Wait()
		displayBoard(peer.PeerState.Boards)
	}
}

func displayBoard(board []service.Board) {
	fmt.Println("IP\tPORT\tID")
	for _, v := range board {
		fmt.Printf("%s\t%d\t%s\n", v.IP, v.Port, v.ID)
	}
}
