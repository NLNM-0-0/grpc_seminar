package org.example.user.service;


import build.buf.gen.seminar.user.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import org.example.user.entity.User;
import org.example.user.mapper.UserMapper;
import org.example.user.repository.UserRepository;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserService extends UserServiceGrpc.UserServiceImplBase {
	private final UserRepository userRepository;
	private final UserMapper userMapper;

	@Override
	public void postUser(PostUserRequest request, StreamObserver<PostUserResponse> responseObserver) {
		User user = User.builder()
						.name(request.getName())
						.build();

		user = userRepository.save(user);

		PostUserResponse response = PostUserResponse.newBuilder()
													.setUser(userMapper.mapToUserGrpc(user))
													.build();

		responseObserver.onNext(response);
		responseObserver.onCompleted();
	}

	@Override
	public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
		User account = userRepository.findById(request.getId())
									 .orElseThrow();

		GetUserResponse response = GetUserResponse.newBuilder()
												  .setUser(userMapper.mapToUserGrpc(account))
												  .build();

		responseObserver.onNext(response);
		responseObserver.onCompleted();
	}
}
