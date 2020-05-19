#### DevOps Pipeline

- create new repository, e.g. https://github.com/stefanhansatos/migros-showcase

- clone it to work with local IDE

- connect it to Cloud Source Repository

- create directory `functions`

- define types in `functions/types.go`

- define function in `functions/http-frontend.go`

- define function in `functions/pubsub-subscriber.go`

- commit and push to repository

- prepare Pub/Sub target for client version 0.0.1

```bash
gcloud pubsub topics create translation_input_0.0.1
```

- prepare deployment (from local)
```bash
cd functions

go mod init
go mod vendor
```

- add vendor directory to git temporarily, commit, and push

- test manual deployment (from local)

```bash
cd functions

gcloud functions deploy translation --region europe-west1  --entry-point TranslationHTTP --runtime go111 --trigger-http \
    --source=https://source.developers.google.com/projects/hybrid-cloud-22365/repos/github_stefanhansatos_migros-showcase/revisions/master/paths/functions \
    --set-env-vars=PUBSUB_TOPIC_TRUNC=translation_input \
    --allow-unauthenticated \
    --service-account=smbe-22365@hybrid-cloud-22365.iam.gserviceaccount.com

gcloud functions deploy PubsubTranslationTaskReceiver --quiet --region europe-west1  --runtime go111 --trigger-topic=translation_input_0.0.1 \
    --source=https://source.developers.google.com/projects/hybrid-cloud-22365/repos/github_stefanhansatos_migros-showcase/revisions/master/paths/functions

   
curl -X POST "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/translation" \
  -d '{ "clientVersion": "0.0.1", "clientId": "beab10c6-deee-4843-9757-719566214526", "text": "Today is Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'



```

- create directory `pipeline`

- configure deployment steps in `pipeline/01_http-frontend.yaml`

- create trigger in the console

```bash
gcloud builds submit 
```

TEST 

gsutil versioning set on gs://hybrid-cloud-22365_migros-showcase-devops

cd development
gcloud functions deploy PubsubTranslationTaskReceiverV001D --quiet --region europe-west1  --runtime go111 --trigger-topic=translation_input_0.0.1

rm version-0-0-1-development.zip
zip -r version-0-0-1-development.zip ./*
gsutil cp version-0-0-1-development.zip gs://hybrid-cloud-22365_migros-showcase-devops

gcloud functions deploy PubsubTranslationTaskReceiverV001D --quiet --region europe-west1  --runtime go111 --trigger-topic=translation_input_0.0.1 \
  --update-labels=environment=development,version=0-0-1 \
  --source=gs://hybrid-cloud-22365_migros-showcase-devops/version-0-0-1-development.zip
  
gcloud alpha builds triggers run cf-http-frontend --branch=master 
  



gcloud alpha functions add-iam-policy-binding translation --region=europe-west1 --member=allUsers --role=roles/cloudfunctions.invoker




