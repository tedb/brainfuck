package main

import "testing"

func TestStartLoopNonNested(t *testing.T) {
	input := []byte("[....]")
	v := next_close_loop_index(input, 0)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}

func TestStartLoopNested(t *testing.T) {
	input := []byte("[.[.].]")
	v := next_close_loop_index(input, 0)
	if v != 6 {
		t.Error("Expected 6, got ", v)
	}
}

func TestStartLoopNested2(t *testing.T) {
	input := []byte("....[..[.[]..].[].].].")
	v := next_close_loop_index(input, 4)
	if v != 18 {
		t.Error("Expected 18, got ", v)
	}
}

func TestStartLoopNested3(t *testing.T) {
	input := []byte("[.[.].]")
	v := next_close_loop_index(input, 0)
	if v != 6 {
		t.Error("Expected 6, got ", v)
	}
}

func TestEndLoopNotNested(t *testing.T) {
	input := []byte("...[.....]..")
	v := prev_open_loop_index(input, 9)
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}

func TestEndLoopNested(t *testing.T) {
	input := []byte("..[.[[].].[]]..")
	v := prev_open_loop_index(input, 12)
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}

func TestEndLoopNested2(t *testing.T) {
        input := []byte("..[.[[].].[]]..")
        v := prev_open_loop_index(input, 11)
        if v != 10 {
                t.Error("Expected 10, got ", v)
        }
}
