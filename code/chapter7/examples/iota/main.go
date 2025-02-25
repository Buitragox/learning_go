package main

import "fmt"

type MailCategory int

const (
	Uncategorized  MailCategory = iota // 0
	Personal                           // 1
	Spam                               // 2
	Social                             // 3
	Advertisements                     // 4
)

// Usually not recommended
const (
	Field1 = 0        // 0
	Field2 = 1 + iota // 2
	Field3 = 20       // 20
	Field4            // 20
	Field5            // 20
	Field6 = iota     // 5
)

type BitField int

// Usually not recommended
const (
	Bit1 BitField = 1 << iota // 1
	Bit2                      // 2
	Bit3                      // 4
	Bit4                      // 8
)

func main() {
	fmt.Println("MailCategory:", Uncategorized, Personal, Spam, Social, Advertisements)
	fmt.Println("Fields:", Field1, Field2, Field3, Field4, Field5, Field6)
	fmt.Println("Bits:", Bit1, Bit2, Bit3, Bit4)
}
