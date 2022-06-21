package controller

import (
	context "context"

	"github.com/matthewmueller/hackernews"
)

type Controller struct {
	HN *hackernews.Client
}

// Story struct
// type Story struct {
// 	// Fields...
// }

// Index of stories
// GET
func (c *Controller) Index(ctx context.Context) (stories []*hackernews.Story, err error) {
	return c.HN.FrontPage(ctx)
}

// Show story
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (story *hackernews.Story, err error) {
	return c.HN.Find(ctx, id)
}
