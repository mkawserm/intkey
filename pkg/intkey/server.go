package intkey

import (
	"context"
	"github.com/mkawserm/intkey/pkg/store"
	"github.com/rs/zerolog/log"
)

type RPCServer struct {
	*UnimplementedIntKeyRPCServer
}

func (r *RPCServer) Insert(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Insert")

	ok, err := store.MemStoreIns().Insert(ctx, in.Key, in.Value)
	if ok {
		return in, nil
	}

	return nil, err
}

func (r *RPCServer) Delete(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Delete")
	ok, err := store.MemStoreIns().Delete(ctx, in.Key)
	if ok {
		return in, nil
	}

	return nil, err
}

func (r *RPCServer) Increment(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Increment")
	ok, err := store.MemStoreIns().Increment(ctx, in.Key, in.Value)
	if ok {
		return in, nil
	}

	return nil, err
}

func (r *RPCServer) Decrement(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Decrement")
	ok, err := store.MemStoreIns().Decrement(ctx, in.Key, in.Value)
	if ok {
		return in, nil
	}

	return nil, err
}
