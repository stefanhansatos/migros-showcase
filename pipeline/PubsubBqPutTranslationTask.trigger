{
  "name": "cf-pubsub-bq-put-translation-task",
  "description": "PubsubBqPutTranslationTask",
  "disabled": true,
  "filename": "pipeline/PubsubBqPutTranslationTask.yaml",
  "substitutions": {
    "_BQ_DATASET": "migros_showcase",
    "_BQ_TABLE": "translations_v0_0_1"
  },
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
