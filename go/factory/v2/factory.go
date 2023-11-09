package v2

import (
	"bytes"
	"fmt"

	"github.com/yeluyang/design_patterns/go/factory"
)

type factoryV2 struct {
	aFactoryV2
	bFactoryV2
}

type aProductV2 struct {
	content string
}

func (p *aProductV2) A() string {
	return fmt.Sprintf("A.v2:%s", p.content)
}

type bProductV2 struct {
	content string
}

func (p *bProductV2) B() string {
	return fmt.Sprintf("B.v2:%s", p.content)
}

type aFactoryV2 struct{}

func (f *aFactoryV2) A(params ...interface{}) factory.ProductA {
	buf := bytes.Buffer{}
	for _, p := range params {
		_, _ = buf.WriteString(p.(string))
		_, _ = buf.WriteString(",")
	}
	return &aProductV2{content: buf.String()}
}

func (f *aFactoryV2) ADefault() factory.ProductA {
	return &aProductV2{content: "default"}
}

func (f *aFactoryV2) AFromConfig(configPath string) factory.ProductA {
	return &aProductV2{content: fmt.Sprintf("from_config:path=%s", configPath)}
}

type bFactoryV2 struct{}

func (f *bFactoryV2) B(params ...interface{}) factory.ProductB {
	buf := bytes.Buffer{}
	for _, p := range params {
		_, _ = buf.WriteString(p.(string))
		_, _ = buf.WriteString(",")
	}
	return &bProductV2{content: buf.String()}
}

func (f *bFactoryV2) BDefault() factory.ProductB {
	return &bProductV2{content: "default"}
}

func (f *bFactoryV2) BFromConfig(configPath string) factory.ProductB {
	return &bProductV2{content: fmt.Sprintf("from_config:path=%s", configPath)}
}
