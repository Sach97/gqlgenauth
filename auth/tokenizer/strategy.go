package tokenizer

type Strategy interface {
	GenerateString() (string, error)
	GetToken(key string) (string, error)
}

type Tokenizer struct {
	Strategy Strategy
}

func (t *Tokenizer) GenerateString() (string, error) {
	str, err := t.Strategy.GenerateString()
	return str, err
}

func (t *Tokenizer) GetToken(key string) (string, error) {
	str, err := t.Strategy.GetToken(key)
	return str, err
}
