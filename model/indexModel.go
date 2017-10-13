package model

import (
	"os/exec"
	"fmt"
	"encoding/json"
	"log"
	"bytes"
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

func (d *DockerInfo) GetGeneralInfo(){
	d.getDockerInfo()
	d.getDockerVersion()
	d.getHostname()
}

func (d *DockerInfo) GetDetailedInfo(){
	d.getTop()
}

func (d *DockerInfo) getTop(){
	cmd := exec.Command("docker", "stats", "-a", "--no-stream", "--format", "{{json .}}")
	output, err := cmd.CombinedOutput()
	if err != nil{
		log.Println(err)
		return
	}
	row := bytes.Split(output, []byte("\n"))

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
	cmd := exec.Command("hostname", )
	output, err := cmd.CombinedOutput()
	if err != nil{
		fmt.Println(err)
		d.Hostname = "N/A"
	}
	d.Hostname = string(output)
}

func (d *DockerInfo) getDockerInfo(){
	cmd := exec.Command("docker", "info", "-f", "{{json .}}")
	output, err := cmd.Output()
	if err != nil{
		fmt.Println(err)
	}

	if err := json.Unmarshal(output, &d); err != nil {
		panic(err)
	}
	d.Memory = d.Memory / 1024
}

func (d *DockerInfo) getDockerVersion(){
	cmd := exec.Command("docker", "version", "-f", "{{json .Server}}")
	output, _ := cmd.CombinedOutput()
	if err := json.Unmarshal(output, &d); err != nil {
		panic(err)
	}

}