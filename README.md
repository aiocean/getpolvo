# get-polvo

Proxy server giữa frontend và backend.

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
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:$CLOUDBUILD_SA --role roles/run.admin
```

4. Tạo build trigger với substitution variables: 

```
_ENV: main
```

`_ENV` này là môi trường đang build, nó sẽ quyết định sự khác biệt về resource.
