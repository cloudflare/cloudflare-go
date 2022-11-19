package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testQueueID           = "6b7efc370ea34ded8327fa20698dfe3a"
	testQueueName         = "example-queue"
	testQueueConsumerName = "example-consumer"
)

func testQueue() Queue {
	CreatedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	ModifiedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	return Queue{
		ID:                  testQueueID,
		Name:                testQueueName,
		CreatedOn:           &CreatedOn,
		ModifiedOn:          &ModifiedOn,
		ProducersTotalCount: 1,
		Producers: []QueueProducer{
			{
				Service:     "example-producer",
				Environment: "production",
			},
		},
		ConsumersTotalCount: 1,
		Consumers: []QueueConsumer{
			{
				Service:     "example-consumer",
				Environment: "production",
				Settings: QueueConsumerSettings{
					BatchSize:   10,
					MaxRetires:  3,
					MaxWaitTime: 5000,
				},
			},
		},
	}
}

func testQueueConsumer() QueueConsumer {
	CreatedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	return QueueConsumer{
		Service:     "example-consumer",
		Environment: "production",
		Settings: QueueConsumerSettings{
			BatchSize:   10,
			MaxRetires:  3,
			MaxWaitTime: 5000,
		},
		QueueName: testQueueName,
		CreatedOn: &CreatedOn,
	}
}

func TestQueue_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": null,
  "messages": null,
  "result": [
    {
      "queue_id": "6b7efc370ea34ded8327fa20698dfe3a",
      "queue_name": "example-queue",
      "created_on": "2023-01-01T00:00:00Z",
      "modified_on": "2023-01-01T00:00:00Z",
      "producers_total_count": 1,
      "producers": [
        {
          "service": "example-producer",
          "environment": "production"
        }
      ],
      "consumers_total_count": 1,
      "consumers": [
        {
          "service": "example-consumer",
          "environment": "production",
          "settings": {
            "batch_size": 10,
            "max_retries": 3,
            "max_wait_time_ms": 5000
          }
        }
      ]
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 100,
    "count": 1,
    "total_count": 1,
    "total_pages": 1
  }
}`)
	})

	_, err := client.QueueList(context.Background(), AccountIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	result, err := client.QueueList(context.Background(), AccountIdentifier(testAccountID))
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
		assert.Equal(t, testQueue(), result[0])
	}
}

func TestQueue_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": null,
		  "messages": null,
		  "result": {
			"queue_id": "6b7efc370ea34ded8327fa20698dfe3a",
			"queue_name": "example-queue",
			"created_on": "2023-01-01T00:00:00Z",
			"modified_on": "2023-01-01T00:00:00Z"
		}
	}`)
	})
	_, err := client.QueueCreate(context.Background(), AccountIdentifier(""), QueueCreateParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueCreate(context.Background(), AccountIdentifier(testAccountID), QueueCreateParams{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}
	results, err := client.QueueCreate(context.Background(), AccountIdentifier(testAccountID), QueueCreateParams{Name: "example-queue"})
	if assert.NoError(t, err) {
		CreatedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
		ModifiedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
		createdQueue := Queue{
			ID:         testQueueID,
			Name:       testQueueName,
			CreatedOn:  &CreatedOn,
			ModifiedOn: &ModifiedOn,
		}

		assert.Equal(t, createdQueue, results)
	}
}

func TestQueue_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s", testAccountID, testQueueName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": null
		}`)
	})
	err := client.QueueDelete(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.QueueDelete(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	err = client.QueueDelete(context.Background(), AccountIdentifier(testAccountID), testQueueName)
	assert.NoError(t, err)
}

func TestQueue_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s", testAccountID, testQueueID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `
		{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"queue_id": "6b7efc370ea34ded8327fa20698dfe3a",
			"queue_name": "example-queue",
			"created_on": "2023-01-01T00:00:00Z",
			"modified_on": "2023-01-01T00:00:00Z",
			"producers_total_count": 1,
			"producers": [
			  {
				"service": "example-producer",
				"environment": "production"
			  }
			],
			"consumers_total_count": 1,
			"consumers": [
			  {
				"service": "example-consumer",
				"environment": "production",
				"settings": {
				  "batch_size": 10,
				  "max_retries": 3,
				  "max_wait_time_ms": 5000
				}
			  }
			]
		  }
		}`)
	})

	_, err := client.QueueGet(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueGet(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	result, err := client.QueueGet(context.Background(), AccountIdentifier(testAccountID), testQueueID)
	if assert.NoError(t, err) {
		assert.Equal(t, testQueue(), result)
	}
}

func TestQueue_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s", testAccountID, testQueueID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": null,
		  "messages": null,
		  "result": {
			"queue_id": "6b7efc370ea34ded8327fa20698dfe3a",
			"queue_name": "renamed-example-queue",
			"created_on": "2023-01-01T00:00:00Z",
			"modified_on": "2023-01-01T00:00:00Z"
		}
	}`)
	})
	_, err := client.QueueUpdate(context.Background(), AccountIdentifier(""), "", Queue{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueUpdate(context.Background(), AccountIdentifier(testAccountID), "", Queue{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	results, err := client.QueueUpdate(context.Background(), AccountIdentifier(testAccountID), testQueueID, Queue{Name: "example-queue"})
	if assert.NoError(t, err) {
		CreatedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
		ModifiedOn, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
		createdQueue := Queue{
			ID:         testQueueID,
			Name:       "renamed-example-queue",
			CreatedOn:  &CreatedOn,
			ModifiedOn: &ModifiedOn,
		}

		assert.Equal(t, createdQueue, results)
	}
}

func TestQueue_ListConsumers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", testAccountID, testQueueID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `
	{
		  "success": true,
		  "errors": null,
		  "messages": null,
		  "result": [
			{
			  "service": "example-consumer",
			  "environment": "production",
			  "settings": {
				"batch_size": 10,
				"max_retries": 3,
				"max_wait_time_ms": 5000
			  },
			  "queue_name": "example-queue",
			  "created_on": "2023-01-01T00:00:00Z"
			}
		  ],
		  "result_info": {
			"page": 1,
			"per_page": 100,
			"count": 1,
			"total_count": 1,
			"total_pages": 1
		  }
		}`)
	})

	_, err := client.QueueListConsumers(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueListConsumers(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	result, err := client.QueueListConsumers(context.Background(), AccountIdentifier(testAccountID), testQueueID)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
		assert.Equal(t, testQueueConsumer(), result[0])
	}
}

func TestQueue_CreateConsumer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", testAccountID, testQueueName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"service": "example-consumer",
			"environment": "production",
			"settings": {
			  "batch_size": 10,
			  "max_retries": 3,
			  "max_wait_time_ms": 5000
			},
			"dead_letter_queue": "example-dlq",
			"queue_name": "example-queue",
			"created_on": "2023-01-01T00:00:00Z"
		  }
		}`)
	})

	_, err := client.QueueCreateConsumer(context.Background(), AccountIdentifier(""), "", QueueConsumer{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueCreateConsumer(context.Background(), AccountIdentifier(testAccountID), "", QueueConsumer{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	result, err := client.QueueCreateConsumer(context.Background(), AccountIdentifier(testAccountID), testQueueName, QueueConsumer{
		Service:     "example-consumer",
		Environment: "production",
	})
	if assert.NoError(t, err) {
		expectedQueueConsumer := testQueueConsumer()
		expectedQueueConsumer.DeadLetterQueue = "example-dlq"
		assert.Equal(t, expectedQueueConsumer, result)
	}
}

func TestQueue_DeleteConsumer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", testAccountID, testQueueName, testQueueConsumerName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": null
		}`)
	})

	err := client.QueueDeleteConsumer(context.Background(), AccountIdentifier(""), "", "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.QueueDeleteConsumer(context.Background(), AccountIdentifier(testAccountID), "", "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	err = client.QueueDeleteConsumer(context.Background(), AccountIdentifier(testAccountID), testQueueName, "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueConsumerName, err)
	}

	err = client.QueueDeleteConsumer(context.Background(), AccountIdentifier(testAccountID), testQueueName, testQueueConsumerName)
	assert.NoError(t, err)
}

func TestQueue_UpdateConsumer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", testAccountID, testQueueName, testQueueConsumerName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		  "success": true,
		  "errors": [],
		  "messages": [],
		  "result": {
			"service": "example-consumer",
			"environment": "production",
			"settings": {
			  "batch_size": 10,
			  "max_retries": 3,
			  "max_wait_time_ms": 5000
			},
			"queue_name": "example-queue",
			"created_on": "2023-01-01T00:00:00Z"
		  }
		}`)
	})

	_, err := client.QueueUpdateConsumer(context.Background(), AccountIdentifier(""), "", "", QueueConsumer{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.QueueUpdateConsumer(context.Background(), AccountIdentifier(testAccountID), "", "", QueueConsumer{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueName, err)
	}

	_, err = client.QueueUpdateConsumer(context.Background(), AccountIdentifier(testAccountID), testQueueName, "", QueueConsumer{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingQueueConsumerName, err)
	}

	result, err := client.QueueUpdateConsumer(context.Background(), AccountIdentifier(testAccountID), testQueueName, testQueueConsumerName, QueueConsumer{
		Service:     "example-consumer",
		Environment: "production",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, testQueueConsumer(), result)
	}
}
