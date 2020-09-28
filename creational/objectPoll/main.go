package main

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]iPollObject,0)
	for i:=0;i <3 ;i ++{
		c := &connection{id:strconv.Itoa(i)}
		connections = append(connections,c)
	}
	poll , err := initPoll(connections)
	if err != nil {
		log.Fatalf("Init poll error: %s",err)
	}

	conn1 ,err := poll.loan()
	if err != nil {
		log.Fatalf("Poll loan error: %s",err)
	}

	conn2 , err := poll.loan()
	if err != nil {
		log.Fatalf("Poll loan error: %s",err)
	}

	poll.receive(conn1)
	poll.receive(conn2)
}
