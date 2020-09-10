package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/dto"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
)

func NewClaimResponse(claimDto *dto.ClaimDTO) *ClaimResponse {
	claimResponse := new(ClaimResponse)
	return claimResponse
}

type ClaimResponse struct {
	*auth.ClaimResponse
}
