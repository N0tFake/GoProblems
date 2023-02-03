package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"

	"github.com/urfave/cli"
	"github.com/zcalusic/sysinfo"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Get system information"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "info",
			Value: "all",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "cpu",
			Usage: "Get CPU info",
			Flags: flags,

			Action: getInfoCPU,
		},
		{
			Name:  "so",
			Usage: "Get OS info",
			Flags: flags,

			Action: getInfoSO,
		},
		{
			Name:  "ram",
			Usage: "Get RAM info",
			Flags: flags,

			Action: getInfoRAM,
		},
		{
			Name:  "storage",
			Usage: "Get Storage info",
			Flags: flags,

			Action: getInfoStorage,
		},
		{
			Name:  "product",
			Usage: "Get Product info",
			Flags: flags,

			Action: getInfoProduct,
		},
		{
			Name:  "kernel",
			Usage: "Get Kernek info",
			Flags: flags,

			Action: getInfoKernel,
		},
		{
			Name:  "network",
			Usage: "Get Network info",
			Flags: flags,

			Action: getInfoNetwork,
		},
	}

	return app
}

func systemInfo() (jsonString string) {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("Requires superuser privilege")
	}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.Marshal(&si)
	if err != nil {
		log.Fatal(err)
	}

	jsonString = string(data)

	return
}

func getInfoCPU(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("CPU infos")

	for i, item := range result["cpu"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoSO(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("Operacional System infos")

	for i, item := range result["os"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoRAM(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("memory RAM infos")

	for i, item := range result["memory"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoStorage(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("Storage infos")

	for i, item := range result["storage"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoProduct(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("Product infos")

	for i, item := range result["product"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoKernel(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("Kernel infos")

	for i, item := range result["kernel"] {
		fmt.Println(i, ":", item)
	}

}

func getInfoNetwork(c *cli.Context) {

	data := systemInfo()

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fmt.Println("Network infos")

	for i, item := range result["network"] {
		fmt.Println(i, ":", item)
	}

}
