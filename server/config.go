package main

import (
	"fmt"
	"time"
)

//定义服务
type TConfig struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	Addr     string
	Port     int
}

type Option func(*TConfig)

func Protocol(p string) Option {
	return func(s *TConfig) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *TConfig) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *TConfig) {
		s.MaxConns = maxconns
	}
}
func NewTConfig(addr string, port int, options ...func(*TConfig)) (*TConfig, error) {
	srv := TConfig{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func testDealConfig() {
	s1, _ := NewTConfig("localhost", 1024)
	s2, _ := NewTConfig("localhost", 2048, Protocol("udp"))
	s3, _ := NewTConfig("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))

	fmt.Println(s1, s2, s3)
}
