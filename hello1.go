package main

import (
	"fmt"
	"sync"
	"time"

	"hello/morestrings"
	"hello/pubsub"
)

const topic = "topic"

func Example() {
	ps := pubsub.New(100)
	ch := ps.Sub(topic)
	ch1 := ps.Sub(topic)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		publish(ps)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; ; i++ {
			if i == 5 {
				// See the documentation of Unsub for why it is called in a new
				// goroutine.
				go ps.Unsub(ch, "topic")
			}

			if msg, ok := <-ch; ok {
				// fmt.Printf("ch Received %s, %d times.\n", msg, i)
				fmt.Println("ch:", time.Now(), " msg:", msg)
				// fmt.Println(time.Now())
			} else {
				break
			}
		}
		// Output:
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 1; ; i++ {
			if i == 10 {
				// See the documentation of Unsub for why it is called in a new
				// goroutine.
				go ps.Unsub(ch1, "topic")
			}
			if msg, ok := <-ch1; ok {
				// fmt.Printf("ch1 Received %s, %d times.\n", msg, i)
				fmt.Println("ch1:", time.Now(), " msg:", msg, " i:", i)
			} else {
				break
			}
		}
		// Output:
	}(&wg)

	wg.Wait()

	// Received message, 1 times.
	// Received message, 2 times.
	// Received message, 3 times.
	// Received message, 4 times.
	// Received message, 5 times.
	// Received message, 6 times.
}

func publish(ps *pubsub.PubSub) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Println("send:", time.Now())
		ps.Pub("message", topic)
	}
}
func main() {
	Example()

	fmt.Println(morestrings.ReverseRunes("!11G ,zzz"))

	// bufio.NewReader(os.Stdin)
	// bufio.NewReader(os.Stdin).ReadString('\n')
	// input, _ := reader.ReadString('\n')

	// fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))

}
