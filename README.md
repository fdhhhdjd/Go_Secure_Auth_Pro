<p align="center"><a href="https://profile-forme.com/" target="_blank"><img src="https://res.cloudinary.com/ecommerce2021/image/upload/v1659065987/avatar/logo_begsn1.png" width="300"></a></p>

<p align="center">
<a href="https://www.linkedin.com/in/tai-nguyen-tien-787545213/"><img src="https://img.icons8.com/color/48/000000/linkedin-circled--v1.png" alt="Linkedin"></a>
<a href="https://profile-forme.surge.sh"><img src="https://img.icons8.com/color/48/000000/internet--v1.png" alt="Profile"></a>
<a href="tel:0798805741"><img src="https://img.icons8.com/color/48/000000/apple-phone.png" alt="Phone"></a>
<a href = "mailto:nguyentientai10@gmail.com"><img src="https://img.icons8.com/fluency/48/000000/send-mass-email.png" alt="License"></a>
</p>

# Project: Coding Go Secure Mastering Auth Pro!

## Note: We will have 2 documents for Vietnamese and English.

<!-- go run gif -->
<div align="center">
  <a href="https://go.dev/"><img src="https://raw.githubusercontent.com/fdhhhdjd/fdhhhdjd/main/assets/go_run.gif"></a>
</div>

---

<!-- go run gif -->

# Structure Folders 📁

```
    .dockerignore
    .env
    .env.example
    .gitignore
    .vscode/
        settings.json
    cmd/
        cli/
        cronjob/
        queue/
        server/
    configs/
        common/
        config.go
        yaml/
    docker-compose.dev.yml
    docker-compose.pro.yml
    docs/
        assets/
        CODE.md
        CODETABLE.md
        GO.md
        postman/
        swagger/
    fsnotify.go
    global/
        global.go
    go.mod
    go.sum
    GUILD.md
    internal/
        controllers/
        messaging/
        middlewares/
        models/
        repo/
        ...
    makefile
    migrations/
    pkg/
    README.md
    response/
    scripts/
    sqlc.yaml
    templates/
    tests/
    third_party/
    tmp/
```

# Tài liệu bằng tiếng Vietnamese 🇻🇳.

- `.dockerignore`: Chứa danh sách các tệp và thư mục mà Docker sẽ bỏ qua khi xây dựng hình ảnh.
- `.env`: Chứa các biến môi trường cho dự án.
- `.env.example`: Một ví dụ về tệp `.env` chứa các biến môi trường cần thiết, dùng để hướng dẫn.
- `.github/`: Chứa các tệp cấu hình cho GitHub, như `FUNDING.yml` để cấu hình tài trợ.
- `.gitignore`: Chứa danh sách các tệp và thư mục mà git sẽ bỏ qua.
- `.vscode/`: Chứa cấu hình cho Visual Studio Code, như `settings.json`.
- `cmd/`: Chứa các điểm vào của ứng dụng như CLI, cronjob, hàng đợi, và máy chủ.
- `configs/`: Chứa các tệp cấu hình cho ứng dụng, bao gồm cả cấu hình chung và cấu hình dưới dạng tệp YAML.
- `docker-compose.dev.yml` và `docker-compose.pro.yml`: Chứa cấu hình Docker Compose cho môi trường phát triển và sản xuất.
- `docs/`: Chứa tài liệu dự án, bao gồm cả mã hóa, bảng mã, Go, Postman và Swagger.
- `fsnotify.go`: Tệp này có thể chứa mã để theo dõi các thay đổi tệp hệ thống.
- `global/`: Chứa các biến toàn cục cho ứng dụng.
- `go.mod` và `go.sum`: Quản lý các phụ thuộc của dự án Go.
- `GUILD.md`: Có thể là hướng dẫn hoặc thông tin về cách tham gia và đóng góp cho dự án.
- `internal/`: Chứa mã nguồn nội bộ của ứng dụng, không dành cho việc tái sử dụng bên ngoài.
- `makefile`: Chứa các lệnh tự động hóa cho việc xây dựng và quản lý dự án.
- `migrations/`: Chứa các tệp di cư cơ sở dữ liệu.
- `pkg/`: Chứa các thư viện và gói có thể tái sử dụng bên ngoài dự án.
- `README.md`: Tệp này chứa thông tin tổng quan và hướng dẫn cho dự án.
- `response/`: Có thể chứa mã để xử lý và trả về các phản hồi HTTP.
- `scripts/`: Chứa các kịch bản hỗ trợ cho việc phát triển và triển khai.
- `sqlc.yaml`: Cấu hình cho sqlc, một công cụ tự động hóa việc viết mã từ SQL.
- `templates/`: Chứa các mẫu cho giao diện người dùng hoặc tệp khác.
- `tests/`: Chứa các bài kiểm tra tự động cho dự án.
- `third_party/`: Chứa mã từ các dự án bên thứ ba.
- `tmp/`: Thư mục tạm thời cho các tệp được tạo ra trong quá trình phát triển.

# Document to language English 🏴󠁧󠁢󠁥󠁮󠁧󠁿

- `.dockerignore`: Contains a list of files and directories that Docker should ignore when building an image.
- `.env`: Contains environment variables for the project.
- `.env.example`: An example `.env` file containing necessary environment variables, meant to guide setup.
- `.github/`: Contains configuration files for GitHub, like `FUNDING.yml` for sponsorship settings.
- `.gitignore`: Contains a list of files and directories that git should ignore.
- `.vscode/`: Contains configurations for Visual Studio Code, such as `settings.json`.
- `cmd/`: Contains the application's entry points like CLI, cronjob, queue, and server.
- `configs/`: Contains configuration files for the application, including common configurations and configurations in YAML format.
- `docker-compose.dev.yml` and `docker-compose.pro.yml`: Contain Docker Compose configurations for development and production environments.
- `docs/`: Contains project documentation, including coding standards, code tables, Go guidelines, Postman collections, and Swagger files.
- `fsnotify.go`: This file may contain code to monitor file system changes.
- `global/`: Contains global variables for the application.
- `go.mod` and `go.sum`: Manage the project's Go dependencies.
- `GUILD.md`: May contain guidelines or information on how to join and contribute to the project.
- `internal/`: Contains the application's internal source code, not intended for external reuse.
- `makefile`: Contains automation commands for building and managing the project.
- `migrations/`: Contains database migration files.
- `pkg/`: Contains libraries and packages that can be reused outside the project.
- `README.md`: This file contains an overview and instructions for the project.
- `response/`: May contain code for handling and returning HTTP responses.
- `scripts/`: Contains support scripts for development and deployment.
- `sqlc.yaml`: Configuration for sqlc, a tool for generating code from SQL.
- `templates/`: Contains templates for user interfaces or other files.
- `tests/`: Contains automated tests for the project.
- `third_party/`: Contains code from third-party projects.
- `tmp/`: A temporary directory for files created during development.

# API Swagger

<img alt="Swagger API Documentation" src="./docs/assets/swagger1.png">
<img alt="Swagger API Documentation" src="./docs/assets/swagger2.png">

---

## License: MIT License

## Team Word: Liên hệ công việc https://profile-forme.com

## Author: Nguyen Tien Tai 🚩.

## Tài Khoản Donate li Cf để có động lực code cho anh em tham khảo 😄.

![giphy](https://3.bp.blogspot.com/-SzGvXn2sTmw/V6k-90GH3ZI/AAAAAAAAIsk/Q678Pil-0kITLPa3fD--JkNdnJVKi_BygCLcB/s1600/cf10-fbc08%2B%25281%2529.gif)

## Mk: NGUYEN TIEN TAI

## STK: 1651002972052

## Chi Nhánh: NGAN HANG TMCP AN BINH (ABBANK).

## SUPPORT CONTACT: [https://profile-forme.com](https://profile-forme.com)

## Thank You <3.
