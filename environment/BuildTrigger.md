
```bash
cd ..

ls -l cf-pubsub-store-translation-task.trigger # wrong filename
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/cf-pubsub-store-translation-task.trigger
  
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l pipeline/PubsubRealtimeDbInsertTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubRealtimeDbInsertTranslationTask.trigger
```

