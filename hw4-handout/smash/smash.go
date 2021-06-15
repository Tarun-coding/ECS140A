package smash

import (
	"io"
	//"bufio"
	"strings"
	//"fmt"
	"sync"
)

type word string

//var mu sync.Mutex
func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)
	buf := new(strings.Builder)
        io.Copy(buf, r)
	arrayOfStrings:=strings.Fields(buf.String())
	var wg sync.WaitGroup
	var mu sync.Mutex

	for stringCount := 0;stringCount< len(arrayOfStrings);stringCount++{
		wg.Add(1)
		go func(stringCount int){
			smasherValue:=smasher(word(arrayOfStrings[stringCount]))

			mu.Lock()
			m[smasherValue]=m[smasherValue]+1
			mu.Unlock()

			wg.Done()
		}(stringCount)
	}
	wg.Wait()

	return m
}
