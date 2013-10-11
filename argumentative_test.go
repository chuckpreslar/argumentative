package argumentative

import (
	"testing"
)

func TestBoolFlag(t *testing.T) {
	var got bool

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t"})

	exepected := bool(true)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestFloat32Flag(t *testing.T) {
	var got float32

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1.01"})

	exepected := float32(1.01)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestFloat64Flag(t *testing.T) {
	var got float64

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1.01"})

	exepected := float64(1.01)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestIntFlag(t *testing.T) {
	var got int

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := int(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestInt8Flag(t *testing.T) {
	var got int8

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := int8(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestInt16Flag(t *testing.T) {
	var got int16

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := int16(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestInt32Flag(t *testing.T) {
	var got int32

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := int32(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestInt64Flag(t *testing.T) {
	var got int64

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := int64(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestUintFlag(t *testing.T) {
	var got uint

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := uint(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestUint8Flag(t *testing.T) {
	var got uint8

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := uint8(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestUint16Flag(t *testing.T) {
	var got uint16

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := uint16(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestUint32Flag(t *testing.T) {
	var got uint32

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := uint32(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}

func TestUint64Flag(t *testing.T) {
	var got uint64

	Initialize("Test", "Test program.").
		RegisterFlag(Flag{Short: "-t", Long: "--test"}, &got)

	Parse([]string{"-t", "1"})

	exepected := uint64(1)

	if exepected != got {
		t.Error("Failed to set flag for type %T, expected `%v`, got `%v`.", got, exepected, got)
	}
}
