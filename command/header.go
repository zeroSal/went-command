package command

type Header struct {
	Use       string
	Short     string
	Long      string
	Flags     *Flags
	Arguments []Argument
}
