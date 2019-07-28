package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

var grpcServerTemplate = `// DO NOT EDIT. This file is generated by app/cmd/tools/generator/main.go

package main

import (
	"app/api/{{ .Name }}"
	"app/idl/{{ .Name }}/{{ .Name }}{{ .Version }}"
	"app/pkg/server"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gRPC endpoint
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGRPCServer(ctx, &server.GRPCOptions{
			ServeAddr: os.Getenv("GRPC_BIND_ADDR"),
			Register: func(s *grpc.Server) {
				{{ .Name }}{{ .Version }}.{{ .Register }}(s, &{{ .Implementor }}{})
			},
		})
	}()

	<-sigs
}
`
var gatewayServerTemplate = `// DO NOT EDIT. This file is generated by app/cmd/tools/generator/main.go

package main

import (
	"app/idl/{{ .Name }}/{{ .Name }}{{ .Version }}"
	"app/pkg/server"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// HTTP endpoint
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: os.Getenv("HTTP_BIND_ADDR"),
			GRPCAddr:  os.Getenv("GRPC_DIAL_ADDR"),
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  {{ .Name }}{{ .Version }}.{{ .Register }},
		})
	}()

	<-sigs
}
`
var dockerfileTemplate = `# DO NOT EDIT. This file is generated by app/cmd/tools/generator/main.go

# Extend gobuild which has predownloaded dependencies.
FROM grpc-tools/gobuild as build
WORKDIR /build
RUN go build -a -o bin/{{ .Bin }} cmd/{{ .Bin }}/main.go

# Create a single binary image.
FROM scratch as {{ .Bin }}
COPY --from=build /build/bin/{{ .Bin }} /{{ .Bin }}/{{ .Bin }}
CMD ["/{{ .Bin }}/{{ .Bin }}"]
`

type grpcConf struct {
	Name        string `mapstructure:"name"`
	Register    string `mapstructure:"register"`
	Implementor string `mapstructure:"implementor"`
}

type gatewayConf struct {
	Name     string `mapstructure:"name"`
	Register string `mapstructure:"register"`
}

type service struct {
	Version string      `mapstructure:"version"`
	GRPC    grpcConf    `mapstructure:"grpc"`
	Gateway gatewayConf `mapstructure:"gateway"`
}

type services map[string]service

type tplData struct {
	Name        string
	Bin         string
	Version     string
	Register    string
	Implementor string
}

var output = os.Getenv("DEVROOT") + "/app/cmd"

func mustRenderTemplate(str string, data interface{}, output string) {
	tpl := template.Must(template.New("").Parse(str))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)
	if err != nil {
		log.Fatalln("Error rendering template", output, err)
	}

	if err := ioutil.WriteFile(output, buf.Bytes(), 0644); err != nil {
		log.Fatalln("Error writing rendered template", output, err)
	}
}

func renderGRPCServer(serviceName string, conf service) {
	// Generate gRPC server
	data := tplData{
		Name:        serviceName,
		Bin:         serviceName + "-grpc-server",
		Version:     conf.Version,
		Register:    conf.GRPC.Register,
		Implementor: conf.GRPC.Implementor,
	}
	serverName := serviceName + "-grpc-server"
	outputDir := output + "/" + serverName
	os.MkdirAll(outputDir, os.ModePerm)
	mustRenderTemplate(grpcServerTemplate, data, outputDir+"/main.go")
	mustRenderTemplate(dockerfileTemplate, data, outputDir+"/Dockerfile")
}

func renderGatewayServer(serviceName string, conf service) {
	// Generate HTTP gateway server
	data := tplData{
		Name:     serviceName,
		Bin:      serviceName + "-gateway-server",
		Version:  conf.Version,
		Register: conf.Gateway.Register,
	}
	serverName := serviceName + "-gateway-server"
	outputDir := output + "/" + serverName
	os.MkdirAll(outputDir, os.ModePerm)
	mustRenderTemplate(gatewayServerTemplate, data, outputDir+"/main.go")
	mustRenderTemplate(dockerfileTemplate, data, outputDir+"/Dockerfile")
}

func main() {
	// Load up configuration.
	viper.AddConfigPath("./")
	viper.SetConfigName("services")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading conf", err)
	}

	serviceConf := services{}
	if err := viper.Unmarshal(&serviceConf); err != nil {
		log.Fatalln("Error loading conf", err)
	}

	for serviceName, conf := range serviceConf {
		renderGRPCServer(serviceName, conf)
		renderGatewayServer(serviceName, conf)
	}
}
