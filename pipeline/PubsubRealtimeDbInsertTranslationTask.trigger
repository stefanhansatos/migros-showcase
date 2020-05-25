{
  "name": "cf-pubsub-realtime-db-insert-translation-task",
  "description": "PubsubRealtimeDbInsertTranslationTask",
  "disabled": true,
  "filename": "pipeline/PubsubRealtimeDbInsertTranslationTask.yaml",
  "substitutions": {
    "_RTDB_URL": "migros_showcase",
    "_RTDB_TABLE": "translations_v0_0_1"
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
