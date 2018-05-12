package actions

import (
	"github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/route"
)

type AcceptAction struct {
}

func (*AcceptAction) Do(p net.Prefix, pa *route.Path) (bool, *route.Path) {
	return true, pa
}
