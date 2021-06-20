// +bdd
package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/TutorialEdge/tutorial-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "ae7b0bf0-fe75-4176-b20a-3ca1371f3226",
					Name: "V9",
					Type: "Falcon Heavy",
				},
			},
		)
		fmt.Println(resp)
		assert.NoError(s.T(), err)
	})

	s.T().Run("invalid uuid returns correct error response", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "seriously-not-valid-uuid",
					Name: "V9",
					Type: "Falcon Heavy",
				},
			},
		)
		fmt.Println(err)
		assert.Error(s.T(), err)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
