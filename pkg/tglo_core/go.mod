module tglo_core

go 1.14

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/jason0x43/go-toggl v0.0.0-20191006134724-98db1b5443ff
	github.com/kyoh86/go-docbase/v2 v2.0.1
	github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0
	github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0
	github.com/slack-go/slack v0.6.3
	github.com/snabb/isoweek v1.0.0
)

replace github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0 => ./time_util
replace github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0 => ./template
