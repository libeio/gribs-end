
package grids

type Head struct {
	Title	string
	Styles	[]string
	Scripts []string
}

type Header struct {
	ImgSrc	string
}

type NavList struct {
	Href	string
	Text	string
}

type Nav struct {
	Lists []NavList
}

type Footer struct {
	Text string	
}

