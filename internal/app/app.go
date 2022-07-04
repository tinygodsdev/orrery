package app

import (
	"context"
	"time"

	"github.com/DanielTitkov/orrery/internal/util"

	"github.com/DanielTitkov/orrery/internal/configs"
	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/logger"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type (
	App struct {
		Cfg           configs.Config
		log           *logger.Logger
		repo          Repository
		systemSummary *domain.SystemSymmary
		Store         sessions.Store
		locales       []string // locale count is not very big so no need to have map
	}
	Repository interface {
		// user
		IfEmailRegistered(context.Context, string) (bool, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)

		// user session
		IfSessionRegistered(context.Context, *domain.UserSession) (bool, error)
		CreateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		CreateOrUpdateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		UpdateUserSessionLastActivityBySID(context.Context, string) error
		GetUserBySession(context.Context, *domain.UserSession) (*domain.User, error)

		// test
		CreateOrUpdateTestFromArgs(context.Context, *domain.CreateTestArgs) error
		GetTests(ctx context.Context, locale string, tagIDs []uuid.UUID) ([]*domain.Test, error)
		GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error)
		GetTakeData(ctx context.Context, take *domain.Take, locale string) (*domain.Test, error)

		// take
		GetTake(ctx context.Context, takeID uuid.UUID) (*domain.Take, error)
		CreateTake(ctx context.Context, take *domain.Take) (*domain.Take, error)
		UpdateTake(ctx context.Context, take *domain.Take) (*domain.Take, error)

		// response
		AddOrUpdateResponse(ctx context.Context, takeID uuid.UUID, itm *domain.Item) (*domain.Response, error)

		// sample
		CreateOrUpdateSample(ctx context.Context, sample *domain.Sample) (*domain.Sample, error)
		GetSamples(ctx context.Context) ([]*domain.Sample, error)

		// norm
		CreateOrUpdateNorm(ctx context.Context, norm *domain.Norm) (*domain.Norm, error)
		GetDataForNormCalculation(ctx context.Context, criteria domain.SampleCriteria) ([]*domain.NormCalculationData, error)
		GetScaleNorms(ctx context.Context, scaleID uuid.UUID, sampleIDs []uuid.UUID) ([]*domain.Norm, error)

		// result
		CreateOrUpdateResult(ctx context.Context, result *domain.Result) (*domain.Result, error)

		// for system summary
		GetUserCount(ctx context.Context) (int, error)

		// tag
		CreateOrUpdateTagFromArgs(ctx context.Context, args *domain.CreateTagArgs) error
		GetTagsByCodes(ctx context.Context, locale string, codes ...string) ([]*domain.Tag, error)
		GetTagIDsByCodes(ctx context.Context, codes ...string) ([]uuid.UUID, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	store sessions.Store,
) (*App, error) {
	defer util.InfoExecutionTime(time.Now(), "app.New", logger)
	app := App{
		Cfg:     cfg,
		log:     logger,
		repo:    repo,
		Store:   store,
		locales: domain.Locales(),
	}

	err := app.loadTagPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadUserPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadTestPresets()
	if err != nil {
		return nil, err
	}

	app.log.Info("finished loading presets", "")

	err = app.initSamples()
	if err != nil {
		return nil, err
	}

	err = app.UpdateNorms(context.TODO())
	if err != nil {
		return nil, err
	}

	// init app jobs, caches and preload data (if any)
	// TODO: move to jobs
	go app.UpdateSystemSummaryJob()
	go app.UpdateNormsJob()

	return &app, nil
}
