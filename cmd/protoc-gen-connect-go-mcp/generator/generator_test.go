package generator

import (
	"testing"

	greetv1 "github.com/yoshihiro-shu/connect-go-mcp/cmd/protoc-gen-connect-go-mcp/testdata/greet/gen"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	greetFileDesc := protodesc.ToFileDescriptorProto(greetv1.File_greet_proto)
	compilerVersion := &pluginpb.Version{
		Major:  ptr(int32(0)),
		Minor:  ptr(int32(0)),
		Patch:  ptr(int32(1)),
		Suffix: ptr("test"),
	}
	req := &pluginpb.CodeGeneratorRequest{
		FileToGenerate:        []string{greetFileDesc.GetName()},
		Parameter:             nil,
		ProtoFile:             []*descriptorpb.FileDescriptorProto{greetFileDesc},
		SourceFileDescriptors: []*descriptorpb.FileDescriptorProto{greetFileDesc},
		CompilerVersion:       compilerVersion,
	}

	gen, err := protogen.Options{}.New(req)
	assert.Nil(t, err)

	config := Config{} // デフォルト設定
	err = Generate(gen, config)
	assert.Nil(t, err)
}

func ptr[T any](v T) *T {
	return &v
}
