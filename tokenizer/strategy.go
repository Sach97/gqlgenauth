package tokenizer

//Strategy holds our strategies
type Strategy interface {
	GenerateToken(userID string) (string, error)
	GetUserID(token string) (string, error)
}

//Tokenizer holds the strategy
type Tokenizer struct {
	Strategy Strategy
}

//GenerateToken return our GenerateToken strategy
func (t *Tokenizer) GenerateToken(userID string) (string, error) {
	str, err := t.Strategy.GenerateToken(userID)
	return str, err
}

//GetUserID return our GetUserID strategy
func (t *Tokenizer) GetUserID(token string) (string, error) {
	str, err := t.Strategy.GetUserID(token)
	return str, err
}
