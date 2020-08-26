# go-fsm
A simple finite state machine for Golang 


# Basic Example
From examples/fsm-test.go:

```go

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

```

will get output like this:

```go
enter main

onLeaveOpened [Opened  closeDoor]
onEnterClosed [Closed Opened closeDoor]
onFailed [Closed Opened closeDoor]
onLeaveClosed [Closed Opened openDoor]
onEnterOpened [Opened Closed openDoor]

leave main


```