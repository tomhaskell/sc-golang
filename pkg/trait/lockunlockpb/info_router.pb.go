// Code generated by protoc-gen-router. DO NOT EDIT.

package lockunlockpb

import (
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.LockUnlockInfoServer that allows routing named requests to specific traits.LockUnlockInfoClient
type InfoRouter struct {
	traits.UnimplementedLockUnlockInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.LockUnlockInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithLockUnlockInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithLockUnlockInfoClientFactory(f func(name string) (traits.LockUnlockInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server grpc.ServiceRegistrar) {
	traits.RegisterLockUnlockInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.LockUnlockInfoClient.
func (r *InfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.LockUnlockInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client any) bool {
	_, ok := client.(traits.LockUnlockInfoClient)
	return ok
}

func (r *InfoRouter) AddLockUnlockInfoClient(name string, client traits.LockUnlockInfoClient) traits.LockUnlockInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.LockUnlockInfoClient)
}

func (r *InfoRouter) RemoveLockUnlockInfoClient(name string) traits.LockUnlockInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.LockUnlockInfoClient)
}

func (r *InfoRouter) GetLockUnlockInfoClient(name string) (traits.LockUnlockInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.LockUnlockInfoClient), nil
}
