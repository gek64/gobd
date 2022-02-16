package main

type Mod struct {
	Module  Module      `json:"Module"`
	Require []Require   `json:"Require"`
	Exclude interface{} `json:"Exclude"`
	Replace []Replace   `json:"Replace"`
	Retract interface{} `json:"Retract"`
}
type Module struct {
	Path string `json:"Path"`
}
type Require struct {
	Path    string `json:"Path"`
	Version string `json:"Version"`
}
type Old struct {
	Path string `json:"Path"`
}
type New struct {
	Path string `json:"Path"`
}
type Replace struct {
	Old Old `json:"Old"`
	New New `json:"New"`
}
