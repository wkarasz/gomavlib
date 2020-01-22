package gomavlib

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCasesFrames = []struct {
	name  string
	key   *Key
	frame Frame
	raw   []byte
}{
	{
		"v1 frame with nil content",
		nil,
		&FrameV1{
			SequenceId:  0x01,
			SystemId:    0x02,
			ComponentId: 0x03,
			Message: &MessageRaw{
				Id:      4,
				Content: nil,
			},
			Checksum: 0x0807,
		},
		[]byte("\xFE\x00\x01\x02\x03\x04\x07\x08"),
	},
	{
		"v1 frame",
		nil,
		&FrameV1{
			SequenceId:  0x27,
			SystemId:    0x01,
			ComponentId: 0x02,
			Message: &MessageTest5{
				'\x10',
				binary.LittleEndian.Uint32([]byte("\x10\x10\x10\x10")),
			},
			Checksum: 0x66e5,
		},
		[]byte("\xFE\x05\x27\x01\x02\x05\x10\x10\x10\x10\x10\xe5\x66"),
	},
	{
		"v1 frame",
		nil,
		&FrameV1{
			SequenceId:  0x27,
			SystemId:    0x01,
			ComponentId: 0x02,
			Message: &MessageTest8{
				'\x10',
				binary.LittleEndian.Uint32([]byte("\x10\x10\x10\x10")),
			},
			Checksum: 0xc7fa,
		},
		[]byte("\xFE\x05\x27\x01\x02\x08\x10\x10\x10\x10\x10\xfa\xc7"),
	},
	{
		"v2 frame with nil content",
		nil,
		&FrameV2{
			IncompatibilityFlag: 0,
			CompatibilityFlag:   0,
			SequenceId:          3,
			SystemId:            4,
			ComponentId:         5,
			Message: &MessageRaw{
				4,
				nil,
			},
			Checksum: 0x0ab7,
		},
		[]byte("\xFD\x00\x00\x00\x03\x04\x05\x04\x00\x00\xb7\x0a"),
	},
	{
		"v2 frame",
		nil,
		&FrameV2{
			IncompatibilityFlag: 0x00,
			CompatibilityFlag:   0x00,
			SequenceId:          0x8F,
			SystemId:            0x01,
			ComponentId:         0x02,
			Message: &MessageTest6{
				'\x10',
				binary.LittleEndian.Uint32([]byte("\x10\x10\x10\x10")),
			},
			Checksum: 0x0349,
		},
		[]byte("\xFD\x05\x00\x00\x8F\x01\x02\x07\x06\x00\x10\x10\x10\x10\x10\x49\x03"),
	},
	{
		"v2 frame signed",
		NewKey(bytes.Repeat([]byte("\x4F"), 32)),
		&FrameV2{
			IncompatibilityFlag: 0x01,
			CompatibilityFlag:   0x00,
			SequenceId:          0x00,
			SystemId:            0x00,
			ComponentId:         0x00,
			Message: &MessageHeartbeat{
				Type:           1,
				Autopilot:      2,
				BaseMode:       3,
				CustomMode:     4,
				SystemStatus:   5,
				MavlinkVersion: 3,
			},
			Checksum:           0xd1d9,
			SignatureLinkId:    1,
			SignatureTimestamp: 2,
			Signature:          &Signature{0x0e, 0x47, 0x04, 0x0c, 0xef, 0x9b},
		},
		[]byte("\xfd\x09\x01\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x01\x02\x03\x05\x03\xd9\xd1\x01\x02\x00\x00\x00\x00\x00\x0e\x47\x04\x0c\xef\x9b"),
	},
	{
		"v2 frame signed",
		NewKey(bytes.Repeat([]byte("\x4F"), 32)),
		&FrameV2{
			IncompatibilityFlag: 0x01,
			CompatibilityFlag:   0x00,
			SequenceId:          0x00,
			SystemId:            0x00,
			ComponentId:         0x00,
			Message: &MessageOpticalFlow{
				TimeUsec:       1,
				SensorId:       2,
				FlowX:          3,
				FlowY:          4,
				FlowCompMX:     5,
				FlowCompMY:     6,
				Quality:        7,
				GroundDistance: 8,
				FlowRateY:      1,
			},
			Checksum:           0xfb77,
			SignatureLinkId:    3,
			SignatureTimestamp: 4,
			Signature:          &Signature{0xa8, 0x88, 0x9, 0x39, 0xb2, 0x60},
		},
		[]byte("\xfd\x22\x01\x00\x00\x00\x00\x64\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa0\x40\x00\x00\xc0\x40\x00\x00\x00\x41\x03\x00\x04\x00\x02\x07\x00\x00\x00\x00\x00\x00\x80\x3f\x77\xfb\x03\x04\x00\x00\x00\x00\x00\xa8\x88\x09\x39\xb2\x60"),
	},
}

func TestParserDecode(t *testing.T) {
	for _, c := range testCasesFrames {
		t.Run(c.name, func(t *testing.T) {
			parser, err := NewParser(ParserConf{
				Reader:      bytes.NewReader(c.raw),
				Writer:      bytes.NewBuffer(nil),
				Dialect:     testDialect,
				OutSystemId: 1,
				InKey:       c.key,
			})
			require.NoError(t, err)
			frame, err := parser.Read()
			require.NoError(t, err)
			require.Equal(t, c.frame, frame)
		})
	}
}

func TestParserEncode(t *testing.T) {
	for _, c := range testCasesFrames {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			parser, err := NewParser(ParserConf{
				Reader:      bytes.NewBuffer(nil),
				Writer:      buf,
				OutSystemId: 1,
				Dialect:     testDialect,
			})
			require.NoError(t, err)
			err = parser.Write(c.frame, true)
			require.NoError(t, err)
			require.Equal(t, c.raw, buf.Bytes())
		})
	}
}

func TestParserEncodeNilMsg(t *testing.T) {
	parser, err := NewParser(ParserConf{
		Reader:      bytes.NewReader(nil),
		Writer:      bytes.NewBuffer(nil),
		Dialect:     nil,
		OutSystemId: 1,
	})
	require.NoError(t, err)

	frame := &FrameV1{Message: nil}
	err = parser.Write(frame, true)
	require.Error(t, err)
}

// ensure that the Frame is left untouched by the encoding procedure, and
// therefore Write() can be called by multiple routines in parallel
func TestParserFrameIsConst(t *testing.T) {
	parser, err := NewParser(ParserConf{
		Reader:      bytes.NewReader(nil),
		Writer:      bytes.NewBuffer(nil),
		Dialect:     MustDialect(3, []Message{&MessageHeartbeat{}}),
		OutSystemId: 1,
		OutKey:      NewKey(bytes.Repeat([]byte("\x7C"), 32)),
	})
	require.NoError(t, err)

	frame := &FrameV2{
		Message: &MessageHeartbeat{
			Type:           1,
			Autopilot:      2,
			BaseMode:       3,
			CustomMode:     4,
			SystemStatus:   5,
			MavlinkVersion: 3,
		},
	}
	original := frame.Clone()

	err = parser.Write(frame, false)
	require.NoError(t, err)
	require.Equal(t, frame, original)
}
