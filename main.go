package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var totalBytes uint64
var maxBitLength uint64
var bitOneCount uint64
var bitLengthCounts [32]uint64
var bitCounts [8]uint64
var byteCounts []uint64
var maxByteValue uint64 // What is the largest value in the byteCounts array?

// getTextSpan returns a text representation of percent of ln characters #############-----------
func getTextSpan(perc float64, ln int) string {
	x := int(((perc / 100.0) + 0.005) * float64(ln))
	return strings.Repeat("#", x) + strings.Repeat("-", ln-x)
}

func bitAnalysis() {
	fmt.Printf("\nBit 0 count %d or %0.3f%% and bit 1 count %d or %0.3f%%\n", 8*totalBytes-bitOneCount, 100.0*float64(8*totalBytes-bitOneCount)/float64(totalBytes*8), bitOneCount, 100.0*float64(bitOneCount)/float64(totalBytes*8))
	fmt.Println("\nCounts by bit")
	for a := 0; a < 8; a++ {
		perc := 100.0 * float64(bitCounts[a]) / float64(4*totalBytes)
		fmt.Printf("%02X = %8d  %6.3f%%  %s\n", 1<<a, bitCounts[a], perc, getTextSpan(perc, 50))
	}
	fmt.Println()

	fmt.Println("\nCounts by bit span, example 01100001 has 2 of length 1, 1 of length 2 and 1 of length 4")
	fmt.Println("Span length    Count       %")
	for a := 0; a < 32; a++ {
		if bitLengthCounts[a] == 0 {
			continue
		}
		perc := 100.0 * float64(bitLengthCounts[a]) / float64(totalBytes*8)
		fmt.Printf("     %2d     %8d    %6.3f%%  %s\n", a+1, bitLengthCounts[a], perc, getTextSpan(perc, 50))
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
	var f *os.File
	var err error

	fmt.Println("File Statistics")

	fn := "/dev/stdin"
	if len(os.Args) == 2 {
		fn = os.Args[1]
		f, err = os.Open(fn)
	} else {
		f = os.Stdin
	}
	if err != nil {
		log.Fatal(err)
	}

	lastBit := 0
	bitSpanCount := -1 // First bit change does not count when starting
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
			bitSpanCount++
			if b&128 == 128 {
				bitOneCount++
				bitCounts[7-a]++
				if 0 == lastBit {
					if bitSpanCount > 32 {
						if bitSpanCount > int(maxBitLength) {
							maxBitLength = uint64(bitSpanCount)
						}
						bitSpanCount = 32
					}
					if bitSpanCount > 0 {
						bitLengthCounts[bitSpanCount-1]++
					}
					bitSpanCount = 0
					lastBit = 1
				}
			} else if lastBit == 1 {
				if bitSpanCount > 32 {
					if bitSpanCount > int(maxBitLength) {
						maxBitLength = uint64(bitSpanCount)
					}
					bitSpanCount = 32
				}
				bitLengthCounts[bitSpanCount-1]++
				bitSpanCount = 0
				lastBit = 0
			}
			b <<= 1
		}
	}
	bitSpanCount++ // Count last bit
	if bitSpanCount > 32 {
		if bitSpanCount > int(maxBitLength) {
			maxBitLength = uint64(bitSpanCount)
		}
		bitSpanCount = 32
	}
	bitLengthCounts[bitSpanCount-1]++

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
