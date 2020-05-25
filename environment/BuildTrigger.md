
```bash
cd ..
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l pipeline/PubsubRealtimeDbInsertTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubRealtimeDbInsertTranslationTask.trigger
  
  
# PubsubStoreTranslationTask
######################################### 
ls -l PubsubStoreTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubStoreTranslationTask.trigger
  
  
# PubsubBqPutTranslationTask
######################################### 
ls -l PubsubBqPutTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubBqPutTranslationTask.trigger
```

