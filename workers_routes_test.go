package cloudflare

// func TestWorkers_CreateWorkerRoute(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/filters", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
// 	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
// 	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_CreateWorkerRouteEnt(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script"}
// 	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
// 	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_CreateWorkerRouteSingleScriptWithAccount(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/filters", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app1.example.com/*", Enabled: true}
// 	res, err := client.CreateWorkerRoute(context.Background(), "foo", route)
// 	want := WorkerRouteResponse{successResponse, WorkerRoute{ID: "e7a57d8746e74ae49c25994dadb421b1"}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_CreateWorkerRouteErrorsWhenMixingSingleAndMultiScriptProperties(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script", Enabled: true}
// 	_, err := client.CreateWorkerRoute(context.Background(), "foo", route)
// 	assert.EqualError(t, err, "Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
// }

// func TestWorkers_CreateWorkerRouteWithNoScript(t *testing.T) {
// 	setup()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, createWorkerRouteResponse) //nolint
// 	})

// 	route := WorkerRoute{Pattern: "app1.example.com/*"}
// 	_, err := client.CreateWorkerRoute(context.Background(), "foo", route)
// 	assert.NoError(t, err)
// }

// func TestWorkers_DeleteWorkerRoute(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, deleteWorkerRouteResponseData) //nolint
// 	})
// 	res, err := client.DeleteWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1")
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID: "e7a57d8746e74ae49c25994dadb421b1",
// 		}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_DeleteWorkerRouteEnt(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, deleteWorkerRouteResponseData) //nolint
// 	})
// 	res, err := client.DeleteWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1")
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID: "e7a57d8746e74ae49c25994dadb421b1",
// 		}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_ListWorkerRoutes(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/filters", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, listRouteResponseData) //nolint
// 	})

// 	res, err := client.ListWorkerRoutes(context.Background(), "foo")
// 	want := WorkerRoutesResponse{successResponse,
// 		[]WorkerRoute{
// 			{ID: "e7a57d8746e74ae49c25994dadb421b1", Pattern: "app1.example.com/*", Enabled: true},
// 			{ID: "f8b68e9857f85bf59c25994dadb421b1", Pattern: "app2.example.com/*", Enabled: false},
// 		},
// 	}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_ListWorkerRoutesEnt(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, listRouteEntResponseData) //nolint
// 	})

// 	res, err := client.ListWorkerRoutes(context.Background(), "foo")
// 	want := WorkerRoutesResponse{successResponse,
// 		[]WorkerRoute{
// 			{ID: "e7a57d8746e74ae49c25994dadb421b1", Pattern: "app1.example.com/*", Script: "test_script_1", Enabled: true},
// 			{ID: "f8b68e9857f85bf59c25994dadb421b1", Pattern: "app2.example.com/*", Script: "test_script_2", Enabled: true},
// 			{ID: "2b5bf4240cd34c77852fac70b1bf745a", Pattern: "app3.example.com/*", Script: "", Enabled: false},
// 		},
// 	}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_GetWorkerRoute(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes/1234", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, getRouteResponseData) //nolint
// 	})

// 	res, err := client.GetWorkerRoute(context.Background(), "foo", "1234")
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID:      "e7a57d8746e74ae49c25994dadb421b1",
// 			Pattern: "app1.example.com/*",
// 			Script:  "script-name"},
// 	}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_UpdateWorkerRoute(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, updateWorkerRouteResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app3.example.com/*", Enabled: true}
// 	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID:      "e7a57d8746e74ae49c25994dadb421b1",
// 			Pattern: "app3.example.com/*",
// 			Enabled: true,
// 		}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_UpdateWorkerRouteEnt(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app3.example.com/*", Script: "test_script_1"}
// 	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID:      "e7a57d8746e74ae49c25994dadb421b1",
// 			Pattern: "app3.example.com/*",
// 			Script:  "test_script_1",
// 		}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_UpdateWorkerRouteSingleScriptWithAccount(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/filters/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
// 	})
// 	route := WorkerRoute{Pattern: "app3.example.com/*", Enabled: true}
// 	res, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
// 	want := WorkerRouteResponse{successResponse,
// 		WorkerRoute{
// 			ID:      "e7a57d8746e74ae49c25994dadb421b1",
// 			Pattern: "app3.example.com/*",
// 			Script:  "test_script_1",
// 		}}
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestWorkers_ListWorkerBindingsMultiScript(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/my-script/bindings", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, listBindingsResponseData) //nolint
// 	})

// 	mux.HandleFunc("/accounts/"+testAccountID+"/workers/scripts/my-script/bindings/MY_WASM/content", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
// 		w.Header().Set("content-type", "application/wasm")
// 		_, _ = w.Write([]byte("mock multi-script wasm"))
// 	})

// 	res, err := client.ListWorkerBindings(context.Background(), &WorkerRequestParams{
// 		ScriptName: "my-script",
// 	})
// 	assert.NoError(t, err)

// 	assert.Equal(t, successResponse, res.Response)
// 	assert.Equal(t, 7, len(res.BindingList))

// 	assert.Equal(t, res.BindingList[0], WorkerBindingListItem{
// 		Name: "MY_KV",
// 		Binding: WorkerKvNamespaceBinding{
// 			NamespaceID: "89f5f8fd93f94cb98473f6f421aa3b65",
// 		},
// 	})
// 	assert.Equal(t, WorkerKvNamespaceBindingType, res.BindingList[0].Binding.Type())

// 	assert.Equal(t, "MY_WASM", res.BindingList[1].Name)
// 	wasmBinding := res.BindingList[1].Binding.(WorkerWebAssemblyBinding)
// 	wasmModuleContent, err := io.ReadAll(wasmBinding.Module)
// 	assert.NoError(t, err)
// 	assert.Equal(t, []byte("mock multi-script wasm"), wasmModuleContent)
// 	assert.Equal(t, WorkerWebAssemblyBindingType, res.BindingList[1].Binding.Type())

// 	assert.Equal(t, res.BindingList[2], WorkerBindingListItem{
// 		Name: "MY_PLAIN_TEXT",
// 		Binding: WorkerPlainTextBinding{
// 			Text: "text",
// 		},
// 	})
// 	assert.Equal(t, WorkerPlainTextBindingType, res.BindingList[2].Binding.Type())

// 	assert.Equal(t, res.BindingList[3], WorkerBindingListItem{
// 		Name:    "MY_SECRET_TEXT",
// 		Binding: WorkerSecretTextBinding{},
// 	})
// 	assert.Equal(t, WorkerSecretTextBindingType, res.BindingList[3].Binding.Type())

// 	environment := "MY_ENVIRONMENT"
// 	assert.Equal(t, res.BindingList[4], WorkerBindingListItem{
// 		Name: "MY_SERVICE_BINDING",
// 		Binding: WorkerServiceBinding{
// 			Service:     "MY_SERVICE",
// 			Environment: &environment,
// 		},
// 	})
// 	assert.Equal(t, WorkerServiceBindingType, res.BindingList[4].Binding.Type())

// 	assert.Equal(t, res.BindingList[5], WorkerBindingListItem{
// 		Name:    "MY_NEW_BINDING",
// 		Binding: WorkerInheritBinding{},
// 	})
// 	assert.Equal(t, WorkerInheritBindingType, res.BindingList[5].Binding.Type())

// 	assert.Equal(t, res.BindingList[6], WorkerBindingListItem{
// 		Name: "MY_BUCKET",
// 		Binding: WorkerR2BucketBinding{
// 			BucketName: "bucket",
// 		},
// 	})
// 	assert.Equal(t, WorkerR2BucketBindingType, res.BindingList[6].Binding.Type())
// }

// func TestWorkers_UpdateWorkerRouteErrorsWhenMixingSingleAndMultiScriptProperties(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	route := WorkerRoute{Pattern: "app1.example.com/*", Script: "test_script", Enabled: true}
// 	_, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
// 	assert.EqualError(t, err, "Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
// }

// func TestWorkers_UpdateWorkerRouteWithNoScript(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/zones"+testAccountID+"/workers/routes/e7a57d8746e74ae49c25994dadb421b1", func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
// 		w.Header().Set("content-type", "application/json")
// 		fmt.Fprintf(w, updateWorkerRouteEntResponse) //nolint
// 	})

// 	route := WorkerRoute{Pattern: "app1.example.com/*"}
// 	_, err := client.UpdateWorkerRoute(context.Background(), "foo", "e7a57d8746e74ae49c25994dadb421b1", route)
// 	assert.NoError(t, err)
// }
