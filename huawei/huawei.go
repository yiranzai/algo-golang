package main

import (
	"fmt"
)

//Phone
type Phone struct {
	numbers  map[byte][]byte
	keys     []byte
	isNumber bool
	lastKey  byte
	index    int
}

func NewPhone(keys string) *Phone {
	p := &Phone{
		numbers: map[byte][]byte{
			'0': {' '},
			'1': {',', '.'},
			'2': {'a', 'b', 'c'},
			'3': {'d', 'e', 'f'},
			'4': {'g', 'h', 'i'},
			'5': {'j', 'k', 'l'},
			'6': {'m', 'n', 'o'},
			'7': {'p', 'q', 'r', 's'},
			'8': {'t', 'u', 'v'},
			'9': {'w', 'x', 'y', 'z'},
		},
		isNumber: true,
		keys:     []byte(keys),
	}
	return p
}

func (p *Phone) Next(key byte) {
	switch key {
	case '/':
		if !p.isNumber {
			p.One(p.lastKey)
		}
	case '#':
		if !p.isNumber {
			if p.lastKey != 0 {
				p.One(p.lastKey)
			}
		}
		p.isNumber = !p.isNumber
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		p.Num(key)
		if p.lastKey == 0 {
			break
		} else {
			if key != p.lastKey {
				p.One(p.lastKey)
				break
			}
		}
	}
	p.index++
	p.lastKey = key
}

func (p *Phone) Num(key byte) {
	if !p.isNumber {
		return
	}
	switch key {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		fmt.Printf("%s", string(key))
	}

}

func (p *Phone) One(key byte) {
	if p.isNumber {
		return
	}
	switch key {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		if p.index > 0 {
			fmt.Printf("%s", string(p.numbers[key][(p.index-1)%len(p.numbers[key])]))
		}
	}
	p.index = 0
}

func (p *Phone) Print() {
	var key byte
	for _, key = range p.keys {
		p.Next(key)
	}
	p.One(key)
}

func (p *Phone) refreshNumber(key byte) {
	switch key {
	case '0':
		p.numbers[key] = []byte{' '}
	case '1':
		p.numbers[key] = []byte{',', '.'}
	case '2':
		p.numbers[key] = []byte{'a', 'b', 'c'}
	case '3':
		p.numbers[key] = []byte{'d', 'e', 'f'}
	case '4':
		p.numbers[key] = []byte{'g', 'h', 'i'}
	case '5':
		p.numbers[key] = []byte{'j', 'k', 'l'}
	case '6':
		p.numbers[key] = []byte{'m', 'n', 'o'}
	case '7':
		p.numbers[key] = []byte{'p', 'q', 'r', 's'}
	case '8':
		p.numbers[key] = []byte{'t', 'u', 'v'}
	case '9':
		p.numbers[key] = []byte{'w', 'x', 'y', 'z'}
	}
}

func main() {
	keys := ""
	fmt.Scan(&keys)
	p := NewPhone(keys)
	p.Print()
}

func test(keys string) {
	p := NewPhone(keys)
	p.Print()
}
