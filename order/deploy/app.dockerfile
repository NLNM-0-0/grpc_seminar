FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go.mod & go.sum, tải dependencies
COPY order/go.mod order/go.sum ./

RUN go mod download

# Copy source code
COPY . .

# Build với target Linux (nếu chạy trên ARM, đổi GOARCH=arm64)
RUN go build -o /app/bin/order ./order/cmd/order/main.go

# Tạo image nhỏ để chạy app
FROM alpine:3.21

WORKDIR /app

# Copy binary từ build stage
COPY --from=build /app/bin/order .

# Cấp quyền thực thi
RUN chmod +x /app/order

CMD ["/app/order"]
