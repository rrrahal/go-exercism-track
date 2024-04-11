package airportrobot

type Greeter interface {
	LanguageName() string
	Greeter(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greeter(name)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greeter(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greeter(name string) string {
	return "Olá " + name + "!"
}
