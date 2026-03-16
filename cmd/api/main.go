package main

import (
	"balancerAI/internal/app/api"
	"balancerAI/internal/app/balancer"
	"balancerAI/internal/app/config"
	llmmodels "balancerAI/internal/app/llm_models"
	"balancerAI/internal/app/observability/tracing"
	"balancerAI/internal/app/service"
	"context"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {

	flag.StringVar(&configPath, "path", "configs/config.toml", "path to config file in .toml format")
}
func main() {
	flag.Parse()
	log.Println("It works")

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	shutdown := tracing.InitTracer(
		"ai-balancer",
		cfg.Jaeger.URL,
	)
	defer shutdown(context.Background())

	chatModel := llmmodels.NewChatLLM(
		cfg.Models.Chat.Endpoint,
		cfg.Models.Chat.APIKey,
	)

	imageModel := llmmodels.NewImageLLM(
		cfg.Models.Image.Endpoint,
		cfg.Models.Image.APIKey,
	)

	audioModel := llmmodels.NewAudioLLM(
		cfg.Models.Audio.Endpoint,
		cfg.Models.Audio.APIKey,
	)

	balancer := balancer.NewBalancer(
		[]llmmodels.ChatModel{chatModel},
		[]llmmodels.ImageModel{imageModel},
		[]llmmodels.AudioModel{audioModel},
	)
	service := service.NewService(balancer)
	api := api.New(
		&config.Config{
			Server: cfg.Server,
			Logger: cfg.Logger,
			Jaeger: cfg.Jaeger,
			Models: cfg.Models,
		},
		service,
	)
	log.Fatal(api.Start())
}
