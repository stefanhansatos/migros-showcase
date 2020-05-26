{
  "name": "cf-http-translation",
  "description": "TranslationHTTP",
  "disabled": true,
  "filename": "pipeline/TranslationHTTP.yaml",
  "substitutions": {
    "_PUBSUB_TOPIC_TRUNC": "translation_input"
  },
  "triggerTemplate": {
    "branchName": "^master$",
    "projectId": "hybrid-cloud-22365",
    "repoName": "github_stefanhansatos_migros-showcase"
  }
}
