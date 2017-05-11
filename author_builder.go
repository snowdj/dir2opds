package main

import (
	"github.com/lann/builder"
	"golang.org/x/tools/blog/atom"
)

type authorBuilder builder.Builder

func (a authorBuilder) Name(name string) authorBuilder {
	return builder.Set(a, "Name", name).(authorBuilder)
}

func (a authorBuilder) URI(uri string) authorBuilder {
	return builder.Set(a, "URI", uri).(authorBuilder)
}

func (a authorBuilder) Email(email string) authorBuilder {
	return builder.Set(a, "Email", email).(authorBuilder)
}

func (a authorBuilder) InnerXml(inner string) authorBuilder {
	return builder.Set(a, "InnerXml", inner).(authorBuilder)
}

func (a authorBuilder) Build() atom.Person {
	return builder.GetStruct(a).(atom.Person)
}

var AuthorBuilder = builder.Register(authorBuilder{}, atom.Person{}).(authorBuilder)