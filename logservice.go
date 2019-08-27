package main

type LogService struct {
	PageSize int
	models   Models
	size     int
}

func NewLogService(pageSize int, models Models) *LogService {
	models.Sort()

	return &LogService{
		PageSize: pageSize,
		models:   models,
		size:     len(models),
	}
}

func (service LogService) Query(cursor string) (Models, *string, *string, error) {
	var window *Window
	var previousWindow *Window
	var nextWindow *Window

	var err error

	if cursor == "" {
		window, err = NewWindow(0, service.PageSize)

		if err != nil {
			return nil, nil, nil, err
		}

		previousWindow = nil
		nextWindow = window.NextWindow(service.PageSize, service.size)
	} else {
		window, err = NewWindowFromCursor(cursor)

		if err != nil {
			return nil, nil, nil, err
		}

		previousWindow = window.PreviousWindow(service.PageSize)
		nextWindow = window.NextWindow(service.PageSize, service.size)
	}

	var previousCursor *string

	if previousWindow != nil {
		cursor := previousWindow.GetCursor()
		previousCursor = &cursor
	}

	var nextCursor *string

	if nextWindow != nil {
		cursor := nextWindow.GetCursor()
		nextCursor = &cursor
	}

	firstModels := service.models[window.Min:window.Max]

	return firstModels, previousCursor, nextCursor, nil
}
