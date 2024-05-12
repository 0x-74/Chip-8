package main

func main() {
	//Slice of 4096 or bytes or 4KB
	memory := make([]byte,4096)
	//index register
	var I uint16
	//subroutine stack
	var stack [16]uint16
	//variable registers
	var V [8]byte
	//delay timer
	var delay byte
	//sound timer 
	var sound byte
	//Program counter 
	var PC byte
	//display
	var	framebuffer [64][32]bool
}
