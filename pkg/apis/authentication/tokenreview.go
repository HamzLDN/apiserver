/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package authentication

import (
	"context"

	rest "github.com/yubo/apiserver/pkg/client"
	"github.com/yubo/golib/api"
)

// TokenReviewsGetter has a method to return a TokenReviewInterface.
// A group's client should implement this interface.
type TokenReviewsGetter interface {
	TokenReviews() TokenReviewInterface
}

// TokenReviewInterface has methods to work with TokenReview resources.
type TokenReviewInterface interface {
	Create(ctx context.Context, tokenReview *TokenReview, opts api.CreateOptions) (*TokenReview, error)
	TokenReviewExpansion
}

// tokenReviews implements TokenReviewInterface
type tokenReviews struct {
	client rest.Interface
}

// newTokenReviews returns a TokenReviews
func newTokenReviews(c *AuthenticationV1Client) *tokenReviews {
	return &tokenReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a tokenReview and creates it.  Returns the server's representation of the tokenReview, and an error, if there is any.
func (c *tokenReviews) Create(ctx context.Context, tokenReview *TokenReview, opts api.CreateOptions) (result *TokenReview, err error) {
	result = &TokenReview{}
	err = c.client.Post().
		Resource("tokenreviews").
		VersionedParams(&opts, rest.ParameterCodec).
		Body(tokenReview).
		Do(ctx).
		Into(result)
	return
}
