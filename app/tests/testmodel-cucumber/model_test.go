package main

import (
	"os"
	"testing"

	cucumber "github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var opts = cucumber.Options{Output: colors.Colored(os.Stdout), Concurrency: 1}

func init() {
	cucumber.BindCommandLineFlags("*.", &opts) // all features
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := cucumber.TestSuite{
		Name:                "watson model",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}

var intent, entity string

func InitializeScenario(ctx *cucumber.ScenarioContext) {
	ctx.Step(`^user asks to watson: \"([^\"]*)\"$`, userAsks)
	ctx.Step(`^watson should respond with intent: \"([^\"]*)\"$`, watsonIntent)
	ctx.Step(`^watson should respond with entity: \"([^\"]*)\"$`, watsonEntity)
}

func userAsks(input string) error {
	resp := askToWatson(input)

	intent = ""
	if len(resp.Output.Intents) > 0 {
		intent = resp.Output.Intents[0].Intent
	}

	entity = ""
	if len(resp.Output.Entities) > 0 {
		entity = resp.Output.Entities[0].Value
	}

	return nil
}

func watsonIntent(expected string) error {
	return assertExpectedAndActual(
		assert.Equal,
		expected,
		intent,
		"Expected %d, but there is %d",
		expected,
		intent,
	)
}

func watsonEntity(expected string) error {
	return assertExpectedAndActual(
		assert.Equal,
		expected,
		entity,
		"Expected %d, but there is %d",
		expected,
		entity,
	)
}
