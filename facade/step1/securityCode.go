package main

import "fmt"

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is Incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
