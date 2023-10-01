package types

import (
	"fmt"
)

type Png struct {
	Signature []byte
	Ihdr *Ihdr
	Idat *Idat
	Iend *Iend
}

type Ihdr struct {
	Lenght []byte
	ChunkType []byte
	ChunkWidth []byte
	ChunkHeight []byte
	BitDepth byte
	ColorType byte
	Compression byte
	Filter byte
	Interrace byte
	Crc []byte
}

type Idat struct {
	Lenght []byte
	ChunkType []byte
	ChunkData []byte
	Crc []byte
}

type Iend struct {
	Lenght []byte
	ChunkType []byte
	Crc []byte
}


func New(signature, ihdr, idat, iend []byte) (*Png, error) {
	sIhdr := &Ihdr {
		Lenght: ihdr[0:4],
		ChunkType: ihdr[4:8],
		ChunkWidth: ihdr[8:12],
		ChunkHeight: ihdr[12:16],
		BitDepth: ihdr[16],
		ColorType: ihdr[17],
		Compression: ihdr[18],
		Filter: ihdr[19],
		Interrace: ihdr[20],
		Crc: ihdr[21:25],
	}
	sIdat := &Idat {
		Lenght: idat[0:4],
		ChunkType: idat[4:8],
		ChunkData: idat[8:len(idat)-4],
		Crc: idat[len(idat)-4:],
	}
	sIend := &Iend {
		Lenght: iend[0:4],
		ChunkType: iend[4:8],
		Crc: iend[8:12],
	}
	sPng := &Png {
		Signature: signature,
		Ihdr: sIhdr,
		Idat: sIdat,
		Iend: sIend,
	}
	return sPng, nil
}

func (p Png) String() string {
	s := fmt.Sprintf("--Signature---------\n")
	s += fmt.Sprintf("%s\n", bytesTo16string(p.Signature))
	s += p.Ihdr.String()
	s += p.Idat.String()
	s += p.Iend.String()
	return s
}

func (i Ihdr) String() string {
	s := fmt.Sprintf("--Ihdr---------\n")
	s += fmt.Sprintf("Length: %s\n", bytesTo16string(i.Lenght))
	s += fmt.Sprintf("ChunkType: %s\n", bytesTo16string(i.ChunkType))
	s += fmt.Sprintf("ChunkWidth: %s\n", bytesTo16string(i.ChunkWidth))
	s += fmt.Sprintf("ChunkHeight: %s\n", bytesTo16string(i.ChunkHeight))
	s += fmt.Sprintf("BitDepth: %#x\n", i.BitDepth)
	s += fmt.Sprintf("ColorType: %#x\n", i.ColorType)
	s += fmt.Sprintf("BitDepth: %#x\n", i.BitDepth)
	s += fmt.Sprintf("Compression: %#x\n", i.Compression)
	s += fmt.Sprintf("Filter: %#x\n", i.Filter)
	s += fmt.Sprintf("Interrace: %#x\n", i.Interrace)
	s += fmt.Sprintf("Crc: %s\n", bytesTo16string(i.Crc))
	return s
}

func (i Idat) String() string {
	s := fmt.Sprintf("--Idat---------\n")
	s += fmt.Sprintf("Length: %s\n", bytesTo16string(i.Lenght))
	s += fmt.Sprintf("ChunkType: %s\n", bytesTo16string(i.ChunkType))
	s += fmt.Sprintf("ChunkData: [xxxxxxxxx]\n")
	s += fmt.Sprintf("Crc: %s\n", bytesTo16string(i.Crc))
	return s
}

func (i Iend) String() string {
	s := fmt.Sprintf("--Iend---------\n")
	s += fmt.Sprintf("Length: %s\n", bytesTo16string(i.Lenght))
	s += fmt.Sprintf("ChunkType: %s\n", bytesTo16string(i.ChunkType))
	s += fmt.Sprintf("Crc: %s\n", bytesTo16string(i.Crc))
	return s
}


func bytesTo16string(bytes []byte) string {
	s := ""
	for _, b := range bytes {
		s += fmt.Sprintf("%#x,", b)
	}
	return s
}