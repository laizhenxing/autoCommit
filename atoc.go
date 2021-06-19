package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

const TimeFormat = "2006-01-02 15:04:05"

func init() {
	viper.SetConfigFile("repo.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

type cmd func(v ...interface{}) error

// type cmdPipeline []cmd

// func (cp *cmdPipeline) add(c cmd) {

// }

func Execute(repo map[string]interface{}) {
	fmt.Println("Auto Commit Start")

	execute(add, repo["path"])
	execute(commit, repo["path"], repo["message"])
	execute(push, repo["path"])

	fmt.Println("Execute Finish!")
}

func execute(fn cmd, v ...interface{}) error {
	return fn(v...)
}

func cd(v ...interface{}) error {
	fmt.Println("v = ", v)
	path, ok := v[0].(string)
	if !ok {
		return fmt.Errorf("%v is not string type", v)
	}
	var out bytes.Buffer
	cmd := exec.Command("cd", path)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println(cmd.Dir)
	fmt.Println(out.String())

	pwd("func=[cd]")
	return nil
}

func pwd(v interface{}) error {
	var out bytes.Buffer
	cmd := exec.Command("pwd")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println(v, out.String())
	return nil
}

func add(v ...interface{}) error {
	gitpath := v[0].(string)

	cmd := exec.Command("git", "add", "-A")
	cmd.Dir = gitpath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func commit(v ...interface{}) error {
	gitpath := v[0].(string)
	msg := v[1].(string)

	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf("%s %s", msg, time.Now().Format(TimeFormat)))
	cmd.Dir = gitpath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}

	// output msg
	fmt.Println("[push]", string(output))
	return nil
}

func push(v ...interface{}) error {
	gitpath := v[0].(string)

	cmd := exec.Command("git", "push")
	cmd.Dir = gitpath

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}

	// output msg
	fmt.Println("[push]", string(output))
	return nil
}
