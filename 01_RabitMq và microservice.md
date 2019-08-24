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

### Ví dụ 
Khi user nhập user info vào giao diện web, ứng dụng web đặt PDF-processing vào message (bao gồm cả thông tin về job). Sau đó message được đặt vào hàng chờ được định nghĩa bởi rabitmq. Kiến trúc cơ bản message queue khá đơn giản: Có 1 client aplications gọi là producers tạo ra các messages và gửi chúng đến messages broker 

Các aplications khác gọi là consumer, kêt nối đến queue và subcribes messages từ broker và process các messages đó. Phần mềm tương tác với broker có thể là producer, consumer, hoặc cả 2. Messages được lưu vào queue cho đến khi consumer vào nhận nó

Luồng

![image](https://user-images.githubusercontent.com/45547213/63631802-4514f800-c656-11e9-8990-62cabdde2399.png)


- Producer gửi message vào rabit mq
- Trong rabitmq có exchange, exchange có nhiệm vụ phân loại các message vào queue phù hợp
- message được đưa vào queue
- Sau đó consumer đến lấy

## RabitMQ và server concept
- Producer: app gửi message
- Consumer: app nhận message
- queue: buffer nhận message
- message: data được gửi đi từ producer đến consumer qua rabitmq
- connection : TCP connection từ app của bạn đến rabitmq
- exchange: nhận message từ producer rồi đẩy vào queue dựa theo nguyên tắc được định nghĩa bởi kiểu exchange 
- binding: là cái link giữa queue và exchange
- routing key: là cái exchange nhìn để biết message này phải gửi vào queue nào hay nói cách khác là cái địa chỉ nhà của 1 lá thư

















































