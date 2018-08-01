package main

import (
	"fmt"
	"time"
)

func MakeTimers() string {
	//create a new Timer, set time limit
	timer1 := time.NewTimer(2 * time.Second)
	//timer sets a channel that will be notified at the setted time
	//send a msg over the channel C
	//waiting for a msg from that channel, during waiting it is blocked
	<-timer1.C
	return fmt.Sprintf("Timer 1 expired\n")
}

func MakeTimersStop() string {
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Printf("Timer 2 expired\n")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		return fmt.Sprintf("Timer 2 stopped\n")
	}
	return ""
}

func TimerLoopExample() {
	timer := time.NewTimer(2 * time.Second)
	correct := 0
	problems := Parse()
	for i, problem := range problems {
		fmt.Printf("Question #%d: %s \n", i+1, problem.Question)

		answerChan := make(chan string)
		//run a goroutine, func(){} is an anonymous function
		go func() {
			//scan an answer
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer //turn the anonymous function into a closure, which takes data outside the function
		}() //() here means calling the function
		select {
		//select has no default, so that it is always waiting for a msg from different channels
		case <-timer.C:
			fmt.Printf("You scored %d of %d\n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.Answer {
				correct++
			}
		}
	}
}
