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
		log.Info().Interface("IntKey", in).Msg("inserted")
		return in, nil
	}

	if err != nil {
		log.Error().Msg(err.Error())
	}

	return nil, err
}

func (r *RPCServer) Delete(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Delete")
	ok, err := store.MemStoreIns().Delete(ctx, in.Key)
	if ok {
		log.Info().Interface("IntKey", in).Msg("deleted")
		return in, nil
	}

	if err != nil {
		log.Error().Msg(err.Error())
	}

	return nil, err
}

func (r *RPCServer) Increment(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Increment")
	ok, err := store.MemStoreIns().Increment(ctx, in.Key, in.Value)
	if ok {
		log.Info().Interface("IntKey", in).Msg("incremented")
		return in, nil
	}

	if err != nil {
		log.Error().Msg(err.Error())
	}
	return nil, err
}

func (r *RPCServer) SafeIncrement(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("SafeIncrement")
	ok, err := store.MemStoreIns().SafeIncrement(ctx, in.Key, in.Value)
	if ok {
		log.Info().
			Interface("IntKey", in).
			Uint64("updated_value", store.MemStoreIns().Get(ctx, in.Key)).Msg("incremented")
		return in, nil
	}

	if err != nil {
		log.Error().Msg(err.Error())
	}
	return nil, err
}

func (r *RPCServer) Decrement(ctx context.Context, in *IntKey) (*IntKey, error) {
	log.Info().Interface("IntKey", in).Msg("Decrement")
	ok, err := store.MemStoreIns().Decrement(ctx, in.Key, in.Value)
	if ok {
		log.Info().Interface("IntKey", in).Msg("decremented")
		return in, nil
	}

	if err != nil {
		log.Error().Msg(err.Error())
	}

	return nil, err
}
