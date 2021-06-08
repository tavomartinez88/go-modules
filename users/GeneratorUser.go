package users

type generatorUser struct {
	builder userInterface
}

func CreateGeneratorUser(b userInterface) *generatorUser {
	return &generatorUser{
		builder: b,
	}
}

func (g generatorUser) setBuilder(b userInterface)  {
	g.builder = b
}