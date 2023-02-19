package app

type (
	Config struct {
		Server Server `yaml:"Server" validate:"required"`
		Logger Logger `yaml:"Logger" validate:"required"`
	}

	Server struct {
		Host string `yaml:"Host" validate:"required"`
		Port string `yaml:"Port" validate:"required"`
	}

	Logger struct {
		Level *int8 `yaml:"Level" validate:"required"`
	}
)
