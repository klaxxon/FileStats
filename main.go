package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var totalBytes uint64
var bitCounts [8]uint64
var byteCounts []uint64
var maxByteValue uint64 // What is the largest value in the byteCounts array?

func bitAnalysis() {
	fmt.Println("\nCounts by bit")
	for a := 0; a < 8; a++ {
		fmt.Printf("%02X = %d  %0.3f%%\n", 1<<a, bitCounts[a], 100.0*float64(bitCounts[a])/float64(totalBytes))
	}
	fmt.Println()
}

func byteAnalysis() {
	fmt.Println("\nCounts by byte value")
	fmt.Print("    ")
	for a := 0; a < 16; a++ {
		fmt.Printf("     %02X", a)
	}
	fmt.Print("\n")
	pos := 0
	for a := 0; a < 16; a++ {
		fmt.Printf("%02X : ", a<<4)
		for b := 0; b < 16; b++ {
			fmt.Printf("%6d ", byteCounts[pos])
			pos++
		}
		fmt.Println()
	}
}

func byteHistogram() {
	vals := " .:-=+*#%@"
	fmt.Println("\nHistogram by byte value")
	fmt.Print("    ")
	for a := 0; a < 16; a++ {
		fmt.Printf(" %02X", a)
	}
	fmt.Print("\n")
	pos := 0
	for a := 0; a < 16; a++ {
		fmt.Printf("%02X :", a<<4)
		for b := 0; b < 16; b++ {
			v := math.Round(9 * float64(byteCounts[pos]) / float64(maxByteValue))
			fmt.Printf("  %c", vals[int(v)])
			pos++
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("File Statistics")
	fn := os.Args[1]

	f, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	maxByteValue = 0
	byteCounts = make([]uint64, 256)
	r := bufio.NewReader(f)
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		totalBytes++
		byteCounts[b]++
		for a := 0; a < 8; a++ {
			if b&1 == 1 {
				bitCounts[a]++
			}
			b >>= 1
		}
	}
	for a := 0; a < 256; a++ {
		if byteCounts[a] > maxByteValue {
			maxByteValue = byteCounts[a]
		}
	}

	// Show the results
	fmt.Printf("File %s, total bytes %d\n", fn, totalBytes)
	byteAnalysis()
	byteHistogram()
	bitAnalysis()
}
