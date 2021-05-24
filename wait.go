package rockset

import (
	"context"

	"github.com/rs/zerolog"
)

const (
	collectionStatusREADY = "READY"
)

// WaitUntilAliasAvailable waits until the alias is available.
func (rc *RockClient) WaitUntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return rc.Retry(ctx, func() (bool, error) {
		_, _, err := rc.AliasesApi.GetAlias(ctx, workspace, alias).Execute()

		if err == nil {
			return false, nil
		}

		re := NewError(err)
		if re.IsNotFoundError() {
			return true, nil
		}
		if re.Retryable() {
			return true, nil
		}

		return false, err
	})
}

// WaitUntilCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitUntilCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.Retry(ctx, rc.collectionHasState(ctx, workspace, name, collectionStatusREADY))
}

// WaitUntilCollectionGone waits until the a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitUntilCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.Retry(ctx, rc.collectionIsGone(ctx, workspace, name))
}

// WaitUntilCollectionDocuments waits until the collection has at least count new documents
func (rc *RockClient) WaitUntilCollectionDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.Retry(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

func (rc *RockClient) collectionIsGone(ctx context.Context, workspace, name string) RetryFunc {
	return func() (bool, error) {
		_, err := rc.GetCollection(ctx, workspace, name)

		if err == nil {
			// the collection still exist, retry
			return true, nil
		}

		re := NewError(err)
		if re.IsNotFoundError() {
			// the collection is gone
			return false, nil
		}
		if re.Retryable() {
			return true, nil
		}

		return false, err
	}
}

func (rc *RockClient) collectionHasState(ctx context.Context, workspace, name, state string) RetryFunc {
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		log.Debug().Str("status", c.GetStatus()).Str("workspace", workspace).
			Str("collection", name).Msg("collectionHasState()")
		if c.GetStatus() == state {
			return false, nil
		}

		return true, nil
	}
}

type docWaiter struct {
	rc        *RockClient
	prevCount int64
}

func (d *docWaiter) collectionHasNewDocs(ctx context.Context, workspace, name string, count int64) RetryFunc {
	d.prevCount = -1
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := d.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		log.Debug().Str("workspace", workspace).Int64("current", current).
			Int64("previous", d.prevCount).Str("collection", name).
			Int64("count", count).Msg("collectionHasNewDocs()")

		if d.prevCount == -1 {
			d.prevCount = current
		}

		if current-d.prevCount >= count {
			return false, nil
		}

		return true, nil
	}
}
