package utils

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const GrpcErrorDomain = "order_service"

func NewGrpcErrorWithMetadata(code codes.Code, message string, reason string, originalErr error, metadata map[string]string) error {
	st := status.New(code, message)

	if metadata == nil {
		metadata = map[string]string{}
	}

	if originalErr != nil {
		origStatus, ok := status.FromError(originalErr)
		if ok {
			for _, detail := range origStatus.Details() {
				if errInfo, ok := detail.(*errdetails.ErrorInfo); ok {
					for k, v := range errInfo.Metadata {
						metadata[k] = v //Merge metadata with original error
					}
				}
			}

			if _, exists := metadata["original_error"]; !exists {
				metadata["original_error"] = message
			}
		} else {
			metadata["original_error"] = originalErr.Error()
		}
	} else {
		metadata["original_error"] = message
	}

	detail := &errdetails.ErrorInfo{
		Reason:   reason,
		Domain:   GrpcErrorDomain,
		Metadata: metadata,
	}

	stWithDetails, err := st.WithDetails(detail)
	if err != nil {
		return st.Err()
	}

	return stWithDetails.Err()
}
