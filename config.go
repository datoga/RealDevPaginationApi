package main

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	PageSize int
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) LoadFromEnv() error {
	var err error

	config.Port, err = config.getPortFromEnv()

	if err != nil {
		return err
	}

	config.PageSize, err = config.getPageSizeFromEnv()

	if err != nil {
		return err
	}

	return nil
}

func (config Config) getPortFromEnv() (int, error) {
	sPort := os.Getenv("PORT")

	if sPort == "" {
		return -1, errors.New("Port should not be void")
	}

	port, err := strconv.Atoi(sPort)

	if err != nil {
		return -1, err
	}

	if port < 0 || port > 65535 {
		return port, errors.New("Port out of range")
	}

	return port, nil
}

func (config Config) getPageSizeFromEnv() (int, error) {
	sPageSize := os.Getenv("PAGE_SIZE")

	if sPageSize == "" {
		return -1, errors.New("Page size should not be void")
	}

	pageSize, err := strconv.Atoi(sPageSize)

	if err != nil {
		return -1, err
	}

	if pageSize < 0 {
		return pageSize, errors.New("Page size out of range")
	}

	return pageSize, nil
}
