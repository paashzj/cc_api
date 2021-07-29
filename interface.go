package cc_api

import "unsafe"

type BaseConfig struct {
	Id      string
	Version int
}

type CcClient interface {
	Run()
	RegisterConfig(value interface{}, listener Listener)
}

type Listener interface {
	Add(value unsafe.Pointer)
	Mod(oldValue, value unsafe.Pointer)
	Del(value unsafe.Pointer)
}
