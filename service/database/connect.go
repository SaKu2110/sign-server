package database

import(
	"time"
	"context"
)

func (sql *DB) Connect () {
	// エラー処理用の変数
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 9*time.Second)
	defer cancel()

	go attempt(sql, ctx)
	select {
		case <-ctx.Done():
		sql.MSG = "Database Connection Timeout"
	}
}
