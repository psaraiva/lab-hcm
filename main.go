package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	configWatcher()

	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		sleepConfig := viper.GetInt("sleep.duration")
		go func(sleep int) {
			defer wg.Done()
			ProcessDiceRoll(NewDiceDefault(), sleep)
		}(sleepConfig)

		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}

func configWatcher() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault("sleep.duration", 10)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found:", err)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
