package types

type UlatCommand struct 
	Verbose 	bool
}

func NewUlatCommandContext(VerboseCmd bool) *UlatCommand {
	return &UlatCommand{
		Verbose: VerboseCmd,
	}
}
