package delivery

//go:generate protoc --proto_path=proto --go_out=../../../proto/log --go_opt=paths=source_relative --go-grpc_out=../../../proto/log --go-grpc_opt=paths=source_relative log.proto

type SourcesDelivery struct {
}

func New() *SourcesDelivery {
	return &SourcesDelivery{}
}

func (d *SourcesDelivery) InitializeSource() {

}

func (d *SourcesDelivery) GetSourceStat() {

}
