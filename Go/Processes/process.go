package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	defer  log.Println("Function end is coming",os.Getenv("WORKER"))
	
	if(len(os.Args) == 2) && (os.Args[1]=="outbound"){
		log.Println("Inside the outbound process")
		time.Sleep(1 * time.Second)

	}else {
	chln:= make(chan int)
	go func (chln chan int) {
		defer func() { chln <- 1 }()
		//create a channel
		cmd := exec.Command(os.Args[0], "outbound")
		cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"WORKER=outbound",    // this value is used
	)		
		// err := cmd.Start()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("waiting on Command ", )
		// err = cmd.Wait()
		out,err:= cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Command finished with output: %s", out)
		
	}(chln)
	os.Setenv("WORKER","inbound")
	<-chln
	}
	log.Println("We are coming to the end of thread - ",os.Getenv("WORKER"))
}