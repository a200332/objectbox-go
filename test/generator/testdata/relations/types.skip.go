package object

type WithGroup struct {
	Group  *Group `link`
	Groups []*Group
}
