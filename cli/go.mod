module tglo_cli

go 1.14

require (
	github.com/joho/godotenv v1.3.0
	github.com/masaki-linkode/tglo/pkg/tglo_core v0.0.0
	github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0
	github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
)

replace github.com/masaki-linkode/tglo/pkg/tglo_core v0.0.0 => ../pkg/tglo_core

replace github.com/masaki-linkode/tglo/pkg/tglo_core/template v0.0.0 => ../pkg/tglo_core/template

replace github.com/masaki-linkode/tglo/pkg/tglo_core/time_util v0.0.0 => ../pkg/tglo_core/time_util
