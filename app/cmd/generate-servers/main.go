package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	tplPath    = os.Getenv("DEVROOT") + "/app/templates"
	outputPath = os.Getenv("DEVROOT") + "/app/cmd"
	version    = "v1"

	serviceName     string
	generateGateway bool
)

func init() {
	flag.StringVar(&serviceName, "servicename", "", "Name of service to generate servers for")
	flag.BoolVar(&generateGateway, "gateway", true, "Generate additional HTTP gateway")

	flag.Parse()
	if serviceName == "" {
		panic("--servicename must be set")
	}
	serviceName = strings.ToLower(serviceName)
}

type tplData struct {
	Name        string
	Bin         string
	Version     string
	Register    string
	Implementor string
}

func main() {
	genGRPC()

	if generateGateway == true {
		genGateway()
	}
}

func genGRPC() {
	data := tplData{
		Name:        serviceName,
		Bin:         serviceName + "-grpc-server",
		Version:     version,
		Register:    "Register" + strings.Title(serviceName) + "APIServer",
		Implementor: serviceName + ".API",
	}

	grpcTpl := tplPath + "/grpc-server.go.tpl"
	grpcOut := outputPath + "/" + data.Bin + "/main.go"

	mustRenderTemplate(grpcOut, grpcTpl, data)

	dockerfileTpl := tplPath + "/Dockerfile.tpl"
	dockerfileOut := outputPath + "/" + data.Bin + "/Dockerfile"
	mustRenderTemplate(dockerfileOut, dockerfileTpl, data)
}

func genGateway() {
	data := tplData{
		Name:     serviceName,
		Bin:      serviceName + "-gateway-server",
		Version:  version,
		Register: "Register" + strings.Title(serviceName) + "APIHandlerFromEndpoint",
	}

	gatewayTpl := tplPath + "/gateway-server.go.tpl"
	gatewayOut := outputPath + "/" + data.Bin + "/main.go"
	mustRenderTemplate(gatewayOut, gatewayTpl, data)

	dockerfileTpl := tplPath + "/Dockerfile.tpl"
	dockerfileOut := outputPath + "/" + data.Bin + "/Dockerfile"
	mustRenderTemplate(dockerfileOut, dockerfileTpl, data)
}

func mustRenderTemplate(outputPath string, inputPath string, data tplData) {
	dir := filepath.Dir(outputPath)
	os.MkdirAll(dir, os.ModePerm)

	tplName := path.Base(inputPath)
	tpl := template.Must(template.New(tplName).ParseFiles(inputPath))

	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(f, data); err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
}
