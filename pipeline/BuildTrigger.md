
```bash
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l pipeline/PubsubRealtimeDbInsertTranslationTask.trigger && \
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=PubsubRealtimeDbInsertTranslationTask.trigger
  
  
# PubsubBqPutTranslationTask
######################################### 
ls -l PubsubBqPutTranslationTask.trigger && \
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=PubsubBqPutTranslationTask.trigger
  
  
# PubsubStorageSaveTranslationTask
######################################### 
ls -l PubsubStorageSaveTranslationTask.trigger && \
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=PubsubStorageSaveTranslationTask.trigger
```

