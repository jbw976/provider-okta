package user

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_user", func(r *config.Resource) {
		r.ShortGroup = "user"
	})
}