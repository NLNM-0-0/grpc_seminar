package utils

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const GrpcErrorDomain = "product_service"

func NewGrpcErrorWithMetadata(code codes.Code, message string, reason string, metadata map[string]string) error {
	st := status.New(code, message)

	detail := &errdetails.ErrorInfo{
		Reason:   reason,
		Domain:   GrpcErrorDomain,
		Metadata: metadata,
	}

	stWithDetails, err := st.WithDetails(detail)
	if err != nil {
		log.Println("Failed to attach metadata:", err)
		return st.Err()
	}

	return stWithDetails.Err()
}
