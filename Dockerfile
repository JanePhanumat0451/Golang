# ใช้ base image ของ Go
FROM golang:alpine

# ติดตั้ง build-base (เครื่องมือสำหรับ build) และ git
RUN apk add --no-cache build-base git

# สร้างไดเรกทอรี /app สำหรับแอปพลิเคชัน
RUN mkdir /app

# กำหนดไดเรกทอรีทำงานเป็น /app
WORKDIR /app

# คัดลอก go.mod และ go.sum มาเพื่อติดตั้ง dependencies ก่อน
ADD go.mod ./
ADD go.sum ./
RUN go mod download

# คัดลอกโค้ดทั้งหมดมาใน container
ADD . .

# ติดตั้ง CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# ใช้ CompileDaemon เพื่อ build และรันแอป
ENTRYPOINT CompileDaemon --build="go build -o main.go" --command=./main
