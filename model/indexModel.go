package model

import (
	"os/exec"
	"encoding/json"
	"log"
	"bytes"
	"strings"
)

type Swarm struct {
	LocalNodeState string  	`json:"LocalNodeState"`
}

type DockerInfoDetailed struct {
	ContainerId 	string	`json:"Container"`
	ContainerName 	string	`json:"Name"`
	BlockIO			string	`json:"BlockIO"`
	CPUPerc			string	`json:"CPUPerc"`
	ID				string	`json:"ID"`
	MemPerc			string	`json:"MemPerc"`
	MemUsage		string	`json:"MemUsage"`
	NetIO			string	`json:"NetIO"`
	PIDs			string	`json:"PIDs"`
}

type DockerInfo struct {
	DockerVersion string	`json:"ServerVersion"`
	ApiVersion string		`json:"ApiVersion"`
	Hostname string
	CPU int					`json:"NCPU"`
	Memory int				`json:"MemTotal"`
	Kernel string			`json:"KernelVersion"`
	OS string				`json:"OperatingSystem"`
	Arch string				`json:"Architecture"`
	Swarm					`json:"Swarm"`

	ListContainer	[]DockerInfoDetailed
}

var GeneralInfo DockerInfo

type Info interface {
	GetGeneralInfo()
}

func (d *DockerInfo) GetGeneralInfo(){
	GeneralInfo.getDockerVersion()
	GeneralInfo.getDockerInfo()
	GeneralInfo.getHostname()
	GeneralInfo.getTop()
}

func (d *DockerInfo) getTop(){
	cmd := exec.Command("docker", "stats", "-a", "--no-stream", "--format", "{{json .}}")
	output, err := cmd.CombinedOutput()
	if err != nil{
		log.Println(err)
		return
	}
	row := bytes.Split(output, []byte("\n"))
	d.ListContainer = nil	// truncate after reload page
	for _, a := range row{
		if string(a) == ""{
			break
		}
		var b DockerInfoDetailed
		if err := json.Unmarshal(a, &b); err != nil {
			panic(err)
		}
		d.ListContainer = append(d.ListContainer, b)
	}

}

func (d *DockerInfo) getHostname(){
	output, err := exeCmd("hostname -s", "")
	if err != nil{
		log.Println(err)
		d.Hostname = "N/A"
		return
	}
	d.Hostname = string(output)
}

func (d *DockerInfo) getDockerInfo(){
	output, err := exeCmd("docker info -f", `{{json .}}`)
	if err != nil{
		log.Println(err)
		return
	}

	if err := json.Unmarshal(output, &d); err != nil {
		panic(err)
	}
	d.Memory = d.Memory / 1024
}

func (d *DockerInfo) getDockerVersion(){
	output, err := exeCmd("docker version -f", `{{json .Server}}`)
	if err != nil{
		log.Println(err)
		return
	}
	if err := json.Unmarshal(output, &d); err != nil {
		panic(err)
	}

}

func exeCmd(cmd, noSplitArgs string) ([]byte, error){
	parts := strings.Split(cmd, " ")
	head := parts[0]
	parts = parts[1:]
	if noSplitArgs != ""{
		parts = append(parts, noSplitArgs)
	}
	out, err := exec.Command(head, parts...).Output()
	return out, err

}