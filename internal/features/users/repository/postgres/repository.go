package users_postgres_repository

import core_repository_postgres_pool "github.com/rrwwmq/golang-todoapp/internal/core/repository/postgres/pool"

type UsersRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewUsersRepository(pool core_repository_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
