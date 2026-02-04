package types

import "github.com/emicklei/proto"

const PbGoExtension = ".pb2.go"
const ProtoExtension = ".proto"

type ProtoFile struct {
	Path         string
	RelativePath string

	Parsed *proto.Proto
	Proto2 bool

	Package  *ProtoPackage
	Imports  []*ProtoImport
	Enums    []*EnumDefinition
	Messages []*MessageDefinition

	BaseFolder string // used for search for imports
}

type ProtoImport struct {
	FSPath     string
	TargetFile *ProtoFile

	ProtoDef *proto.Import
	Alias    map[TargetPlatform]string
}

type ProtoPackage struct {
	Name ScopedName

	ProtoDef *proto.Package
}
