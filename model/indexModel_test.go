package model

import (
	"testing"
)

func TestExeCmd(t *testing.T) {
	type InputExeCmd struct {
		command string
		noSplit string
	}

	var ValidCmdArr []InputExeCmd
	var InvalidCmdArr []InputExeCmd
	ValidCmdArr = append(ValidCmdArr, InputExeCmd{"docker", "version"})
	ValidCmdArr = append(ValidCmdArr, InputExeCmd{"docker version", ""})
	ValidCmdArr = append(ValidCmdArr, InputExeCmd{"hostname -s", ""})
	ValidCmdArr = append(ValidCmdArr, InputExeCmd{"hostname", "-s"})

	InvalidCmdArr = append(InvalidCmdArr, InputExeCmd{"", "hostname"})
	InvalidCmdArr = append(InvalidCmdArr, InputExeCmd{"", "docker version"})

	// Test valid data
	for _, elem := range ValidCmdArr{
		_, err := ExeCmd(elem.command, elem.noSplit)
		if err != nil{
			t.Error(err.Error())
		}
	}

	// Test invalid data
	for _, elem := range InvalidCmdArr{
		_, err := ExeCmd(elem.command, elem.noSplit)
		if err == nil{
			t.Error(err.Error())
		}
	}

}
