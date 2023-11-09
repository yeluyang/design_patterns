package factory

type (
	ProductA interface {
		A() string
	}

	FactoryA interface {
		A(params ...interface{}) ProductA
		ADefault() ProductA
		AFromConfig(configPath string) ProductA
	}

	ProductB interface {
		B() string
	}

	FactoryB interface {
		B(params ...interface{}) ProductB
		BDefault() ProductB
		BFromConfig(configPath string) ProductB
	}

	Factory interface {
		FactoryA
		FactoryB
	}
)
