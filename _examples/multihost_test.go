package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"../fluent"
)

func BenchmarkSingleHost(b *testing.B) {
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1"})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()
	tag := "myapp.access"
	var data = map[string]string{
		"foo":  "bar",
		"hoge": "hoge"}
	for i := 0; i < 1000000; i++ {
		data["no"] = strconv.Itoa(i + 1)
		e := logger.Post(tag, data)
		if e != nil {
			log.Println("Error while posting log: ", e)
		}
	}
	log.Println("Success to post logs against single host")
}

func BenchmarkMultipleHosts(b *testing.B) {
	logger, err := fluent.New(fluent.Config{FluentHost: "127.0.0.1:24224,127.0.0.1:24225,127.0.0.1:24226"})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()
	tag := "myapp.access"
	var data = map[string]string{
		"foo":  "bar",
		"hoge": "hoge"}
	for i := 0; i < 1000000; i++ {
		data["no"] = strconv.Itoa(i + 1)
		e := logger.Post(tag, data)
		if e != nil {
			log.Println("Error while posting log: ", e)
		}
	}
	log.Println("Success to post logs against multiple hosts")
}
