package main

import llrss "github.com/SlyMarbo/rss"
import (
	"fmt"
	"time"
	"sync"
)

type Item struct {
	Index int
}

func (item Item) String() string {
	return fmt.Sprintf("Item %d", item.Index)
}

type Fetcher interface {
	Fetch() (items []*llrss.Item, next time.Time, err error)
	Domain() string
}

type RealFetcher struct {
	domain string
}

func (fetcher *RealFetcher) Fetch() (items []*llrss.Item, next time.Time, err error) {
	feed, err := llrss.Fetch(fetcher.domain)
	if err != nil {
		return nil, time.Now(), err
	}

	return feed.Items, feed.Refresh, nil
}

func (fetcher *RealFetcher) Domain() string {
	return fetcher.domain
}

type Subscription interface {
	Updates() <-chan Item
	Close() error
	Description() string
}

type SingleSubscription struct {
	updates chan Item
	done chan bool
	description string
}

func (sub *SingleSubscription) Updates() <- chan Item {
	return sub.updates
}

func (sub *SingleSubscription) Close() error {
	sub.done <- true
	return nil
}

func (sub *SingleSubscription) Description() string {
	return sub.description
}

// convert fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {

	sub := &SingleSubscription{
		updates: make(chan Item),
		done: make(chan bool),
		description: fetcher.Domain(),
	}

	go func() {

		defer func() {
			fmt.Println("Subscribe goroutine exit")
		}()

		for done := false; !done; {
			items, nextUpdate, err := fetcher.Fetch()
			fmt.Println("Fetched")
			if err != nil {
				panic(err)
			}

			nextUpdate = time.Now().Add(1 * time.Second)

			for i, _ := range items {
				sub.updates <- Item{i}
			}

			timer := time.NewTimer(time.Until(nextUpdate))
			select {
				case <-timer.C:
				case <-sub.done:
					fmt.Printf("Fetcher %s received close\n", fetcher.Domain())
					close(sub.updates)
					done = true
			}
		}

	}()

	return sub
}

// merges several streams
func Merge(subs ...Subscription) Subscription {
	merged := &MergedSubscription{
		subs: subs,
		updates: make(chan Item),
		description: "merged feed",
	}

	go func() {
		var wg sync.WaitGroup
		for _, sub := range subs {
			wg.Add(1)
			go func(sub Subscription) {
				defer wg.Done()
				fmt.Printf("Merged goroutine start for %s\n", sub.Description())
				for item := range sub.Updates() {
					merged.updates <- item
				}
				fmt.Printf("Merged goroutine exit for %s\n", sub.Description())
			}(sub)
		}

		wg.Wait()
		close(merged.updates)
	}()
	 
	return merged
}

type MergedSubscription struct {
	subs []Subscription
	updates chan Item
	description string
}

func (merged *MergedSubscription) Updates() <-chan Item {
	return merged.updates
}

func (merged *MergedSubscription) Close() error {
	for _, sub := range merged.subs {
		sub.Close()
	}

	// merged.done <- 1

	return nil
}

func (merged *MergedSubscription) Description() string {
	return merged.description
}

func Fetch(domain string) Fetcher {
	return &RealFetcher{domain}
}
 
func main() {
	domain0 := "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"
	domain1 := "https://blog.golang.org/feed.atom?format=xml"

	var merged Subscription
	if true {
		merged = Merge(
			Subscribe(Fetch(domain0)),
			Subscribe(Fetch(domain1)),
		)
	} else {
		merged = Subscribe(Fetch(domain0))
	}

	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("closing")
		fmt.Println("closed:, ", merged.Close())
	})

	for item := range merged.Updates() {
		if false { fmt.Println(item); }
	}
}

