
```bash
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/cf-pubsub-store-translation-task.trigger
  
  
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubRealtimeDbInsertTranslationTask.trigger
```

