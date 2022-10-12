package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type IntegrationTestSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestIntegrationTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	suite.Run(t, &IntegrationTestSuite{rc: rc})
}

func (s *IntegrationTestSuite) TearDown() {
	ctx := testCtx()

	// clean up any lingering integrations
	_ = s.rc.DeleteIntegration(ctx, "s3test")
	_ = s.rc.DeleteIntegration(ctx, "testGCSIntegration")
}

func (s *IntegrationTestSuite) TestGetIntegration() {
	ctx := testCtx()

	const iName = "confluent-cloud"
	integration, err := s.rc.GetIntegration(ctx, iName)
	s.NoError(err)
	s.Assert().Equal(iName, integration.GetName())
}

func (s *IntegrationTestSuite) TestListIntegrations() {
	ctx := testCtx()

	_, err := s.rc.ListIntegrations(ctx)
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestCreateS3Integration() {
	ctx := testCtx()
	const name = "s3test"

	_, err := s.rc.CreateS3Integration(ctx, name,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription("test"))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, name)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestCreateGCSIntegration() {
	saKeyFile := skipUnlessEnvSet(s.T(), "GCP_SA_KEY_JSON")

	ctx := testCtx()
	name := "testGCSIntegration"

	_, err := s.rc.CreateGCSIntegration(ctx, name, saKeyFile,
		option.WithGCSIntegrationDescription("created by rockset integration tests"))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, name)
	s.Require().NoError(err)
}
