package discourse

import (
	"context"
	"net/http"

	"github.com/mrusme/gobbs/models/post"
	"github.com/mrusme/gobbs/system/adapter"
	"go.uber.org/zap"
)

type System struct {
	config map[string]interface{}
	logger *zap.SugaredLogger
}

func (sys *System) GetConfig() map[string]interface{} {
	return sys.config
}

func (sys *System) SetConfig(cfg *map[string]interface{}) {
	sys.config = *cfg
}

func (sys *System) SetLogger(logger *zap.SugaredLogger) {
	sys.logger = logger
}

func (sys *System) Load() error {

	return nil
}

func (sys *System) GetCapabilities() []adapter.Capability {
	var caps []adapter.Capability

	caps = append(caps, adapter.Capability{
		ID:   "posts",
		Name: "Posts",
	})
	caps = append(caps, adapter.Capability{
		ID:   "groups",
		Name: "Groups",
	})
	caps = append(caps, adapter.Capability{
		ID:   "search",
		Name: "Search",
	})

	return caps
}

func (sys *System) ListPosts() ([]post.Post, error) {
	credentials := make(map[string]string)
	for k, v := range (sys.config["credentials"]).(map[string]interface{}) {
		credentials[k] = v.(string)
	}
	c := NewClient(&ClientConfig{
		Endpoint:    sys.config["url"].(string),
		Credentials: credentials,
		HTTPClient:  http.DefaultClient,
		Logger:      sys.logger,
	})

	posts, err := c.Posts.List(context.Background())
	if err != nil {
		return []post.Post{}, err
	}

	var mPosts []post.Post
	for _, p := range posts {
		mPosts = append(mPosts, post.Post{
			ID:      string(p.ID),
			Subject: p.TopicTitle,
		})
	}

	return mPosts, nil
}
