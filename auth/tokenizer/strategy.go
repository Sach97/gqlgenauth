package tokenizer

//Strategy holds our strategies
type Strategy interface {
	GenerateString() (string, error)
	GetToken(key string) (string, error)
}

//Tokenizer holds the strategy
type Tokenizer struct {
	Strategy Strategy
}

//GenerateString return our GenerateString strategy
func (t *Tokenizer) GenerateString() (string, error) {
	str, err := t.Strategy.GenerateString()
	return str, err
}

//GetToken return our GetToken strategy
func (t *Tokenizer) GetToken(key string) (string, error) {
	str, err := t.Strategy.GetToken(key)
	return str, err
}
