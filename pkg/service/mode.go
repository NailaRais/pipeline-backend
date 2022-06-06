package service

import (
	"context"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/instill-ai/pipeline-backend/pkg/datamodel"
	"google.golang.org/grpc/codes"

	connectorPB "github.com/instill-ai/protogen-go/vdp/connector/v1alpha"
	pipelinePB "github.com/instill-ai/protogen-go/vdp/pipeline/v1alpha"
)

func (s *service) getModeByConnRscName(srcConnRscName string, dstConnRscName string) (datamodel.PipelineMode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	srcConnResp, err := s.connectorServiceClient.GetSourceConnector(ctx, &connectorPB.GetSourceConnectorRequest{Name: srcConnRscName})
	if err != nil {
		return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_UNSPECIFIED), err
	}

	srcConnDefResp, err := s.connectorServiceClient.GetSourceConnectorDefinition(ctx, &connectorPB.GetSourceConnectorDefinitionRequest{Name: srcConnResp.GetSourceConnector().GetSourceConnectorDefinition()})
	if err != nil {
		return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_UNSPECIFIED), err
	}

	srcConnType := srcConnDefResp.GetSourceConnectorDefinition().GetConnectorDefinition().GetConnectionType()

	dstConnResp, err := s.connectorServiceClient.GetDestinationConnector(ctx, &connectorPB.GetDestinationConnectorRequest{Name: dstConnRscName})
	if err != nil {
		return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_UNSPECIFIED), err
	}

	dstConnDefResp, err := s.connectorServiceClient.GetDestinationConnectorDefinition(ctx,
		&connectorPB.GetDestinationConnectorDefinitionRequest{Name: dstConnResp.GetDestinationConnector().GetDestinationConnectorDefinition()})
	if err != nil {
		return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_UNSPECIFIED), err
	}

	dstConnType := dstConnDefResp.GetDestinationConnectorDefinition().GetConnectorDefinition().GetConnectionType()

	if srcConnType == connectorPB.ConnectionType_CONNECTION_TYPE_DIRECTNESS &&
		dstConnType == connectorPB.ConnectionType_CONNECTION_TYPE_DIRECTNESS {

		// Relying on a hardcoding naming rule "source-*" and "destination-*" for directness connectors
		if strings.Split(srcConnDefResp.GetSourceConnectorDefinition().GetId(), "-")[1] == strings.Split(dstConnDefResp.GetDestinationConnectorDefinition().GetId(), "-")[1] {
			return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_SYNC), nil
		}

		return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_UNSPECIFIED),
			status.Error(codes.InvalidArgument, "Source and destination connector definition id must be the same if they are both directness connection type")
	}

	return datamodel.PipelineMode(pipelinePB.Pipeline_MODE_ASYNC), nil
}
