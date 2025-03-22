package org.example.user;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class UserApplication implements CommandLineRunner {
	@Value("${spring.datasource.url}")
	private String datasourceUrl;

	public static void main(String[] args) {
		SpringApplication.run(UserApplication.class, args);
	}
	@Override
	public void run(String... args) {
		System.out.println("Datasource URL: " + datasourceUrl);
	}

}
