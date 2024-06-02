package types

type Instruction struct {
	Type    string   `json:"type"`
	Data    string   `json:"data"`
	Tokens  []Token  `json:"tokens"`
	Command *Command `json:"command"`
}

type Token struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
	Data   string `json:"data"`
}

type Command struct {
	Keyword string   `json:"keyword"`
	Args    []string `json:"args"`
}

func NewInstruction(dataType byte, data byte) *Instruction {

	return &Instruction{
		Type:    string(dataType),
		Data:    string(data),
		Tokens:  make([]Token, 0),
		Command: NewCommand(),
	}
}

func NewToken(dataType byte, length int, line []byte) Token {

	return Token{
		Type:   string(dataType),
		Length: length,
		Data:   string(line),
	}
}

func NewCommand() *Command {

	return &Command{Args: make([]string, 0)}
}

func (t *Token) Validate() error {

	return nil
}

func (c *Command) SetKeyword(keyword []byte) {

	if c.Keyword == "" {
		c.Keyword = string(keyword)
	}
}

func (c *Command) PushArg(arg string) {

	if arg == c.Keyword {
		return
	}

	c.Args = append(c.Args, arg)
}
