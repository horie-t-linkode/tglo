module tglo_core

go 1.14

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/jason0x43/go-toggl v0.0.0-20210518224841-9af0e37bf5f1
	github.com/kyoh86/go-docbase/v2 v2.0.1
	github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0
	github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/slack-go/slack v0.9.1
)

replace github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0 => ./time_util

replace github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0 => ./template
