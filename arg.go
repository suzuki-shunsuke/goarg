package goarg

type (
	// Parser represents a parser.
	Parser struct {
		handlers []Handler
	}

	// Handler represents a handler.
	Handler struct {
		args []string
		cmd  Command
	}

	// Command represents a command.
	Command func(args ...string) error
)

// NewHandler returns a handler.
func NewHandler(cmd Command, args ...string) Handler {
	return Handler{args: args, cmd: cmd}
}

// NewParser returns a parser.
func NewParser(handlers ...Handler) Parser {
	p := Parser{handlers: handlers}
	return p
}

// Match returns arguments of command and whether arguments include match of the handler.
func (handler *Handler) Match(args ...string) ([]string, bool) {
	s := len(handler.args)
	if len(args) < s {
		return nil, false
	}
	for i, a := range handler.args {
		if args[i] != a {
			return nil, false
		}
	}
	return args[s:], true
}

// Add adds a handler to the parser.
func (parser *Parser) Add(cmd Command, args ...string) *Parser {
	if parser == nil {
		return parser
	}
	parser.handlers = append(parser.handlers, Handler{args: args, cmd: cmd})
	return parser
}

// Parse parses arguments and run the command of matched handler.
func (parser *Parser) Parse(args ...string) (bool, error) {
	if parser == nil {
		return false, nil
	}
	for _, handler := range parser.handlers {
		if arr, f := handler.Match(args...); f {
			return true, handler.cmd(arr...)
		}
	}
	return false, nil
}
