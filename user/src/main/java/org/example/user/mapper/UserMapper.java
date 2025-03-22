package org.example.user.mapper;

import org.example.user.entity.User;
import org.springframework.stereotype.Component;

@Component
public class UserMapper {
	public build.buf.gen.seminar.user.v1.User mapToUserGrpc(User user) {
		return build.buf.gen.seminar.user.v1.User.newBuilder()
												 .setId(user.getId())
												 .setName(user.getName())
												 .build();
	}
}
