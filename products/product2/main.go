package main

import (
	"gitlab.com/tsuchinaga/monorepo-test/products/libs/scheduler"
	"log"
)

func main() {
	s := scheduler.NewScheduler()
	s.AddJob("greet", "* * * * *", new(greet))
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
	defer s.Stop()

	<-make(chan int)
}

type greet struct{}

func (j *greet) Run() {
	log.Println("こんにちわーるど")
}
