package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/dihedron/brokerd/log"
	"go.uber.org/zap"
)

const (
	// WAIT is the tim to wait between iterations.
	WAIT time.Duration = 50 * time.Millisecond
)

type generator func() string

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func random() string {
	return fmt.Sprintf("%d", rand.Int())
}

func exactly(value string) generator {
	return func() string {
		return value
	}
}

func get(id int, url string, key string, iter int, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	for i := 0; i < iter; i++ {
		time.Sleep(WAIT)
		err = func() error {
			var response *http.Response
			var err error
			var body []byte
			if response, err = http.Get(url + key); err != nil {
				return err
			}
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			m := map[string]string{}
			if err = json.Unmarshal(body, &m); err != nil {
				return err
			}
			if value, ok := m[key]; ok {
				fmt.Printf("- GET[%d]: %s => %s\n", id, key, value)
			}
			return nil
		}()
		if err != nil {
			log.L.Error("error getting value", zap.Int("id", id), zap.Error(err))
			break
		}
	}
	return err
}

func set(id int, url string, key string, value generator, iter int, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	for i := 0; i < iter; i++ {
		time.Sleep(WAIT)
		err = func() error {
			var response *http.Response
			var err error
			var body []byte
			value := value()
			m := map[string]string{
				key: value,
			}
			if body, err = json.Marshal(m); err != nil {
				return err
			}
			if response, err = http.Post(url, "application/json", bytes.NewBuffer(body)); err != nil {
				return err
			}
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			fmt.Printf("+ SET[%d]: %s => %s\n", id, key, value)
			return nil
		}()
		if err != nil {
			log.L.Error("error setting value", zap.Int("id", id), zap.Error(err))
			break
		}
	}
	return err
}

func main() {
	log.L.Sync()
	var wg sync.WaitGroup
	const iterations int = 500

	wg.Add(6)
	// go set(0, "http://localhost:11000/key/", "foo", exactly("bar"), 10, &wg)
	go set(0, "http://localhost:11000/key/", "foo", random, iterations, &wg)
	go get(0, "http://localhost:11000/key/", "foo", iterations, &wg)
	go get(1, "http://localhost:11001/key/", "foo", iterations, &wg)
	go get(2, "http://localhost:11002/key/", "foo", iterations, &wg)
	go get(3, "http://localhost:11003/key/", "foo", iterations, &wg)
	go get(4, "http://localhost:11004/key/", "foo", iterations, &wg)

	wg.Wait()
}
