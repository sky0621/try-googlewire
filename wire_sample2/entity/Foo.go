package entity

type Foo struct {
	Name string
}

func ProviderFoo(name string) Foo {
	return Foo{
		Name: name,
	}
}
