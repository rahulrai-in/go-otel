package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func init() {
	// Initialize OpenTelemetry
	initOpenTelemetry()
}

func initOpenTelemetry() {
	// Create a new export pipeline using stdout exporter
	// This will show traces in the console
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}

	// Configure the SDK with the exporter, trace provider, and default sampler
	bsp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
	otel.SetTracerProvider(tp)

	// Get a tracer instance from the global trace provider
	tracer = otel.Tracer("todo-app")
}

func executeChoiceInstrumented(choice int) {
	switch choice {
	case 1:
		addTaskInstrumented()
	case 2:
		listTasksInstrumented()
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}

func addTaskInstrumented() {
	// Start a new span for the addTask function
	var ctx, span = tracer.Start(context.Background(), "addTaskProcess")
	defer span.End()

	// Child span for getting user input
	_, inputSpan := tracer.Start(ctx, "getUserInput")
	fmt.Print("Enter task title: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	title = strings.TrimSuffix(title, "\n")
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	inputSpan.End()

	// Child span for updating the in-memory list
	_, updateSpan := tracer.Start(ctx, "updateTaskList")
	idCounter++
	task := Task{
		ID:    idCounter,
		Title: title,
	}
	tasks = append(tasks, task)
	updateSpan.End()

	// Add an event to the span for task addition
	span.AddEvent("Task Added", trace.WithAttributes())
	span.End()

	fmt.Println("Task added successfully!")
}

func listTasksInstrumented() {
	// Start a new span for the listTasks function
	var _, span = tracer.Start(context.Background(), "listTasks")
	defer span.End()

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Title)
	}

	// Add an event to the span for task listing
	span.AddEvent(fmt.Sprintf("Listed %d tasks", len(tasks)), trace.WithAttributes())
}
