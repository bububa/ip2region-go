package ip2region

const (
	INDEX_BLOCK_LENGTH  = 12
	TOTAL_HEADER_LENGTH = 4096
	SUPER_BLOCK_LENGTH  = 8
)

type Locator struct {
	HeaderSip []uint
	HeaderPtr []uint
	HeaderLen int
	Data      []byte
}

type Location struct {
	Country  string
	Region   string
	Province string
	City     string
	ISP      string
}
