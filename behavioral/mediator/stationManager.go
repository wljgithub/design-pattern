package main

import "sync"

type stationManager struct {
	isPlatformFree bool
	lock *sync.Mutex
	trainQueue []train
}

func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (this *stationManager) canLand(t train) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.isPlatformFree{
		this.isPlatformFree = false
		return true
	}
	this.trainQueue = append(this.trainQueue,t)
	return false
}

func (this *stationManager) notifyFree() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if !this.isPlatformFree{
		this.isPlatformFree = true
	}
	if len(this.trainQueue) > 0 {
		firstTrainInQueue := this.trainQueue[0]
		this.trainQueue = this.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
