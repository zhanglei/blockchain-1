package blockchain

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

// Size 序列化的大小.
func (d *Block) Size() (s uint64) {
	{
		l := uint64(len(d.Hash))
		{
			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.PrevHash))
		{
			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++
		}
		s += l
	}
	{
		l := uint64(len(d.Data))
		{
			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++
		}
		s += l
	}
	s += 16
	return
}

// Marshal 序列化.
func (d *Block) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)
	{
		buf[0] = byte(d.Height >> 0)
		buf[1] = byte(d.Height >> 8)
		buf[2] = byte(d.Height >> 16)
		buf[3] = byte(d.Height >> 24)
		buf[4] = byte(d.Height >> 32)
		buf[5] = byte(d.Height >> 40)
		buf[6] = byte(d.Height >> 48)
		buf[7] = byte(d.Height >> 56)

	}
	{
		buf[0+8] = byte(d.Timestamp >> 0)
		buf[1+8] = byte(d.Timestamp >> 8)
		buf[2+8] = byte(d.Timestamp >> 16)
		buf[3+8] = byte(d.Timestamp >> 24)
		buf[4+8] = byte(d.Timestamp >> 32)
		buf[5+8] = byte(d.Timestamp >> 40)
		buf[6+8] = byte(d.Timestamp >> 48)
		buf[7+8] = byte(d.Timestamp >> 56)

	}
	{
		l := uint64(len(d.Hash))
		{
			t := uint64(l)
			for t >= 0x80 {
				buf[i+16] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+16] = byte(t)
			i++
		}
		copy(buf[i+16:], d.Hash)
		i += l
	}
	{
		l := uint64(len(d.PrevHash))
		{
			t := uint64(l)
			for t >= 0x80 {
				buf[i+16] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+16] = byte(t)
			i++
		}
		copy(buf[i+16:], d.PrevHash)
		i += l
	}
	{
		l := uint64(len(d.Data))
		{
			t := uint64(l)
			for t >= 0x80 {
				buf[i+16] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+16] = byte(t)
			i++
		}
		copy(buf[i+16:], d.Data)
		i += l
	}
	return buf[:i+16], nil
}

// Unmarshal 反序列化.
func (d *Block) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	{
		d.Height = 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
	}
	{
		d.Timestamp = 0 | (int64(buf[i+0+8]) << 0) | (int64(buf[i+1+8]) << 8) | (int64(buf[i+2+8]) << 16) | (int64(buf[i+3+8]) << 24) | (int64(buf[i+4+8]) << 32) | (int64(buf[i+5+8]) << 40) | (int64(buf[i+6+8]) << 48) | (int64(buf[i+7+8]) << 56)
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i+16] & 0x7F)
			for buf[i+16]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+16]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		d.Hash = string(buf[i+16 : i+16+l])
		i += l
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i+16] & 0x7F)
			for buf[i+16]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+16]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		d.PrevHash = string(buf[i+16 : i+16+l])
		i += l
	}
	{
		l := uint64(0)
		{
			bs := uint8(7)
			t := uint64(buf[i+16] & 0x7F)
			for buf[i+16]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+16]&0x7F) << bs
				bs += 7
			}
			i++
			l = t
		}
		if uint64(cap(d.Data)) >= l {
			d.Data = d.Data[:l]
		} else {
			d.Data = make([]byte, l)
		}
		copy(d.Data, buf[i+16:])
		i += l
	}
	return i + 16, nil
}
