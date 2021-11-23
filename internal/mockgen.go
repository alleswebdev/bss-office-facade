package internal

//go:generate mockgen -destination=./mocks/event_repo_mock.go -package=mocks github.com/ozonmp/bss-office-facade/internal/repo EventRepo
//go:generate mockgen -destination=./mocks/office_repo_mock.go -package=mocks github.com/ozonmp/bss-office-facade/internal/repo OfficeRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/bss-office-facade/internal/app/sender EventSender
