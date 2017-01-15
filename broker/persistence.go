package hrotti

import (
	"github.com/blackbeans/go-uuid"
	. "github.com/gopheracademy/hrotti/packets"
)

type dirFlag byte

const (
	INBOUND  = 1
	OUTBOUND = 2
)

type Persistence interface {
	Init() error
	Open(string)
	Close(string)
	Add(string, dirFlag, ControlPacket) bool
	Replace(string, dirFlag, ControlPacket) bool
	AddBatch(map[string]*PublishPacket)
	Delete(string, dirFlag, uuid.UUID) bool
	GetAll(string) []ControlPacket
	Exists(string) bool
}
