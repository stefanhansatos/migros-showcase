
#### Cloud Build

Firebase Realtime DB Permissions

```bash
gcloud projects add-iam-policy-binding hybrid-cloud-22365 \
  --member="serviceAccount:335804897202@cloudbuild.gserviceaccount.com" \
  --role=roles/firebasedatabase.viewer
```
---
BigQuery Permissions

```bash
gcloud projects add-iam-policy-binding hybrid-cloud-22365 \
  --member="serviceAccount:335804897202@cloudbuild.gserviceaccount.com" \
  --role=roles/bigquery.admin
```
---

Cloud Storage Permissions
```bash
gcloud projects add-iam-policy-binding hybrid-cloud-22365 \
  --member="serviceAccount:335804897202@cloudbuild.gserviceaccount.com" \
  --role=roles/storage.objectViewer

```
---

Pub/Sub Permissions
```bash
gcloud projects add-iam-policy-binding hybrid-cloud-22365 \
  --member="serviceAccount:335804897202@cloudbuild.gserviceaccount.com" \
  --role=roles/storage.objectViewer

```