package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Window struct {
	Min int
	Max int
}

func NewWindowFromCursor(cursor string) (*Window, error) {
	var err error

	parts := strings.Split(cursor, "_")

	if len(parts) != 2 {
		return nil, errors.New("Format not recognised")
	}

	var windowMin, windowMax int64

	windowMin, err = strconv.ParseInt(parts[0], 16, 64)

	if err != nil {
		return nil, err
	}

	windowMax, err = strconv.ParseInt(parts[1], 16, 64)

	if err != nil {
		return nil, err
	}

	return &Window{
		int(windowMin),
		int(windowMax),
	}, nil
}

func NewWindow(min int, max int) (*Window, error) {
	if min < 0 || max < 0 {
		return nil, errors.New("Bounds out of range")
	}

	if min > max {
		return nil, errors.New("Incorrect window")
	}

	return &Window{
		Min: min,
		Max: max,
	}, nil
}

func (window Window) GetCursor() string {
	minHex := fmt.Sprintf("%06x", window.Min)
	maxHex := fmt.Sprintf("%06x", window.Max)

	return minHex + "_" + maxHex
}

func (window Window) NextWindow(pageSize int, max int) *Window {
	if window.Max == max {
		return nil
	}

	var end int

	if window.Max+pageSize >= max {
		end = max
	} else {
		end = window.Max + pageSize
	}

	newW, err := NewWindow(window.Max, end)

	if err != nil {
		panic(err)
	}

	return newW
}

func (window Window) PreviousWindow(pageSize int) *Window {
	if window.Min == 0 {
		return nil
	}

	var ini int

	if window.Min-pageSize < 0 {
		ini = 0
	} else {
		ini = window.Min - pageSize
	}

	newW, err := NewWindow(ini, window.Min)

	if err != nil {
		panic(err)
	}

	return newW
}
