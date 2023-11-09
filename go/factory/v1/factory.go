package v1

import (
	"bytes"
	"fmt"

	"github.com/yeluyang/design_patterns/go/factory"
)

type factoryV1 struct {
	aFactoryV1
	bFactoryV1
}

type aProductV1 struct {
	content string
}

func (p *aProductV1) A() string {
	return fmt.Sprintf("A.v1:%s", p.content)
}

type bProductV1 struct {
	content string
}

func (p *bProductV1) B() string {
	return fmt.Sprintf("B.v1:%s", p.content)
}

type aFactoryV1 struct{}

func (f *aFactoryV1) A(params ...interface{}) factory.ProductA {
	buf := bytes.Buffer{}
	for _, p := range params {
		_, _ = buf.WriteString(p.(string))
		_, _ = buf.WriteString(",")
	}
	return &aProductV1{content: buf.String()}
}

func (f *aFactoryV1) ADefault() factory.ProductA {
	return &aProductV1{content: "default"}
}

func (f *aFactoryV1) AFromConfig(configPath string) factory.ProductA {
	return &aProductV1{content: fmt.Sprintf("from_config:path=%s", configPath)}
}

type bFactoryV1 struct{}

func (f *bFactoryV1) B(params ...interface{}) factory.ProductB {
	buf := bytes.Buffer{}
	for _, p := range params {
		_, _ = buf.WriteString(p.(string))
		_, _ = buf.WriteString(",")
	}
	return &bProductV1{content: buf.String()}
}

func (f *bFactoryV1) BDefault() factory.ProductB {
	return &bProductV1{content: "default"}
}

func (f *bFactoryV1) BFromConfig(configPath string) factory.ProductB {
	return &bProductV1{content: fmt.Sprintf("from_config:path=%s", configPath)}
}
