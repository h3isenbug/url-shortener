// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/h3isenbug/url-shortener/internal/monitoring"
	"github.com/h3isenbug/url-shortener/pkg/log"
	"net/http"
	"sync"
)

// Injectors from inject.go:

func Inject() (*App, func(), error) {
	logger, err := provideLogger()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := provideSQLXConnection(logger)
	if err != nil {
		return nil, nil, err
	}
	metricCollector, cleanup2 := provideMetricCollector()
	repository := provideAccountRepository(db, metricCollector)
	refreshTokenRepository := provideRefreshTokenRepository(db, metricCollector)
	diAccessTokenSecretsType, err := provideAccessTokenSecrets()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service := provideAuthenticationService(logger, repository, refreshTokenRepository, diAccessTokenSecretsType)
	authenticationAPI := provideAuthenticationAPI(logger, service)
	client := provideRedisClient()
	urlRepository := provideUrlRepository(logger, db, client, metricCollector)
	urlService := provideUrlService(logger, urlRepository)
	urlAPI := provideUrlAPI(logger, urlService)
	router := provideMuxRouter(logger, service, authenticationAPI, urlAPI, metricCollector)
	server, cleanup3 := provideHTTPServer(logger, router)
	app := provideApp(logger, server, metricCollector)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// inject.go:

type App struct {
	httpServer      *http.Server
	metricCollector monitoring.MetricCollector
	logger          log.Logger
}

func (a App) Start() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Warn("http server terminated", map[string]interface{}{
				"errorMessage": err,
			})
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.metricCollector.Start(); err != nil && err != http.ErrServerClosed {
			a.logger.Warn("metric collector terminated", map[string]interface{}{
				"errorMessage": err,
			})
		}
	}()

	wg.Wait()
}

func (a App) Handler() http.Handler {
	return a.httpServer.Handler
}

func provideApp(logger log.Logger, httpServer *http.Server, metricCollector monitoring.MetricCollector) *App {
	return &App{
		httpServer:      httpServer,
		metricCollector: metricCollector,
		logger:          logger,
	}
}
