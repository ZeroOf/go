package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond fater %s", url, timeout)
}

func main()  {
	url := "www.baidu.cn"
	if err := WaitForServer(url); err != nil {
		log.Fatalf("site is down : %v\n", err)
	}

	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		fmt.Errorf("failed to create temprary dir: %v", err)
		return
	}
	os.RemoveAll(dir)
}
