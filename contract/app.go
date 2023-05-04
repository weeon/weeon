package contract

type AppInterface interface {
	Run() error
	Service() string
	Env() string
}
