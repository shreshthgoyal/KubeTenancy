package registry

import (
	"github.com/loft-sh/kiosk/pkg/apis/tenancy"
	"github.com/loft-sh/kiosk/pkg/apiserver/registry/account"
	"github.com/loft-sh/kiosk/pkg/apiserver/registry/space"
)

func init() {
	tenancy.NewAccountRESTFunc = account.NewAccountREST
	tenancy.NewSpaceRESTFunc = space.NewSpaceREST
}
