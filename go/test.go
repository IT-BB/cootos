package main

import "fmt"

type cod interface {
  area()
}

type codec struct {
	codec []int
}

type codecs []codec

func (cc *codecs) area() {
	for _, v := range *cc {
		*cc = append(*cc, v)
	}
	for _, v := range *cc {
		fmt.Println(v.codec)
	}
}

func main() {

	codecss := &codecs{
		codec{codec: []int{1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0}},
		codec{codec: []int{0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1}},
		codec{codec: []int{1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0}},
		codec{codec: []int{0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0}},
		codec{codec: []int{1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0}},
	}

	for i := 0; i < 100; i++ {
		codecss.area()
	}
}
