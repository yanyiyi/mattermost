// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindEndpointForURL(t *testing.T) {
	var youtubeProvider *ProviderEndpoint
	for _, provider := range providers {
		if provider.URL == "https://www.youtube.com/oembed" {
			youtubeProvider = provider
		}
	}

	require.NotNil(t, youtubeProvider)

	for _, testCase := range []struct {
		Name     string
		Input    string
		Expected *ProviderEndpoint
	}{
		{
			Name:     "random URL",
			Input:    "https://example.com/some/random.url",
			Expected: nil,
		},
		{
			Name:     "YouTube home page",
			Input:    "https://www.youtube.com",
			Expected: nil,
		},
		{
			Name:     "YouTube video",
			Input:    "https://www.youtube.com/watch?v=szfZfQFUSnU",
			Expected: youtubeProvider,
		},
		{
			Name:     "YouTube video with short link and tracking information",
			Input:    "https://youtu.be/Qq3zukqBFqQ?si=iK_TPT20H30mH90G",
			Expected: youtubeProvider,
		},
		{
			Name:     "YouTube video with playlist",
			Input:    "https://www.youtube.com/watch?v=Qq3zukqBFqQ&list=PL-jqvaPsjQpMqnRgFEw_3fuGQbcVDTpaM",
			Expected: youtubeProvider,
		},
		{
			Name:     "YouTube playlist",
			Input:    "https://www.youtube.com/playlist?list=PL-jqvaPsjQpMqnRgFEw_3fuGQbcVDTpaM",
			Expected: youtubeProvider,
		},
		{
			Name:     "YouTube channel",
			Input:    "https://www.youtube.com/@MattermostHQ",
			Expected: nil,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, testCase.Expected, FindEndpointForURL(testCase.Input))
		})
	}
}