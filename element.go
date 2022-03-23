package gom

import (
	"fmt"
)

type Element struct {
	Tag        string
	isFinite   bool
	noClose    bool
	Attributes []Attribute
	Children   []*Element
}

func H(t string, options ...*Option) *Element {
	el := &Element{Tag: t, Attributes: []Attribute{}, Children: []*Element{}}
	for _, option := range options {
		if option.Name == IsFinite.Name {
			el.isFinite = true
		} else if option.Name == NoClose.Name {
			el.noClose = true
		}
	}

	return el
}

func (el *Element) A(attrs ...Attribute) *Element {
	el.Attributes = attrs
	return el
}

func (el *Element) C(children ...*Element) *Element {
	el.Children = children
	return el
}

func (el Element) Build() (html string) {
	if el.isFinite {
		html = el.Tag
		return
	}

	attrs := ""
	for _, attr := range el.Attributes {
		attrs += " " + attr.Build()
	}

	noClose := ""
	if el.noClose {
		noClose = " /"
	}
	html += fmt.Sprintf("<%s%s%s>", el.Tag, attrs, noClose)

	for _, child := range el.Children {
		html += child.Build()
	}

	if !el.noClose {
		html += fmt.Sprintf("</%s>", el.Tag)
	}
	return html
}
