// +acceptance

package test

import (
	"context"
	"testing"

	"github.com/TutorialEdge/tutorial-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	s.T().Run("validates the uuid in the new rocket is a uuid", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "not-a-valid-uuid",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.Error(s.T(), err)
		st := status.Convert(err)
		assert.Equal(s.T(), codes.InvalidArgument, st.Code())
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
