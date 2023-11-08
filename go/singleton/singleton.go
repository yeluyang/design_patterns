package singleton

type (
	Singleton interface {
		Do()
	}

	SingletonGetter interface {
		Get() Singleton
	}
)
