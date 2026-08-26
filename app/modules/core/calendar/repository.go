package calendar

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
)

func FindEvents(userID uint, resModel string, start string, end string) ([]CalendarEvent, error) {
	var events []CalendarEvent
	query := config.DB.Model(&CalendarEvent{}).Preload("User").Preload("Attendees")

	if resModel != "" {
		query = query.Where("res_model = ?", resModel)
	}

	query = query.Where(
		config.DB.Where("visibility = ?", "public").
			Or("user_id = ?", userID).
			Or("id IN (SELECT calendar_event_id FROM calendar_event_attendees WHERE user_id = ?)", userID),
	)

	if start != "" && end != "" {
		query = query.Where("start >= ? AND start <= ?", start, end)
	}

	err := query.Find(&events).Error
	return events, err
}

func CreateEventRepo(event *CalendarEvent) error {
	if len(event.AttendeeIDs) > 0 {
		var users []*auth.User
		config.DB.Where("id IN ?", event.AttendeeIDs).Find(&users)
		event.Attendees = users
	}
	return config.DB.Create(event).Error
}

func GetEventByID(id string) (CalendarEvent, error) {
	var event CalendarEvent
	err := config.DB.First(&event, id).Error
	return event, err
}

func UpdateEventRepo(event *CalendarEvent) error {
	if len(event.AttendeeIDs) > 0 {
		var users []*auth.User
		config.DB.Where("id IN ?", event.AttendeeIDs).Find(&users)
		config.DB.Model(event).Association("Attendees").Replace(users)
	}
	return config.DB.Save(event).Error
}

func DeleteEventRepo(id string) error {
	return config.DB.Delete(&CalendarEvent{}, id).Error
}
