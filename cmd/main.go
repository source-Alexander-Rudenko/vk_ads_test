package main

import (
	"fmt"
	"github.com/source-Alexander-Rudenko/vk_ads_test/pkg/pool"
	"os"
	"time"
)

func main() {
	p := pool.NewPool(10, os.Stdout)
	p.Run()

	for i := 0; i < 3; i++ {
		if err := p.AddWorker(); err != nil {
			fmt.Println("AddWorker error:", err)
		}
	}

	for i := 0; i < 5; i++ {
		task := fmt.Sprintf("task-%d", i)
		if err := p.Submit(task); err != nil {
			fmt.Println("Submit error:", err)
		}
	}

	time.Sleep(time.Second)

	if err := p.RemoveWorker(1); err != nil {
		fmt.Println("RemoveWorker error:", err)
	}

	if err := p.Submit("task-after-remove"); err != nil {
		fmt.Println("Submit error:", err)
	}

	time.Sleep(time.Second)

	p.Shutdown()
	time.Sleep(500 * time.Millisecond)
}
