package main

import (
	"fmt"
	fsm "go-fsm/src"
)

func main() {
	fmt.Println("enter main\n")
	defer fmt.Println("\nleave main")
	FsmTest()
}

func FsmTest() {
	doorCtrl := fsm.NewFSM(
		[]fsm.Transition{
			{"openDoor", []string{"Closed"}, "Opened"},
			{"closeDoor", []string{"Opened"}, "Closed"},
			{"lockDoor", []string{"UnLocked", "Closed"}, "Locked"},
			{"unLockDoor", []string{"Locked", "Opened"}, "UnLocked"},
		}, []fsm.CCallBack{
			onLeaveClosed,
			onLeaveOpened,
			onLeaveLocked,
			onLeaveUnlocked,
			onEnterClosed,
			onEnterOpened,
			onEnterLocked,
			onEnterUnlocked,
			onTransition,
			onFailed,
		},
		"Opened")


	doorCtrl.Trans("closeDoor")
	doorCtrl.Trans("closeDoor")
	doorCtrl.Trans("openDoor")
}

func onTransition(opt ...string) {
	//fmt.Println("onTransition", opt)
}

func onFailed(opt ...string) {
	fmt.Println("onFailed", opt)
}

func onLeaveClosed(opt ...string) {
	fmt.Println("onLeaveClosed", opt)
}

func onLeaveOpened(opt ...string) {
	fmt.Println("onLeaveOpened", opt)
}

func onLeaveLocked(opt ...string) {
	fmt.Println("onLeaveLocked", opt)
}

func onLeaveUnlocked(opt ...string) {
	fmt.Println("onLeaveUnlocked", opt)
}

func onEnterClosed(opt ...string) {
	fmt.Println("onEnterClosed", opt)
}

func onEnterOpened(opt ...string) {
	fmt.Println("onEnterOpened", opt)
}

func onEnterLocked(opt ...string) {
	fmt.Println("onEnterLocked", opt)
}

func onEnterUnlocked(opt ...string) {
	fmt.Println("onEnterUnlocked", opt)
}

