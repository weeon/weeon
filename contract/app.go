package contract

type App interface {
	Run() error
	Service() string
	Env() string
}
