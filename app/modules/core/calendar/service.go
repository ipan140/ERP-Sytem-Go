package calendar

func FetchEventsService(userID uint, resModel, start, end string) ([]CalendarEvent, error) {
	return FindEvents(userID, resModel, start, end)
}

func SaveEventService(event *CalendarEvent) error {
	return CreateEventRepo(event)
}

func ModifyEventService(id string, input *CalendarEvent) (CalendarEvent, error) {
	event, err := GetEventByID(id)
	if err != nil {
		return event, err
	}

	// Update fields safely
	event.Title = input.Title
	event.Start = input.Start
	event.End = input.End
	event.AllDay = input.AllDay
	event.Color = input.Color
	event.Location = input.Location
	event.MeetingURL = input.MeetingURL
	event.Visibility = input.Visibility
	event.AttendeeIDs = input.AttendeeIDs

	err = UpdateEventRepo(&event)
	return event, err
}

func RemoveEventService(id string) error {
	return DeleteEventRepo(id)
}
