#### DevOps Pipeline

- create new repository, e.g. https://github.com/stefanhansatos/migros-showcase

- clone it to work with local IDE

- connect it to Cloud Source Repository

- create directory `functions`

- define types in `functions/types.go`

- define function in `functions/http-frontend.go`

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
    --source=https://source.developers.google.com/projects/hybrid-cloud-22365/repos/github_stefanhansatos_migros-showcase/revisions/master/paths/functions
    --set-env-vars=PUBSUB_TOPIC_TRUNC=translation_input \
    --service-account=smbe-22365@hybrid-cloud-22365.iam.gserviceaccount.com
    

```

- create directory `pipeline`

- configure deployment steps in `pipeline/01_http-frontend.yaml`



