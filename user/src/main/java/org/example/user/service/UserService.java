package org.example.user.service;


import build.buf.gen.seminar.user.v1.*;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import jakarta.persistence.PersistenceException;
import lombok.RequiredArgsConstructor;
import org.example.user.entity.User;
import org.example.user.mapper.UserMapper;
import org.example.user.repository.UserRepository;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserService extends UserServiceGrpc.UserServiceImplBase {
	private final UserRepository userRepository;
	private final UserMapper userMapper;

	@Override
	public void postUser(PostUserRequest request, StreamObserver<PostUserResponse> responseObserver) {
		try {
			User user = User.builder()
							.name(request.getName())
							.build();

			user = userRepository.save(user);

			PostUserResponse response = PostUserResponse.newBuilder()
														.setUser(userMapper.mapToUserGrpc(user))
														.build();

			responseObserver.onNext(response);
			responseObserver.onCompleted();
		} catch (DataIntegrityViolationException | PersistenceException e) {
			responseObserver.onError(
					new StatusRuntimeException(Status.ALREADY_EXISTS.withDescription("User name already exists"))
			);
		} catch (Exception e) {
			responseObserver.onError(
					new StatusRuntimeException(Status.INTERNAL.withDescription("Internal Server Error"))
			);
		}
	}

	@Override
	public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
		try {
			User account = userRepository.findById(request.getId())
										 .orElseThrow(() -> new StatusRuntimeException(
												 Status.NOT_FOUND.withDescription("User not found")));

			GetUserResponse response = GetUserResponse.newBuilder()
													  .setUser(userMapper.mapToUserGrpc(account))
													  .build();

			responseObserver.onNext(response);
			responseObserver.onCompleted();
		} catch (StatusRuntimeException e) {
			responseObserver.onError(e);
		} catch (Exception e) {
			responseObserver.onError(Status.INTERNAL.withDescription("Internal Server Error").asRuntimeException());
		}
	}
}
