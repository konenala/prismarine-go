package chat

// Color constants for Minecraft chat colors
const (
	ColorBlack       = "black"
	ColorDarkBlue    = "dark_blue"
	ColorDarkGreen   = "dark_green"
	ColorDarkAqua    = "dark_aqua"
	ColorDarkRed     = "dark_red"
	ColorDarkPurple  = "dark_purple"
	ColorGold        = "gold"
	ColorGray        = "gray"
	ColorDarkGray    = "dark_gray"
	ColorBlue        = "blue"
	ColorGreen       = "green"
	ColorAqua        = "aqua"
	ColorRed         = "red"
	ColorLightPurple = "light_purple"
	ColorYellow      = "yellow"
	ColorWhite       = "white"
)

// Style represents the styling of a chat component
type Style struct {
	Bold          bool
	Italic        bool
	Underlined    bool
	Strikethrough bool
	Obfuscated    bool
	Color         string
}

// ApplyStyle applies a style to a component
func (c *Component) ApplyStyle(style Style) {
	c.Bold = style.Bold
	c.Italic = style.Italic
	c.Underlined = style.Underlined
	c.Strikethrough = style.Strikethrough
	c.Obfuscated = style.Obfuscated
	if style.Color != "" {
		c.Color = style.Color
	}
}

// WithColor returns a new component with the specified color
func (c Component) WithColor(color string) Component {
	c.Color = color
	return c
}

// WithBold returns a new component with bold styling
func (c Component) WithBold(bold bool) Component {
	c.Bold = bold
	return c
}

// WithItalic returns a new component with italic styling
func (c Component) WithItalic(italic bool) Component {
	c.Italic = italic
	return c
}
