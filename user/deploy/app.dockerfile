FROM maven:3.9.6-eclipse-temurin-17 AS build
WORKDIR /app

# Copy toàn bộ thư mục user vào container
COPY user /app/user/

# Di chuyển vào thư mục user để build
WORKDIR /app/user
RUN chmod +x mvnw && ./mvnw package -DskipTests

FROM eclipse-temurin:17-jdk
WORKDIR /app

# Copy đúng file JAR
COPY --from=build /app/user/target/*.jar /app/app.jar

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
