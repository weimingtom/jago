package jago

import "fmt"

/*202 (0xCA)*/
func BREAKPOINT(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*254 (0xFE)*/
func IMPDEP1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*255 (0xFF)*/
func IMPDEP2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}