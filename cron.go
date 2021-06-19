package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func Cron() {

	repos := repoParse()

	c := cron.New(cron.WithSeconds())

	for _, repo := range repos {
		go func(repo map[string]interface{}) {
			autoCommit := repo["autoCommit"].(bool)
			if autoCommit {
				sec := repo["at"].(string)
				timeSection := timeParse(sec)
				c.AddFunc(fmt.Sprintf("%s %s %s * * *", timeSection[2], timeSection[1], timeSection[0]), func() {
					Execute(repo)
				})
				time.Sleep(time.Second)
			} else {
				Execute(repo)
			}
		}(repo)
	}

	c.Start()

	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}

func timeParse(t string) []string {
	return strings.Split(t, ":")
}

func repoParse() []map[string]interface{} {
	repos := viper.Get("repos")
	reposSlice, ok := repos.([]interface{})

	if !ok {
		fmt.Println("parse repos config fail")
		return nil
	}

	reposMapSlice := make([]map[string]interface{}, 0)

	for _, rps := range reposSlice {
		r, ok := rps.(map[string]interface{})
		if !ok {
			return nil
		}
		reposMapSlice = append(reposMapSlice, r)
	}

	return reposMapSlice
}
