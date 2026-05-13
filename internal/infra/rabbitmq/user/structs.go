package userRbMq

type AppExchanges string

const (
	UsersExchanges AppExchanges = "user.exchange"
)

type AppRoutingKeys string

const (
	UserCreatedRK AppRoutingKeys = "user.created"
)

type AppQueues string

const (
	UserQueue1 AppQueues = "user.queue.1"
	UserQueue2 AppQueues = "user.queue.2"
)
