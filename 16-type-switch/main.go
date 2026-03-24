package main

import (
	"errors"
	"fmt"
)

func main() {

	println(IsNumberIf(float32(100.1)))
	println(IsNumberIf(true))
	println(IsNumberIf("Hello World"))

	println(IsNumber(float32(100.1)))
	println(IsNumber(true))
	println(IsNumber("Hello World"))

	if sq, err := GetSq(100); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sq:", sq)
	}

	if sq, err := GetSq(true); err != nil { // This is caught by the runtime
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sq:", sq)
	}

	sq := GetSqG(100)
	fmt.Println("Sq:", sq)
	//sq = GetSqG(true) // This is caught by the compiler
	//fmt.Println("Sq:", sq)

}

func IsNumberIf(a any) bool {
	if _, ok := a.(int); !ok {
		if _, ok := a.(uint); !ok {
			if _, ok := a.(int8); !ok {
				if _, ok := a.(uint8); !ok {
					if _, ok := a.(int16); !ok {
						if _, ok := a.(uint16); !ok {
							if _, ok := a.(int32); !ok {
								if _, ok := a.(uint32); !ok {
									if _, ok := a.(int64); !ok {
										if _, ok := a.(uint64); !ok {
											if _, ok := a.(float32); !ok {
												if _, ok := a.(float64); !ok {
													return false
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}

func IsNumber(a any) bool {
	switch a.(type) { // type switch
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func GetSq(a any) (any, error) {
	if IsNumber(a) {
		switch v := a.(type) {
		case int:
			return a.(int) * a.(int), nil //can type assert and get the value
		case uint:
			return v * v, nil // can also use the value which is automatically asserted
		case int8:
			return v * v, nil // can also use the value which is automatically asserted
		case uint8:
			return v * v, nil // can also use the value which is automatically asserted
		case int16:
			return v * v, nil // can also use the value which is automatically asserted
		case uint16:
			return v * v, nil // can also use the value which is automatically asserted
		case int32:
			return v * v, nil // can also use the value which is automatically asserted
		case uint32:
			return v * v, nil // can also use the value which is automatically asserted
		case int64:
			return v * v, nil // can also use the value which is automatically asserted
		case uint64:
			return v * v, nil // can also use the value which is automatically asserted
		case float32:
			return v * v, nil // can also use the value which is automatically asserted
		case float64:
			return v * v, nil // can also use the value which is automatically asserted
			// case uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
			// 	return v * v, nil
		}
	}
	return nil, errors.New("invalid number or not a number type")
}

func GetSqG[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](a T) T {
	return a * a
}
