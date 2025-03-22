package org.example.user.config;

import io.grpc.Server;
import io.grpc.netty.NettyServerBuilder;
import jakarta.annotation.PreDestroy;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.example.user.service.UserService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.net.InetSocketAddress;

@Slf4j
@Configuration
@RequiredArgsConstructor
public class GrpcServerConfig {
	private Server server;
	private final UserService userService;

	@Value("${grpc.port}")
	private int port;

	@Bean
	public Server grpcServer() throws IOException {
		server = NettyServerBuilder
				.forPort(port)
				.addService(userService)
				.build()
				.start();

		log.info("🚀 GRPC Server is running on port {}...", port);

		new Thread(() -> {
			try {
				server.awaitTermination();
			} catch (InterruptedException e) {
				Thread.currentThread().interrupt();
			}
		}).start();

		return server;
	}

	@PreDestroy
	public void stopServer() {
		if (server != null) {
			log.info("🛑 Shutting down GRPC Server...");
			server.shutdown();
		}
	}
}
