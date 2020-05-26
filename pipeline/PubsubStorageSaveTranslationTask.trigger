{
  "name": "cf-pubsub-storage-save-translation-task",
  "description": "PubsubStorageSaveTranslationTask",
  "disabled": true,
  "filename": "pipeline/PubsubStorageSaveTranslationTask.yaml",
  "substitutions": {
    "_GS_BUCKET_URL": "hybrid-cloud-22365.appspot.com",
    "_GS_TABLE": "translations_v0_0_1"
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
