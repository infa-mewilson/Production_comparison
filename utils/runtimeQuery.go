package utils

import "fmt"

func QueryRuntime(startDate string, EndDate string, apis string, filepath string, env string) string {
	var query string
	query = `{
  "aggs": {
    "0": {
      "percentiles": {
        "field": "temp.duration",
        "percents": [
          95,90
        ]
      }
    },"1": {
      "avg": {
        "field": "temp.duration"
      }
    }
  },
  "size": 0,
  "fields": [
    {
      "field": "@timestamp",
      "format": "date_time"
    },
    {
      "field": "dissect.request_date",
      "format": "date_time"
    },
    {
      "field": "dissect_apim_audit.date",
      "format": "date_time"
    },
    {
      "field": "event.created",
      "format": "date_time"
    },
    {
      "field": "event.end",
      "format": "date_time"
    },
    {
      "field": "event.ingested",
      "format": "date_time"
    },
    {
      "field": "event.logstash_timestamp",
      "format": "date_time"
    },
    {
      "field": "event.start",
      "format": "date_time"
    },
    {
      "field": "formatted.discale.tst",
      "format": "date_time"
    },
    {
      "field": "process.parent.start",
      "format": "date_time"
    },
    {
      "field": "process.start",
      "format": "date_time"
    }
  ],
  "script_fields": {},
  "stored_fields": [
    "*"
  ],
  "runtime_mappings": {
    "aggregate_total_time": {
      "type": "long",
      "script": {
        "source": "String str = params._source['message'];\ndef m = /.*?\\[aggregate\\] Total time taken:(\\d+) ms .*/.matcher(str);\nif (m.matches()) {\n    emit(Integer.parseInt(m.group(1)))\n}"
      }
    },
    "aggregate_increments": {
      "type": "long",
      "script": {
        "source": "String str = params._source['message'];\ndef m = /.*?\\[aggregate\\] Total time taken:\\d+ ms to process (\\d+) increments.*/.matcher(str);\nif (m.matches()) {\n    emit(Integer.parseInt(m.group(1)))\n}"
      }
    }
  },
  "_source": {
    "excludes": []
  },
  "query": {
    "bool": {
      "must": [],
      "filter": [
        {
          "bool": {
            "filter": [
              {
                "multi_match": {
                  "type": "phrase",
                  "query": "%s",
                  "lenient": true
                }
              },
              {
                "bool": {
                  "should": [
                    {
                      "term": {
                        "environment": "%s"
                      }
                    }
                  ],
                  "minimum_should_match": 1
                }
              },
              {
                "bool": {
                  "should": [
                    {
                      "term": {
                        "log.file.path": "%s"
                      }
                    }
                  ],
                  "minimum_should_match": 1
                }
              }
            ]
          }
        },
        {
          "range": {
            "@timestamp": {
              "format": "strict_date_optional_time",
              "gte": "%s",
              "lte": "%s"
            }
          }
        }
      ],
      "should": [],
      "must_not": []
    }
  }
}`
	query = fmt.Sprintf(query, apis, env, filepath, startDate, EndDate)

	return query
}
