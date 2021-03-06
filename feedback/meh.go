package feedback

type MehCommand struct {
	Description string `short:"d" long:"description" description:"Write your message" required:"true"`
}

func (m *MehCommand) Execute(args []string) error {
	r := RetroItem{
		Description: m.Description,
		Category:    CategoryMeh,
	}

	return Run(r)
}
