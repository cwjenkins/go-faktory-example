package main

import (
	"log"
	"time"
	"context"
	worker "github.com/contribsys/faktory_worker_go"

	"go-faktory-example.id.me/config"
	"go-faktory-example.id.me/user"
)

// The actual func to work/perform
func perform(ctx context.Context, args ...interface{}) (error) {
	t1 := time.Now()
	help := worker.HelperFor(ctx)
	log.Println("GoWorker JID:", help.Jid(), "INFO: start")

	var u user.User
	config.DB().First(&u)
	log.Println("GoWorker JID:", help.Jid(), u.FullName)
	log.Println("GoWorker JID:", help.Jid(), "INFO: done:", time.Now().Sub(t1))

	return nil
}

func main() {
        log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	config.Init()
	mgr := worker.NewManager()

	// register job types and the function to execute them
	mgr.Register("GoWorker", perform)
	//mgr.Register("AnotherJob", anotherFunc)

	// use up to N goroutines to execute jobs
	// TODO: read on why this is recommended to be so low
	// goroutines can scale thousands/millions
	// bottleneck should be external:
	// Examples: underlying nix somaxconn/ulimit, DB max conn, API rate limit, etc.
	mgr.Concurrency = 20

	// pull jobs from these queues, in this order of precedence
	mgr.ProcessStrictPriorityQueues("for_go")

	// alternatively you can use weights to avoid starvation
	//mgr.ProcessWeightedPriorityQueues(map[string]int{"critical":3, "default":2, "bulk":1})

	// Start processing jobs, this method does not return.
	mgr.Run()
}
