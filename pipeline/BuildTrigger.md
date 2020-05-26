
```bash
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l TranslationHTTP.trigger && \
gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=TranslationHTTP.trigger
  
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l PubsubRealtimeDbInsertTranslationTask.trigger && \
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

