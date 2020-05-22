#### BigQuery

```bash
bq mk migros_showcase

bq mk migros_showcase.translations_v0_0_1

bq load \
  --autodetect \
  --source_format=NEWLINE_DELIMITED_JSON \
  migros_showcase.translations_v0_0_1 \
  "gs://hybrid-cloud-22365.appspot.com/0.0.1/17d55af7-ceb4-4f4a-bfa0-ddcffb46fcde"
``` 