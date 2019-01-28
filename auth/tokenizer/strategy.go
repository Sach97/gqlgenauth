package tokenizer

type Strategy interface {
	GenerateString() (string, error)
	GetToken(key string) (string, error)
}

type Tokenizer struct {
	Strategy Strategy
}

func (o *Tokenizer) GenerateString() (string, error) {
	str, err := o.Strategy.GenerateString()
	return str, err
}

func (o *Tokenizer) GetToken(key string) (string, error) {
	str, err := o.Strategy.GetToken(key)
	return str, err
}
