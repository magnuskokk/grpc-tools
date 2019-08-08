package main

import (
	"app/pkg/server"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	output      = os.Getenv("DEVROOT") + "/app/cmd"
	serviceName string
)

func init() {
	flag.StringVar(&serviceName, "servicename", "", "Name of service to generate servers for")
	flag.Parse()
	if serviceName == "" {
		panic("--servicename must be set")
	}
	serviceName = strings.ToLower(serviceName)
}

func mustRenderTemplate(str string, data interface{}, outputFile string) {
	tpl := template.Must(template.New("").Parse(str))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)
	if err != nil {
		log.Fatalln("Error rendering template", outputFile, err)
	}

	if err := ioutil.WriteFile(outputFile, buf.Bytes(), 0644); err != nil {
		log.Fatalln("Error writing rendered template", outputFile, err)
	}
}

func generateServer(outputDir, tpl string, data server.TemplateData) {
	os.MkdirAll(outputDir, os.ModePerm)
	mustRenderTemplate(tpl, data, outputDir+"/main.go")
	mustRenderTemplate(server.DockerfileTemplate, data, outputDir+"/Dockerfile")
}

func renderGRPCServer(service server.Service) {
	// Generate gRPC server
	data := server.TemplateData{
		Name:        serviceName,
		Bin:         service.GRPC.Bin,
		Version:     service.Version,
		Register:    service.GRPC.Register,
		Implementor: service.GRPC.Implementor,
	}
	outputDir := output + "/" + serviceName + "-grpc-server"

	generateServer(outputDir, server.GrpcTpl, data)
}

func renderGatewayServer(service server.Service) {
	// Generate HTTP gateway server
	data := server.TemplateData{
		Name:     serviceName,
		Bin:      service.Gateway.Bin,
		Version:  service.Version,
		Register: service.Gateway.Register,
	}
	outputDir := output + "/" + serviceName + "-gateway-server"

	generateServer(outputDir, server.GwTpl, data)
}

func main() {
	service := server.Service{
		Version: "v1",
		GRPC: server.GrpcConf{
			Bin:         serviceName + "-grpc-server",
			Register:    "Register" + strings.Title(serviceName) + "APIServer",
			Implementor: serviceName + ".API",
		},
		Gateway: server.GatewayConf{
			Bin:      serviceName + "-gateway-server",
			Register: "Register" + strings.Title(serviceName) + "APIHandlerFromEndpoint",
		},
	}

	renderGRPCServer(service)
	renderGatewayServer(service)
}
