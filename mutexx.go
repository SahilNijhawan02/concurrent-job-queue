package main

import "sync"

var jobStore = make(map[string]Job)
var mu sync.Mutex
