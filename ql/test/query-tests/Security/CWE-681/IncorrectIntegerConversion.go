package main

import (
	"math"
	"strconv"
)

func main() {

}

type something struct {
}
type config struct {
}
type registry struct {
}

func lookupTarget(conf *config, num int32) (int32, error) {
	return 567, nil
}
func lookupNumberByName(reg *registry, name string) (int32, error) {
	return 567, nil
}
func lab(s string) (*something, error) {
	num, err := strconv.Atoi(s)

	if err != nil {
		number, err := lookupNumberByName(&registry{}, s)
		if err != nil {
			return nil, err
		}
		num = int(number)
	}
	target, err := lookupTarget(&config{}, int32(num)) // NOT OK
	if err != nil {
		return nil, err
	}

	// convert the resolved target number back to a string

	s = strconv.Itoa(int(target))

	return nil, nil
}

func testParseInt() {
	{
		parsed, err := strconv.ParseInt("3456", 10, 8)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // OK
		_ = uint8(parsed)  // OK
		_ = int16(parsed)  // OK
		_ = uint16(parsed) // OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseInt("3456", 10, 16)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // OK
		_ = uint16(parsed) // OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseInt("3456", 10, 32)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseInt("3456", 10, 64)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // NOT OK
		_ = uint32(parsed) // NOT OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // NOT OK
		_ = uint(parsed)   // NOT OK
	}
	{
		parsed, err := strconv.ParseInt("3456", 10, 0)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // NOT OK
		_ = uint32(parsed) // NOT OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
}

func testParseUint() {
	{
		parsed, err := strconv.ParseUint("3456", 10, 8)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // OK
		_ = uint8(parsed)  // OK
		_ = int16(parsed)  // OK
		_ = uint16(parsed) // OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseUint("3456", 10, 16)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // OK
		_ = uint16(parsed) // OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseUint("3456", 10, 32)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // OK
		_ = uint32(parsed) // OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
	{
		parsed, err := strconv.ParseUint("3456", 10, 64)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // NOT OK
		_ = uint32(parsed) // NOT OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // NOT OK
		_ = uint(parsed)   // NOT OK
	}
	{
		parsed, err := strconv.ParseUint("3456", 10, 0)
		if err != nil {
			panic(err)
		}
		_ = int8(parsed)   // NOT OK
		_ = uint8(parsed)  // NOT OK
		_ = int16(parsed)  // NOT OK
		_ = uint16(parsed) // NOT OK
		_ = int32(parsed)  // NOT OK
		_ = uint32(parsed) // NOT OK
		_ = int64(parsed)  // OK
		_ = uint64(parsed) // OK
		_ = int(parsed)    // OK
		_ = uint(parsed)   // OK
	}
}

func testAtoi() {
	parsed, err := strconv.Atoi("3456")
	if err != nil {
		panic(err)
	}
	_ = int8(parsed)   // NOT OK
	_ = uint8(parsed)  // NOT OK
	_ = int16(parsed)  // NOT OK
	_ = uint16(parsed) // NOT OK
	_ = int32(parsed)  // NOT OK
	_ = uint32(parsed) // NOT OK
	_ = int64(parsed)  // OK
	_ = uint64(parsed) // OK
	_ = int(parsed)    // OK
	_ = uint(parsed)   // OK
}

type customInt int16

// these should be caught:
func typeAliases(input string) {
	{
		parsed, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			panic(err)
		}
		// NOTE: byte is uint8
		_ = byte(parsed)      // NOT OK
		_ = customInt(parsed) // NOT OK
	}
}

func testBoundsChecking(input string) {
	{
		parsed, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		if parsed <= math.MaxInt8 && parsed >= math.MinInt8 {
			_ = int8(parsed) // OK
		}
		if parsed < math.MaxInt8 {
			_ = int8(parsed) // OK (because we only check for upper bounds)
			if parsed >= 0 {
				_ = int16(parsed) // OK
			}
		}
		if parsed >= math.MinInt8 {
			_ = int8(parsed) // NOT OK
			if parsed <= 0 {
				_ = int16(parsed) // OK
			}
		}
	}
	{
		parsed, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			panic(err)
		}
		if parsed <= math.MaxInt8 {
			_ = uint8(parsed) // OK
		}
		if parsed < 5 {
			_ = uint16(parsed) // OK
		}
		if err == nil && 1 == 1 && parsed < math.MaxInt8 {
			_ = int8(parsed) // OK
		}
		if parsed > 42 {
			_ = uint16(parsed) // NOT OK
		}
		if parsed < 5 {
			return
		}
		_ = uint8(parsed) // OK
	}
}

func testRightShifted(input string) {
	{
		parsed, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			panic(err)
		}
		_ = byte(parsed) // OK
		_ = byte(parsed >> 8)
		_ = byte(parsed >> 16)
		_ = byte(parsed >> 24)
	}
	{
		parsed, err := strconv.ParseInt(input, 10, 16)
		if err != nil {
			panic(err)
		}
		_ = byte(parsed) // NOT OK
		_ = byte(parsed << 8)
	}
}

func testPathWithMoreThanOneSink(input string) {
	parsed, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic(err)
	}
	v := int16(parsed) // NOT OK
	_ = int8(v)        // OK
}
