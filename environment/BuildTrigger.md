
```bash
cd ..
  
# PubsubRealtimeDbInsertTranslationTask
######################################### 
ls -l pipeline/PubsubRealtimeDbInsertTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubRealtimeDbInsertTranslationTask.trigger
  
  
# PubsubBqPutTranslationTask
######################################### 
ls -l PubsubBqPutTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubBqPutTranslationTask.trigger
  
  
# PubsubStorageSaveTranslationTask
######################################### 
ls -l PubsubStorageSaveTranslationTask.trigger

gcloud alpha builds triggers create cloud-source-repositories \
  --trigger-config=pipeline/PubsubStorageSaveTranslationTask.trigger
```

