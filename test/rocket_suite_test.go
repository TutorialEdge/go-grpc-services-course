// +acceptance

package test

import (
	"context"
	"testing"

	"github.com/TutorialEdge/tutorial-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "ae7b0bf0-fe75-4176-b20a-3ca1371f3226",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "ae7b0bf0-fe75-4176-b20a-3ca1371f3226", resp.Rocket.Id)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
