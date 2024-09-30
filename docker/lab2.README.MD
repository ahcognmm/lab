## More container?

Mô tả: Chúng ta có api sinh random số ở folder `api`. Hãy viết Dockerfile cho app này với hướng dẫn chạy ở readme.md

note: Docker image sau khi build ra phải nhẹ nhất có thể.
Hint: Hãy dùng nhiều stage trong Dockerfile.

## Not the first run

Mô tả: `api` có chạy kèm với redis. Hãy chạy cả hai ứng dụng này lên bằng docker sao chúng có thể giao tiếp được với nhau để `api` không bị lỗi. ( ở đây chúng ta sẽ không sử dụng docker-compose ).

Hint: Hãy tạo một network riêng và cho 2 container vào network đó khi chạy.
