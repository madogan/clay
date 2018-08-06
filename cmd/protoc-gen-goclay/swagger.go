package main

import (
	"github.com/golang/glog"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/utrack/grpc-gateway/protoc-gen-swagger/genswagger"
)

func genSwaggerDef(reg *descriptor.Registry, req *plugin.CodeGeneratorRequest) (map[string][]byte, error) {
	gsw := genswagger.New(reg)
	var targets []*descriptor.File
	for _, target := range req.FileToGenerate {
		f, err := reg.LookupFile(target)
		if err != nil {
			glog.Fatal(err)
		}
		targets = append(targets, f)
	}
	outSwag, err := gsw.Generate(targets)
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]byte, len(outSwag))
	for pos := range outSwag {
		ret[req.FileToGenerate[pos]] = []byte(outSwag[pos].GetContent())
	}
	return ret, nil
}
