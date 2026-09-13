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

func FetchCategoriesService(module string) ([]CalendarCategory, error) {
	return FindCategories(module)
}

func SaveCategoryService(cat *CalendarCategory) error {
	return CreateCategoryRepo(cat)
}

func ModifyCategoryService(id string, input *CalendarCategory) (CalendarCategory, error) {
	cat, err := GetCategoryByID(id)
	if err != nil {
		return cat, err
	}

	cat.Name = input.Name
	cat.Color = input.Color
	cat.Icon = input.Icon
	cat.Description = input.Description
	if input.Module != "" {
		cat.Module = input.Module
	}

	err = UpdateCategoryRepo(&cat)
	return cat, err
}

func RemoveCategoryService(id string) error {
	return DeleteCategoryRepo(id)
}
