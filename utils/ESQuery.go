package utils

import (
	"fmt"
	"strings"
)

func buildmyquery(startDate string, EndDate string, apis string, cont string, filepath string, env string) string {
	var query string
	if strings.Contains(apis, "*and*") {
		partsofAPI := strings.Split(apis, "*and*")
		// Print the parts
		fmt.Println(partsofAPI[0]) // GET scheduler-service/api/v1/Organizations
		fmt.Println(partsofAPI[1]) // /Jobs
		query = `{
  "aggs": {
    "0": {
      "percentiles": {
        "field": "%s",
        "percents": [
          95,90
        ]
      }
    },"1": {
      "avg": {
        "field": "%s"
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
  "script_fields": {
    
  },
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
    "excludes": [
      
    ]
  },
  "query": {
    "bool": {
      "must": [
        
      ],
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
              },{
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
                      "wildcard": {
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
          "exists": {
            "field": "%s"
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
      "should": [
        
      ],
      "must_not": [
      ]
    }
  }
}`
		if strings.Contains(apis, "internal") {
			ResponseTime := "svctimeInMS"
			query = fmt.Sprintf(query, ResponseTime, ResponseTime, partsofAPI[0], partsofAPI[1], env, filepath, ResponseTime, startDate, EndDate)
		} else {
			ResponseTime := "temp.duration"
			haproxy := "/var/log/haproxy.log"
			query = fmt.Sprintf(query, ResponseTime, ResponseTime, partsofAPI[0], partsofAPI[1], env, haproxy, ResponseTime, startDate, EndDate)
		}
	} else {
		// String doesn't contain " and "
		//fmt.Println("String does not contain ' and '")
		query = `{
  "aggs": {
    "0": {
      "percentiles": {
        "field": "%s",
        "percents": [
          95,90
        ]
      }
    },"1": {
      "avg": {
        "field": "%s"
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
  "script_fields": {
    
  },
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
    "excludes": [
      
    ]
  },
  "query": {
    "bool": {
      "must": [
        
      ],
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
                      "wildcard": {
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
          "exists": {
            "field": "%s"
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
      "should": [
        
      ],
      "must_not": [
        
      ]
    }
  }
}`
		if strings.Contains(apis, "internal") {
			ResponseTime := "svctimeInMS"
			query = fmt.Sprintf(query, ResponseTime, ResponseTime, apis, env, filepath, ResponseTime, startDate, EndDate)
		} else {
			ResponseTime := "temp.duration"
			haproxy := "/var/log/haproxy.log"
			query = fmt.Sprintf(query, ResponseTime, ResponseTime, apis, env, haproxy, ResponseTime, startDate, EndDate)
		}
	}
	return query
}
