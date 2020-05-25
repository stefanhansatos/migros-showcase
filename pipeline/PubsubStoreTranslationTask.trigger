{
  "name": "cf-pubsub-store-translation-task",
  "description": "PubsubStoreTranslationTask",
  "disabled": true,
  "filename": "pipeline/PubsubStoreTranslationTask.yaml",
  "substitutions": {
    "_GS_BUCKET_URL": "hybrid-cloud-22365.appspot.com",
    "_GS_TABLE": "translations_v0_0_1"
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
