module tgl_cli

go 1.14

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/jason0x43/go-toggl v0.0.0-20191006134724-98db1b5443ff
	github.com/joho/godotenv v1.3.0
	github.com/masaki-linkode/tgl/pkg v0.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
)

replace github.com/masaki-linkode/tgl/pkg v0.0.0 => ../pkg
