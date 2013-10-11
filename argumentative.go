package argumentative

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	ErrInvalidVariable      = errors.New("Command's RegisterFlag expects a pointer to a variable.")
	ErrUninitializedProgram = errors.New("You must first call Initialize before registering flags or commands.")
)

type Command struct {
	Name        string
	Description string
	Commands    []*Command
	Flags       []*Flag
}

type Flag struct {
	Short       string
	Long        string
	Description string
	value       reflect.Value
}

func (c *Command) RegisterCommand(name, desc string) *Command {
	n := NewCommand(name, desc)
	c.Commands = append(c.Commands, n)

	return n
}

func (c *Command) RegisterFlag(flag Flag, variable interface{}) *Command {
	value := reflect.ValueOf(variable)

	if kind := value.Kind(); reflect.Ptr != kind {
		panic(ErrInvalidVariable)
	}

	flag.value = value.Elem()
	c.Flags = append(c.Flags, &flag)

	return c
}

func (c *Command) HasSubCommand(name string) (*Command, bool) {
	for i := 0; i < len(c.Commands); i++ {
		if name == c.Commands[i].Name {
			return c.Commands[i], true
		}
	}

	return nil, false
}

func (c *Command) HasFlag(flag string) (*Flag, bool) {
	for i := 0; i < len(c.Flags); i++ {
		if c.Flags[i].Matches(flag) {
			return c.Flags[i], true
		}
	}

	return nil, false
}

func NewCommand(name, desc string) (c *Command) {
	c = new(Command)
	c.Name = name
	c.Description = desc

	return
}

func (f *Flag) Matches(flag string) bool {
	return f.Short == flag || f.Long == flag
}

func (f *Flag) setValue(kind reflect.Kind, value interface{}) {
	switch kind {
	case reflect.Float32:
		f.value.Set(reflect.ValueOf(float32(value.(float64))))
	case reflect.Float64:
		f.value.Set(reflect.ValueOf(float64(value.(float64))))
	case reflect.Int:
		f.value.Set(reflect.ValueOf(int(value.(int64))))
	case reflect.Int8:
		f.value.Set(reflect.ValueOf(int8(value.(int64))))
	case reflect.Int16:
		f.value.Set(reflect.ValueOf(int16(value.(int64))))
	case reflect.Int32:
		f.value.Set(reflect.ValueOf(int32(value.(int64))))
	case reflect.Int64:
		f.value.Set(reflect.ValueOf(int64(value.(int64))))
	case reflect.Uint:
		f.value.Set(reflect.ValueOf(uint(value.(uint64))))
	case reflect.Uint8:
		f.value.Set(reflect.ValueOf(uint8(value.(uint64))))
	case reflect.Uint16:
		f.value.Set(reflect.ValueOf(uint16(value.(uint64))))
	case reflect.Uint32:
		f.value.Set(reflect.ValueOf(uint32(value.(uint64))))
	case reflect.Uint64:
		f.value.Set(reflect.ValueOf(uint64(value.(uint64))))
	}
}

func (f *Flag) Parse(arguments *[]string) {
	index, length := 0, len((*arguments))
	kind := f.value.Kind()

	if reflect.Bool == kind {
		f.value.Set(reflect.ValueOf(true))
		return
	}

	if index >= length {
		panic(errors.New(fmt.Sprintf("Unable to parse argument for `%s`", (*arguments)[index])))
	}

	var value interface{}
	var err error

	switch kind {
	case reflect.Float32:
		value, err = strconv.ParseFloat((*arguments)[index+1], 32)
	case reflect.Float64:
		value, err = strconv.ParseFloat((*arguments)[index+1], 64)
	case reflect.Int:
		value, err = strconv.ParseInt((*arguments)[index+1], 10, 0)
	case reflect.Int8:
		value, err = strconv.ParseInt((*arguments)[index+1], 10, 8)
	case reflect.Int16:
		value, err = strconv.ParseInt((*arguments)[index+1], 10, 16)
	case reflect.Int32:
		value, err = strconv.ParseInt((*arguments)[index+1], 10, 32)
	case reflect.Int64:
		value, err = strconv.ParseInt((*arguments)[index+1], 10, 64)
	case reflect.Uint:
		value, err = strconv.ParseUint((*arguments)[index+1], 10, 0)
	case reflect.Uint8:
		value, err = strconv.ParseUint((*arguments)[index+1], 10, 8)
	case reflect.Uint16:
		value, err = strconv.ParseUint((*arguments)[index+1], 10, 16)
	case reflect.Uint32:
		value, err = strconv.ParseUint((*arguments)[index+1], 10, 32)
	case reflect.Uint64:
		value, err = strconv.ParseUint((*arguments)[index+1], 10, 64)
	}

	*arguments = (*arguments)[index+1:]

	if nil != err {
		panic(err)
	}

	f.setValue(kind, value)
}

var program *Command

func Initialize(name, desc string) *Command {
	program = NewCommand(name, desc)
	return program
}

func RegisterFlag(flag Flag, variable interface{}) *Command {
	if nil == program {
		panic(ErrUninitializedProgram)
	}

	return program.RegisterFlag(flag, variable)
}

func RegisterCommand(name, desc string) *Command {
	if nil == program {
		panic(ErrUninitializedProgram)
	}

	return program.RegisterCommand(name, desc)
}

func Parse(arguments []string) {
	context := program
	for i := 0; i < len(arguments); i++ {
		if c, ok := context.HasSubCommand(arguments[i]); ok {
			context = c
			continue
		} else if f, ok := context.HasFlag(arguments[i]); ok {
			remainder := arguments[i:]
			f.Parse(&remainder)
			continue
		}
	}
}
