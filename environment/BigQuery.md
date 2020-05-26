#### BigQuery

```bash
bq mk migros_showcase

bq mk migros_showcase.translations_v0_0_1

bq load \
  --autodetect \
  --source_format=NEWLINE_DELIMITED_JSON \
  migros_showcase.translations_v0_0_1 \
  "gs://hybrid-cloud-22365.appspot.com/<version>/<client id>/<task id>"
``` 