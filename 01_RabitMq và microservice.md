# Microservices and RabbitMQ

## Microservices là gì
Trong kiến trúc microservice (còn gọi là kiến trúc mô-đun), các mô-đun khác nhau được tách rời khỏi nhau. Các mô-đun này thường cung cấp các chức năng khác nhau, tạo thành một ứng dụng phần mềm hoàn chỉnh, thông qua các hoạt động và giao diện được khớp nối với nhau. Không giống như một ứng dụng nguyên khối, nơi tất cả các khối chức năng có mặt trong một mô-đun, một ứng dụng microservice chia các chức năng khác nhau cho các mô-đun khác nhau, để tăng cường khả năng bảo trì và dễ sử dụng. 

### Lợi ích
- Dễ phát triển và duy trì
- Cô lập được lỗi
- Cải thiện năng suất và tốc độ
- Cải thiện khả năng mở rộng
- Dễ hiểu

## Role của message queue trong microservices
Thông thường trong kiến trúc microservices, luôn tồn tại cross-dependencies, mà đòi hỏi rằng không 1 service đơn lẻ nào có thể  perform chức năng của nó mà không có sự trợ giúp của các services khác. Chính vì vậy, hệ thống của bạn cần có một cơ chế cho phép các dịch vụ giữ liên lạc với nhau mà không bị chặn bởi các phản hồi. message queue đáp ứng mục đích này bằng cách cung cấp cơ chế cho các services đẩy các messages đến hàng đợi không đồng bộ và đảm bảo rằng chúng được gửi đến đúng đích. Để thực hiện message queue, bạn cần một message broker; nghĩ về nó như một người đưa thư, người nhận thư từ người gửi và đưa nó đến đúng đích.

## RabitMq
RabbitMQ là một phần mềm message queue được gọi là message broker hoặc queue manager. Nói đơn giản; đó là một phần mềm nơi queue có thể được xác định và các ứng dụng có thể kết nối với queue và chuyển một thông điệp lên nó. message queue cho phép giao tiếp không đồng bộ, có nghĩa là các ứng dụng khác (điểm cuối) đang tạo và tiêu thụ thư tương tác với hàng đợi, thay vì giao tiếp trực tiếp với nhau.

Một message có thể bao gồm bất kỳ thông tin. Ví dụ, nó có thể chứa thông tin về một quy trình / công việc nên bắt đầu trên một ứng dụng khác (có thể trên một máy chủ khác) hoặc đó cũng có thể chỉ là một tin nhắn văn bản đơn giản

message broker lưu trữ các tin nhắn cho đến khi một ứng dụng nhận kết nối và đưa một tin nhắn ra khỏi hàng đợi. Các ứng dụng nhận sau đó xử lý thích hợp tin nhắn. Một nhà sản xuất có thể thêm tin nhắn vào hàng đợi mà không phải chờ xử lý chúng.
