package aoc_assert

import "log"

func Assert(condition bool, error string) {
	if !condition {
		log.Fatalln(error)
	}
}
