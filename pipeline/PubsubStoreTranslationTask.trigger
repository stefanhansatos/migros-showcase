{
  "name": "cf-pubsub-store-translation-task",
  "description": "PubsubStoreTranslationTask",
  "disabled": true,
  "filename": "pipeline/PubsubStoreTranslationTask.yaml",
  "substitutions": {
    "_GCP_PROJECT": "hybrid-cloud-22365",
    "_GCP_REGION": "europe-west1",
    "_BQ_DATASET": "migros_showcase",
    "_BQ_TABLE": "translations_v0_0_1"
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
