package gremlin

type ProtoReader interface {
	Unmarshal(data []byte) error
	SourceBytes() []byte
}

type ProtoWriter interface {
	Marshal() []byte
}
