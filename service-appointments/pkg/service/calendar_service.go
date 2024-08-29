package service

type ICalendar interface {
	registerEvent(result any)
}

func RegisterEvent(calendar ICalendar, event *any) {

}

func NewCalendarService() *CalendarService {
	service := *NewService[any]()
	return &CalendarService{&service}
}

type CalendarService struct {
	*Service[any]
}

func (r *CalendarService) registerEvent(result any) {

}
