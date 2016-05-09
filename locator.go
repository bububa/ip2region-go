package ip2region

import (
	"net"
	"os"
	"strings"
)

func NewLocator(dbFile string) (*Locator, error) {
	fd, err := os.Open(dbFile)
	if err != nil {
		return nil, err
	}
	fdInfo, err := fd.Stat()
	if err != nil {
		return nil, err
	}
	loc := &Locator{}
	loc.Data = make([]byte, fdInfo.Size())
	_, err = fd.Read(loc.Data)
	if err != nil {
		return nil, err
	}
	var (
		i   = SUPER_BLOCK_LENGTH
		idx int
	)
	for i < TOTAL_HEADER_LENGTH {
		sip := bytesLittleEndianToUint(loc.Data[i : i+4])
		idxptr := bytesLittleEndianToUint(loc.Data[i+4 : i+8])
		i += 8
		if idxptr == 0 {
			break
		}
		loc.HeaderSip = append(loc.HeaderSip, sip)
		loc.HeaderPtr = append(loc.HeaderPtr, idxptr)
		idx += 1
	}
	loc.HeaderLen = idx
	return loc, nil
}

func (this *Locator) Search(ipstr string) *Location {
	ip := ip2long(ipstr)
	var (
		l    int
		h    = this.HeaderLen - 1
		m    int
		p    int
		sptr uint
		eptr uint
		dptr uint
		sip  uint
		eip  uint
	)
	for l <= h {
		m = (l + h) >> 1
		//perfetc matched, just return it
		if ip == this.HeaderSip[m] {
			if m > 0 {
				sptr = this.HeaderPtr[m-1]
				eptr = this.HeaderPtr[m]
			} else {
				sptr = this.HeaderPtr[m]
				eptr = this.HeaderPtr[m+1]
			}
			break
		}

		//less then the middle value
		if ip < this.HeaderSip[m] {
			if m == 0 {
				sptr = this.HeaderPtr[m]
				eptr = this.HeaderPtr[m+1]
				break
			} else if ip > this.HeaderSip[m-1] {
				sptr = this.HeaderPtr[m-1]
				eptr = this.HeaderPtr[m]
				break
			}
			h = m - 1
		} else {
			if m == this.HeaderLen-1 {
				sptr = this.HeaderPtr[m-1]
				eptr = this.HeaderPtr[m]
				break
			} else if ip < this.HeaderSip[m+1] {
				sptr = this.HeaderPtr[m]
				eptr = this.HeaderPtr[m+1]
				break
			}
			l = m + 1
		}
	}

	if sptr == 0 {
		return nil
	}
	indexBlockLen := eptr - sptr
	buf := this.Data[sptr : sptr+indexBlockLen+INDEX_BLOCK_LENGTH]
	dptr = 0
	l = 0
	h = int(indexBlockLen) / INDEX_BLOCK_LENGTH
	for l <= h {
		m = (l + h) >> 1
		p = m * INDEX_BLOCK_LENGTH
		sip = bytesLittleEndianToUint(buf[p : p+4])
		if ip < sip {
			h = m - 1
		} else {
			eip = bytesLittleEndianToUint(buf[p+4 : p+8])
			if ip > eip {
				l = m + 1
			} else {
				dptr = bytesLittleEndianToUint(buf[p+8 : p+12])
				break
			}
		}
	}
	if dptr == 0 {
		return nil
	}

	dataLen := (dptr >> 24) & 0xFF
	dataptr := dptr & 0x00FFFFFF

	buf = this.Data[dataptr : dataptr+dataLen]

	//city_id = bytesBigEndianToUint(buf)
	dataLen -= 4
	ret := strings.Split(string(buf[4:4+dataLen]), "|")
	return &Location{
		Country:  ret[0],
		Region:   ret[1],
		Province: ret[2],
		City:     ret[3],
		ISP:      ret[4],
	}
}

func bytesBigEndianToUint(b []byte) uint {
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24
}

func bytesLittleEndianToUint(b []byte) uint {
	return uint(b[0]) | uint(b[1])<<8 | uint(b[2])<<16 | uint(b[3])<<24
}

func ip2long(ipstr string) uint {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return bytesBigEndianToUint(ip)
}
