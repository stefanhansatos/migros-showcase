
#### Cloud Build

```bash
gcloud projects add-iam-policy-binding hybrid-cloud-22365 \
  --member="serviceAccount:335804897202@cloudbuild.gserviceaccount.com" \
  --role=roles/storage.objectViewer

```