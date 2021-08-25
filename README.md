# Polvo Proxy

Trước đây mình đã xây dựng polvo endpoint nhằm mục đích cho phép cả frontend của phần admin và của client sử dụng chung. nhưng có vài điểm sau làm mình thay đổi:

1. Thư viện sẽ nặng
2. thư viện cần gọi request #1 để lấy được url của endpoint.js, sau đó phải tải cái endpoint này.

Nên mình làm thêm cái proxy này để có thể trực tiếp redirect request của thư viện.

Nhưng Proxy này lấy dữ liệu từ database hay từ polvo service. Mình nghĩ nên lấy từ service, mọi xử lý nên tựu về một chỗ, proxy ra đời là để giải quyết vấn đề giao tiếp, nên nó nên làm đúng việc nó cần làm.
