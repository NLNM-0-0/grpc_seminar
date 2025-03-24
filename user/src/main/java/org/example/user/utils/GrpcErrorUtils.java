package org.example.user.utils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.grpc.Metadata;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;

import java.util.HashMap;
import java.util.Map;

public class GrpcErrorUtils {
	public static final String ERROR_DOMAIN = "user_service";
	private static final ObjectMapper objectMapper = new ObjectMapper();
	private static final Metadata.Key<byte[]> METADATA_KEY =
			Metadata.Key.of("grpc-error-bin", Metadata.BINARY_BYTE_MARSHALLER);

	public static StatusRuntimeException newGrpcError(
			Status status,
			String message,
			String reason,
			Throwable originalError,
			Map<String, String> metadataMap) {

		Metadata metadata = new Metadata();

		ErrorDetails errorDetails = new ErrorDetails(ERROR_DOMAIN, message, reason, metadataMap);
		if (originalError != null) {
			errorDetails.metadata.put(
					"original_error",
					originalError.getMessage()
			);
		} else {
			errorDetails.metadata.put(
					"original_error",
					message
			);
		}

		try {
			metadata.put(METADATA_KEY, objectMapper.writeValueAsBytes(errorDetails));
		} catch (JsonProcessingException e) {
			throw new RuntimeException("Failed to serialize error metadata", e);
		}

		return status.withDescription(message).asRuntimeException(metadata);
	}

	public static class ErrorDetails {
		public String domain;
		public String reason;
		public String message;
		public Map<String, String> metadata;

		public ErrorDetails(String domain, String message, String reason, Map<String, String> metadata) {
			this.domain   = domain;
			this.reason   = reason;
			this.message  = message;
			this.metadata = (metadata != null) ? new HashMap<>(metadata) : new HashMap<>();
		}
	}
}
