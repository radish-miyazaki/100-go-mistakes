package main

import (
	"log"
	"net/http"
)

var tracing bool

func listening1() error {
	var client *http.Client

	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}

		log.Println(client)
	} else {
		client, err := createClientDefault()
		if err != nil {
			return err
		}

		log.Println(client)
	}

	// client を使用する
	_ = client

	return nil
}

func listening2() error {
	var client *http.Client

	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}

		client = c

		log.Println(client)
	} else {
		c, err := createClientDefault()
		if err != nil {
			return err
		}

		client = c

		log.Println(client)
	}

	// client を使用する
	_ = client

	return nil
}

func listening3() error {
	var client *http.Client
	var err error

	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}

		log.Println(client)
	} else {
		client, err = createClientDefault()
		if err != nil {
			return err
		}

		log.Println(client)
	}

	// client を使用する
	_ = client

	return nil
}

func listening4() error {
	var client *http.Client
	var err error

	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createClientDefault()
	}

	if err != nil {
		return err
	}

	log.Println(client)

	// client を使用する
	_ = client

	return nil
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createClientDefault() (*http.Client, error) {
	return nil, nil
}
