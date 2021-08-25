# Polvo Proxy

Trước đây mình đã xây dựng polvo endpoint nhằm mục đích cho phép cả frontend của phần admin và của client sử dụng chung. nhưng có vài điểm sau làm mình thay đổi:

1. Thư viện sẽ nặng
2. thư viện cần gọi request #1 để lấy được url của endpoint.js, sau đó phải tải cái endpoint này.

Nên mình làm thêm cái proxy này để có thể trực tiếp redirect request của thư viện.

Nhưng Proxy này lấy dữ liệu từ database hay từ polvo service. Mình nghĩ nên lấy từ service, mọi xử lý nên tựu về một chỗ, proxy ra đời là để giải quyết vấn đề giao tiếp, nên nó nên làm đúng việc nó cần làm.

## Cài đặt

> Các command dưới này để được gọi trong Cloud Shell.

1. In Cloud Shell, create the Cloud Storage bucket:

```
PROJECT_ID=$(gcloud config get-value project)
gsutil mb gs://${PROJECT_ID}-tfstate
```

2. Enable Object Versioning to keep the history of your deployments:

```
gsutil versioning set on gs://${PROJECT_ID}-tfstate
```

3. Cấp quyền cho account build

```
CLOUDBUILD_SA="$(gcloud projects describe $PROJECT_ID --format 'value(projectNumber)')@cloudbuild.gserviceaccount.com"
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:$CLOUDBUILD_SA --role roles/editor
```
