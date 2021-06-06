package distlock

//go:generate mockgen -package=distlock -self_package=github.com/Panzer1119/claircore/pkg/distlock -destination=./locker_mock.go github.com/Panzer1119/claircore/pkg/distlock Locker
