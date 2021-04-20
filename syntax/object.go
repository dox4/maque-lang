package syntax

import (
	"fmt"
	"os"
)

type value struct {
	value interface{}
}

func (v *value) GetValue() interface{} {
	return v.value
}

type NumberValue struct {
	value
}

type BoolValue struct {
	value
}

type StringValue struct {
	value
}

type NullValue struct {
	value
}

const (
	TypeNumber = iota
	TypeBool
	TypeNull
	TypeString
)

type Expression interface {
	Eval() interface{}
	Type() int32
}

// func (v *value) Eval() interface{} {
// 	return v.GetValue()
// }

func (v *value) Type() interface{} {
	panic("value has no type.")
}

func (v *NumberValue) Type() int32 {
	return TypeNumber
}

func (v *NumberValue) Eval() interface{} {
	return v.GetValue()
}

func (v *BoolValue) Type() int32 {
	return TypeBool
}

func (v *BoolValue) Eval() interface{} {
	return v.GetValue()
}

func (v *StringValue) Type() int32 {
	return TypeString
}

func (v *StringValue) Eval() interface{} {
	return v.GetValue()
}

func (v *NullValue) Eval() interface{} {
	return nil
}

func (v *NullValue) Type() int32 {
	return TypeNull
}

type PrimaryExpression = value

func Value(v interface{}) interface{} {
	if v == nil {
		return &NullValue{
			value{nil},
		}
	}
	switch t := v.(type) {
	case float64:
		return &NumberValue{
			value{v.(float64)},
		}
	case bool:
		return &BoolValue{
			value{v.(bool)},
		}
	case string:
		return &StringValue{
			value{v.(string)},
		}
	default:
		_ = fmt.Errorf("unexpected type: %v", t)
		os.Exit(1)
	}
	return nil
}

type NameExpr struct {
	Identifier string
	Env        *Context
}

const (
	NAME = iota + 255
)

type Context struct {
	ctx map[string]interface{}
}

func NewNameExpr(id string, ctx *Context) *NameExpr {
	return &NameExpr{
		Identifier: id,
		Env:        ctx,
	}
}

func (n *NameExpr) GetType() int {
	return NAME
}

func (n *NameExpr) Eval() interface{} {
	val, ok := n.Env.ctx[n.Identifier]
	if ok {
		return val
	}
	panic("unresolved name: " + n.Identifier)
}

