package main

import (
	"fmt"
	"sync"
)

type poll struct {
	idle     []iPollObject
	active   []iPollObject
	capacity int
	mulock   *sync.Mutex
}

func initPoll(pollObjects []iPollObject) (*poll, error) {
	if len(pollObjects) == 0 {
		return nil,fmt.Errorf("Cannot create a poll of 0 length")
	}
	active := make([]iPollObject,0)
	poll := &poll{
		idle: pollObjects,
		active: active,
		capacity: len(pollObjects),
		mulock: new(sync.Mutex),
	}
	return poll,nil
}

func (p *poll) loan() (iPollObject,error){
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0{
		return nil,fmt.Errorf("No poll object free.Please request after some time")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active,obj)
	fmt.Printf("Loan poll object with ID:%s\n",obj.getId())
	return obj,nil
}

func (p *poll) receive(target iPollObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle,target)
	fmt.Printf("Return poll object with ID: %s\n",target.getId())
	return nil
}

func (p *poll) remove(target iPollObject) error {
	currentActiveLength := len(p.active)
	for i,obj := range p.active{
		if obj.getId() == target.getId(){
			p.active[currentActiveLength-1],p.active[i] = p.active[i],p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Target poll Oject doesn't belong to the poll")
}