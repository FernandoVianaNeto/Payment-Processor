package cli

import (
	"context"

	configs "payment-gateway/cmd/config"
	domain_repository "payment-gateway/internal/domain/repository"
	"payment-gateway/internal/infra/web"
	natsclient "payment-gateway/pkg/nats"
)

type Application struct {
	UseCases UseCases
}
type UseCases struct {
}

type Adapters struct {
}

type Repositories struct {
	Payments domain_repository.PaymentRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	eventClient := natsclient.New(configs.NatsCfg.Host)
	eventClient.Connect()

	srv := web.NewServer(
		ctx,
	)

	return srv
}

// func NewRepositories(
// 	ctx context.Context,
// ) Repositories {
// 	paymentRepsitory := redis_payment_repository.NewPaymentsRepository(db)

// 	return Repositories{
// 		PaymentsRepository: paymentRepsitory,
// 	}
// }

// func NewAdapters(
// 	ctx context.Context,
// ) Adapters {
// 	emailSenderAdapter := sendgrid.NewEmailSenderAdapter(ctx)
// 	minioAdapter := NewStorageAdapter(ctx)

// 	return Adapters{
// 		emailSenderAdapter: emailSenderAdapter,
// 		storageAdapter:     minioAdapter,
// 	}
// }

// func NewUseCases(
// 	ctx context.Context,
// 	userRepository domain_repository.UserRepositoryInterface,
// 	workoutRepository domain_repository.WorkoutRepositoryInterface,
// 	gamificationRepository domain_repository.GamificationRepositoryInterface,
// 	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
// 	services Services,
// 	adapters Adapters,
// 	eventClient messaging.Client,
// ) UseCases {
// 	userUsecase := user_usecase.NewCreateUserUseCase(userRepository, services.encryptStringService, adapters.storageAdapter)
// 	getUserUsecase := user_usecase.NewGetUserProfileUseCase(userRepository, adapters.storageAdapter)
// 	updateUserUsecase := user_usecase.NewUpdateUserUseCase(userRepository, adapters.storageAdapter)

// 	//AUTH
// 	authUsecase := auth_usecase.NewAuthUsecase(userRepository)
// 	googleAuthUsecase := auth_usecase.NewGoogleAuthUsecase(userRepository)
// 	generateResetPasswordCodeUsecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(resetPasswordCodeRepository, userRepository, adapters.emailSenderAdapter)
// 	resetPasswordUsecase := auth_usecase.NewResetPasswordUsecase(userRepository, resetPasswordCodeRepository, services.encryptStringService)
// 	validateResetPasswordCodeUsecase := auth_usecase.NewValidateResetPasswordCodeUsecase(resetPasswordCodeRepository)

// 	//TEMPORARY
// 	updateLeaderboardUsecase := gamification_usecase.NewUpdateLeaderboardUsecase(gamificationRepository, workoutRepository)

// 	createWorkoutUsecase := workout_usecase.NewCreateWorkoutUseCase(workoutRepository, updateLeaderboardUsecase, adapters.storageAdapter, eventClient)
// 	getWorkoutDetailsUsecase := workout_usecase.NewGetWorkoutDetailsUseCase(workoutRepository, adapters.storageAdapter)
// 	listUserWorkoutsUsecase := workout_usecase.NewListUserWorkoutsUseCase(workoutRepository, adapters.storageAdapter)
// 	addWorkoutInteractionUsecase := workout_usecase.NewAddInteractionUseCase(workoutRepository, updateLeaderboardUsecase)
// 	getInteractionsUsecase := workout_usecase.NewGetWorkoutInteractionUsecase(workoutRepository, userRepository, adapters.storageAdapter)
// 	deleteWorkoutUsecase := workout_usecase.NewDeleteWorkoutUseCase(workoutRepository)

// 	createGamificationUsecase := gamification_usecase.NewCreateGamificationUsecase(gamificationRepository, adapters.storageAdapter)
// 	getGamificationUsecase := gamification_usecase.NewGetGamificationUsecase(gamificationRepository, adapters.storageAdapter)
// 	addUsersUsecase := gamification_usecase.NewAddUserUsecase(gamificationRepository)
// 	getUserGamificationsUsecase := gamification_usecase.NewGetUserGamificationsUsecase(gamificationRepository, adapters.storageAdapter)

// 	return UseCases{
// 		userUseCase:                      userUsecase,
// 		GetUserUsecase:                   getUserUsecase,
// 		UpdateUserUsecase:                updateUserUsecase,
// 		AuthUsecase:                      authUsecase,
// 		GoogleAuthUsecase:                googleAuthUsecase,
// 		CreateWorkoutUsecase:             createWorkoutUsecase,
// 		GetWorkoutDetailsUsecase:         getWorkoutDetailsUsecase,
// 		ListUserWorkoutsUsecase:          listUserWorkoutsUsecase,
// 		CreateGamificationUsecase:        createGamificationUsecase,
// 		GetGamificationUsecase:           getGamificationUsecase,
// 		AddUserUsecase:                   addUsersUsecase,
// 		GetUserGamificationsUsecase:      getUserGamificationsUsecase,
// 		UpdateLeaderboardUsecase:         updateLeaderboardUsecase,
// 		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
// 		ResetPasswordUsecase:             resetPasswordUsecase,
// 		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
// 		AddWorkoutInteractionsUsecase:    addWorkoutInteractionUsecase,
// 		GetWorkoutInteractionUsecase:     getInteractionsUsecase,
// 		DeleteWorkoutUsecase:             deleteWorkoutUsecase,
// 	}
// }
